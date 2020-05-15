package server

import (
	"NameWorm/cnnumber"
	"NameWorm/db"
	"NameWorm/utils/aesutil"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/emicklei/go-restful"
	"io/ioutil"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

const (
	LIMIT_TYPE_CENTER = "center"
	LIMIT_TYPE_END = "end"
	LIMIT_IS_SINGLE = "1"
	LIMIT_IS_DOUBLE = "0"
)

type ApiNameEvaluate struct {
	FirsName 		string 		`json:"firstName"`
	SuffixName 		string		`json:"suffixName"`
	Gender 			string		`json:"gender"`
	BirthYear		string 		`json:"year"`
	BirthMonth		string 		`json:"month"`
	BirthDay		string		`json:"day"`
	LimitWord 		string 		`json:"limitWord"` //特定字
	LimitType		string 		`json:"limitType"` //特定字位置{center:居中，end:末尾}
	SingleName 		string 		`json:"single"`	//单字 {0:否，1：是}
	LicenseCode		string		`json:"licenseCode"` // 授权码

}

func (v ApiNameEvaluate) RegistRoute(server *restful.WebService) {
	server.Route(server.POST("/api/name/parse").Filter(aesFilter).To(v.nameEvaluate).
		Doc("post NameEvaluate"))
	fmt.Printf("listener join post api : %s \n","/api/name/parse")

	server.Route(server.POST("/api/name/grasp").Filter(aesFilter).To(v.nameGrasp).
		Doc("post NameGrasp"))
	fmt.Printf("listener join post api : %s \n","/api/name/grasp")

	server.Route(server.POST("/api/name/links").To(v.links).Doc("post Links"))
	fmt.Printf("listener join post api : %s \n","/api/name/links")

	server.Route(server.POST("/api/name/search").Filter(aesFilter).To(v.searchName).Doc("post Links"))
	fmt.Printf("listener join post api : %s \n","/api/name/search")

}

// Route Filter (defines FilterFunction)
func aesFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	// wrap responseWriter into a compressing one
	var bodyBytes []byte
	bodyBytes, _ = ioutil.ReadAll(req.Request.Body)
	if bodyBytes != nil{
		if decrypt, e := aesutil.Decrypt(fmt.Sprintf("%s", bodyBytes)); e != nil || decrypt==""{
			resp.WriteErrorString(401, "401: 参数错误")
			return
		}else  {
			bodyBytes = []byte(decrypt)
		}
	}
	req.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	chain.ProcessFilter(req, resp)
}

func init() {
	ApiNameEvaluate{}.RegistRoute(WS)
}

/**
 * 解析姓名
 */
func (v ApiNameEvaluate) nameEvaluate(request *restful.Request, response *restful.Response){


	if err := request.ReadEntity(&v);err != nil{
		result := fmt.Sprintf(`{"resultstatus":"0","errorcode":"-2","errorinfo":"request parameter err "}`)
		fmt.Printf("error : 	request read error : %v \n",err)
		response.Write([]byte(result))
		return
	}

	fmt.Println("licensCode: ",v.LicenseCode)

	preDatas := cnnumber.GetNameIndexPage()
	numNameData := cnnumber.BuildNumNameData(v.FirsName, v.SuffixName, v.Gender, v.BirthYear, v.BirthMonth, v.BirthDay)
	numberCal := cnnumber.GetNameNumber(preDatas,numNameData)
	if len(numberCal.Scores) == 0{
		result := fmt.Sprintf(`{"resultstatus":"0","errorcode":"-2","errorinfo":"解析测算姓名失败，未匹配该名字素材 "}`)
		response.Write([]byte(result))
	}
	if data, e := json.Marshal(numberCal);e != nil{
		fmt.Println(e)
	}else{
		response.Header().Add("Content-Type","application/json")
		response.Write(data)
	}
}

/**
 * 获取姓名
 */
func (v ApiNameEvaluate) nameGrasp(request *restful.Request, response *restful.Response){
	if v.beforeCheck(request, response) {
		return
	}

	firstName := v.FirsName
	sql := "SELECT ID FROM 	t_first_name WHERE first_name = ?"
	result, e := db.Conn.QueryFirst(sql, firstName)
	if result == nil || result["ID"] == nil {
		result := fmt.Sprintf(`{"resultstatus":"0","errorcode":"-1","errorinfo":"检索不到该姓氏"}`)
		fmt.Printf("error : 	request read error : %v \n", e)
		response.Write([]byte(result))
		return
	}
	if v.Gender != "1" && v.Gender != "2" {
		result := fmt.Sprintf(`{"resultstatus":"0","errorcode":"-1","errorinfo":"数据非法"}`)
		fmt.Printf("error : 	request read error : %v \n", e)
		response.Write([]byte(result))
		return
	}
	firstNameId := fmt.Sprintf("%d", result["ID"])

	resultList := v.randomNameList(firstNameId, firstName)

	if len(resultList) == 0 {
		result := fmt.Sprintf(`{"resultstatus":"0","errorcode":"-1","errorinfo":"系统未录入该姓名相关素材"}`)
		fmt.Printf("error : 	request read error : %v \n", e)
		response.Write([]byte(result))
		return
	}

	resultMap := make(map[string]interface{})
	resultMap["resultstatus"] = "1"
	resultMap["list"] = resultList
	if data, e := json.Marshal(resultMap); e == nil {
		fmt.Printf("%s \n", data)
		response.Header().Add("Content-Type", "application/json")
		response.Write(data)
	}

}

func (v ApiNameEvaluate) searchName(request *restful.Request, response *restful.Response){
	if v.beforeCheck(request, response) {
		return
	}
	sql := "SELECT * FROM t_suffix_name t1 LEFT JOIN t_first_name t2 ON t1.first_name_id=t2.id WHERE t2.first_name=? AND t1.suffix_name = ? ;"
	if result, e := db.Conn.Query(sql,v.FirsName,v.SuffixName);e != nil{
		result := fmt.Sprintf(`{"resultstatus":"0","errorcode":"-1","errorinfo":"查询失败"}`)
		response.Write([]byte(result))
		return
	}else{
		if len(result)>0{
			return
		}

		goto saveSearchName
	}

	saveSearchName:func(){
		sql = "SELECT * FROM t_first_name WHERE first_name = ?"
		result, _ := db.Conn.QueryFirst(sql, v.FirsName)
		if len(result) == 0{
			result := fmt.Sprintf(`{"resultstatus":"0","errorcode":"-1","errorinfo":"姓氏不存在"}`)
			response.Write([]byte(result))
			return
		}
		i := fmt.Sprintf("%d",result["id"])
		sql = "INSERT INTO  t_suffix_name(first_name_id,suffix_name,sex) VALUES(?,?,?)"
		db.Conn.Insert(sql,i,v.SuffixName,v.Gender)
	}()
}

func (v *ApiNameEvaluate) beforeCheck(request *restful.Request, response *restful.Response) bool {
	bt, _ := ioutil.ReadAll(request.Request.Body)
	if err := json.Unmarshal(bt, &v); err != nil {
		result := fmt.Sprintf(`{"resultstatus":"0","errorcode":"-2","errorinfo":"request parameter err "}`)
		fmt.Printf("error : 	request read error : %v \n", err)
		response.Write([]byte(result))
		return true
	}
	fmt.Println("licensCode: ", v.LicenseCode)

	if runtime.GOOS != "windows" {
		if v.LicenseCode == "" || LicenseCache[v.LicenseCode] == 0 {
			result := fmt.Sprintf(`{"resultstatus":"0","errorcode":"-1","errorinfo":"授权码错误"}`)
			response.Write([]byte(result))
			return true
		} else if LicenseCache[v.LicenseCode] < time.Now().Unix() {
			result := fmt.Sprintf(`{"resultstatus":"0","errorcode":"-1","errorinfo":"授权码过期"}`)
			response.Write([]byte(result))
			return true
		}

	}
	return false
}

func (v ApiNameEvaluate) randomNameList(firstNameId string,firstName string) []string{
	resultList := make([]string, 0)

	limitWord := v.LimitWord
	limitType := v.LimitType
	gender := v.Gender

	countSql := `SELECT COUNT(t1.id) as count FROM t_suffix_name t1 WHERE t1.first_name_id = '`+firstNameId+`' AND t1.sex = '`+gender+`' %s `

	sql := `SELECT t1.*,(@i:=@i+1) as i
		FROM t_suffix_name t1,(select @i:=0) t2 WHERE t1.first_name_id = '`+firstNameId+`' AND t1.sex = '`+gender+`' %s`

	if v.SingleName == LIMIT_IS_SINGLE{
		limitTsql := ` AND CHAR_LENGTH(t1.suffix_name) =1 %s`
		sql = fmt.Sprintf(sql,limitTsql)
		countSql = fmt.Sprintf(countSql,limitTsql)

	}else if v.SingleName == LIMIT_IS_DOUBLE{
		limitTsql := ` AND CHAR_LENGTH(t1.suffix_name) =2 %s`
		sql = fmt.Sprintf(sql,limitTsql)
		countSql = fmt.Sprintf(countSql,limitTsql)
	}

	if limitType != "" && limitWord != ""{
		limitTsql := "AND t1.suffix_name LIKE '%"+limitWord+"'"
		if limitType == LIMIT_TYPE_CENTER{
			limitTsql = "AND t1.suffix_name LIKE '"+limitWord+"%'"
		}
		sql = fmt.Sprintf(sql,limitTsql)
		countSql = fmt.Sprintf(countSql,limitTsql)
	}else if limitWord != ""{
		limitTsql := "AND t1.suffix_name = '"+limitWord+"'"
		sql = fmt.Sprintf(sql,limitTsql)
		countSql = fmt.Sprintf(countSql,limitTsql)
	} else{
		charset := " "
		sql = fmt.Sprintf(sql,charset)
		countSql = fmt.Sprintf(countSql,charset)
	}

	count := 0
	startIndex := 0
	step := 40
	if result, _ := db.Conn.QueryFirst(countSql);result != nil{
		count, _ = strconv.Atoi(fmt.Sprintf("%s", result["count"]))

	}
	if count == 0{
		return resultList
	}else if count > step {
		startIndex = rand.Intn(count-step+1)
	}
	sql = fmt.Sprintf("SELECT a.* FROM ( %s ) a WHERE a.i > %d LIMIT %d ",sql,startIndex,step)

	//每次随机查询出一条，遍历40次
	nameResult, _ := db.Conn.Query(sql);
	for _,v := range nameResult {
		resultList = append(resultList,fmt.Sprintf("%s%s",firstName,v["suffix_name"]))
	}

	return resultList
}

type Links struct {
	QQ 			string	`json:"qq"`
	QQgroup 	string 	`json:"qqg"`
	WX		string 	`json:"weixin"`
	WXgroup	string	`json:"weixing"`
}

func (v ApiNameEvaluate)links(request *restful.Request, response *restful.Response){
	links := Links{"1264712553", "", "", ""}
	if data, e := json.Marshal(links);e == nil{
		fmt.Printf("%s \n",data)
		response.Header().Add("Content-Type","application/json")
		response.Write(data)
	}
}