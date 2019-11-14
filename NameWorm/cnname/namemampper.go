package cnname

import (
	"fmt"
	"pyl/NameWorm/db"
)

func DeleteAll(){
	sql := "DELETE FROM t_first_name"
	db.Conn.Delete(sql)
}

func SaveName2Db(firstName string,suffixNames []SuffixName){
	sql := "INSERT INTO t_first_name(first_name) values(?)"
	result, e := db.Conn.Insert(sql, firstName)
	if e!=nil{
		panic(e)
	}
	firstid,_ := result.LastInsertId()
	i :=0
	for i< len(suffixNames){
		sql = "INSERT INTO t_suffix_name(first_name_id,suffix_name,sex) VALUES"
		for {
			suffixName := suffixNames[i]
			value := fmt.Sprintf(`(%d,"%s",%d)`, firstid, suffixName.Suffix, suffixName.Sex)
			sql = sql+value
			if i%500 == 499 || i == (len(suffixNames)-1){
				db.Conn.Insert(sql)
				i++
				break
			}else{
				sql = sql+","
			}
			i++

		}

	}


}


