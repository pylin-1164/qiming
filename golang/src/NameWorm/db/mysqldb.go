package db

import (
	. "NameWorm/common"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type MysqlConnect struct {
	Host 		string
	DbName 		string
	UserName	string
	Password 	string
	db 			*sql.DB
}

var Conn *MysqlConnect

func init() {
	Conn = &MysqlConnect{Host:DataBaseInfo.Host,DbName:DataBaseInfo.Name,UserName:DataBaseInfo.User,Password:DataBaseInfo.Password}
	//connStr := "%s:%s@127.0.0.1:3306/test?charset=utf8"
	connStr :=  strings.Join([]string{Conn.UserName, ":", Conn.Password, "@tcp(",Conn.Host, ":",DataBaseInfo.Port, ")/", Conn.DbName, "?charset=utf8"}, "")
	db, err := sql.Open("mysql", connStr)
	if err != nil{
		panic(err)
	}
	Conn.db = db
}

func (conn *MysqlConnect)Insert(sql string,args ...interface{}) (sql.Result,error){
	stmt, e := conn.db.Prepare(sql)
	if e != nil{
		return nil,e
	}

	result, e := stmt.Exec(args[:]...)
	stmt.Close()
	return result,e
}

func (conn *MysqlConnect)Delete(sql string,args ...interface{})(sql.Result,error){
	stmt, e := conn.db.Prepare(sql)
	if e != nil{
		return nil,e
	}

	result, e := stmt.Exec(args[:]...)
	stmt.Close()
	return result,e
}

func (conn *MysqlConnect)QueryFirst(sql string,args ...interface{})(map[string]interface{},error){
	item := make(map[string]interface{},0)
	rows, e := conn.db.Query(sql, args[:]...)
	defer rows.Close()
	if e != nil{
		return nil,e
	}else{
		columns, _ := rows.Columns()
		columnLength := len(columns)
		cache := make([]interface{}, columnLength) //临时存储每行数据
		for index, _ := range cache { //为每一列初始化一个指针
			var a interface{}
			cache[index] = &a
		}
		if rows.Next(){
			_ = rows.Scan(cache...)
			for i, data := range cache {
				item[columns[i]] = *data.(*interface{}) //取实际类型
			}
		}
	}
	return item,nil
}

func (conn *MysqlConnect)Query(sql string,args ...interface{})([]map[string]interface{},error) {
	var list []map[string]interface{} //返回的切片
	rows, e := conn.db.Query(sql, args[:]...)
	if e != nil{
		return nil,e
	}else{
		columns, _ := rows.Columns()
		columnLength := len(columns)
		cache := make([]interface{}, columnLength) //临时存储每行数据
		for index, _ := range cache { //为每一列初始化一个指针
			var a interface{}
			cache[index] = &a
		}

		for rows.Next() {
			_ = rows.Scan(cache...)

			item := make(map[string]interface{})
			for i, data := range cache {
				item[columns[i]] = *data.(*interface{}) //取实际类型
			}
			list = append(list, item)
		}
		_ = rows.Close()
	}
	return list,nil
}
