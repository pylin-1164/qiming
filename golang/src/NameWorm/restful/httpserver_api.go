package restful

import (
	"NameWorm/cnnumber"
	"NameWorm/common"
	"NameWorm/db"
	"NameWorm/utils/aesutil"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/emicklei/go-restful"
	"io/ioutil"
)

const (
	LIMIT_TYPE_CENTER = "center"
	LIMIT_TYPE_END = "end"
	LIMIT_IS_SINGLE = "single"
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

}

func (v ApiNameEvaluate) RegistRoute(server *restful.WebService) {
	server.Route(server.POST("/api/name/parse").Filter(aesFilter).To(v.nameEvaluate).
		Doc("post NameEvaluate"))
	fmt.Printf("listener join post api : %s \n","/api/name/parse")
	server.Route(server.POST("/api/name/grasp").Filter(aesFilter).To(v.nameGrasp).
		Doc("post NameGrasp"))
	fmt.Printf("listener join post api : %s \n","/api/name/grasp")
	server.Route(server.POST("/api/name/links").To(v.links).Doc("post Links"))

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

	preDatas := cnnumber.GetNameIndexPage()
	numNameData := cnnumber.BuildNumNameData(v.FirsName, v.SuffixName, v.Gender, v.BirthYear, v.BirthMonth, v.BirthDay)
	numberCal := cnnumber.GetNameNumber(preDatas, numNameData)
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
	if err := request.ReadEntity(&v);err != nil{
		result := fmt.Sprintf(`{"resultstatus":"0","errorcode":"-2","errorinfo":"request parameter err "}`)
		fmt.Printf("error : 	request read error : %v \n",err)
		response.Write([]byte(result))
		return
	}
	
	firstName := v.FirsName
	sql := "SELECT ID FROM 	t_first_name WHERE first_name = ?"
	result, e := db.Conn.QueryFirst(sql, firstName)
	if result == nil || result["ID"] == nil{
		result := fmt.Sprintf(`{"resultstatus":"0","errorcode":"-1","errorinfo":"检索不到该姓氏"}`)
		fmt.Printf("error : 	request read error : %v \n",e)
		response.Write([]byte(result))
		return
	}
	if v.Gender != "1" && v.Gender != "2"{
		result := fmt.Sprintf(`{"resultstatus":"0","errorcode":"-1","errorinfo":"数据非法"}`)
		fmt.Printf("error : 	request read error : %v \n",e)
		response.Write([]byte(result))
		return
	}
	firstNameId := fmt.Sprintf("%d",result["ID"])



	resultList := v.randomNameList(firstNameId, firstName)

	if len(resultList) == 0{
		result := fmt.Sprintf(`{"resultstatus":"0","errorcode":"-1","errorinfo":"系统未录入该姓名相关素材"}`)
		fmt.Printf("error : 	request read error : %v \n",e)
		response.Write([]byte(result))
		return
	}

	resultMap := make(map[string]interface{})
	resultMap["resultstatus"] = "1"
	resultMap["list"] = resultList
	if data, e := json.Marshal(resultMap);e == nil{
		fmt.Printf("%s \n",data)
		response.Header().Add("Content-Type","application/json")
		response.Write(data)
	}


}

func (v ApiNameEvaluate) randomNameList(firstNameId string,firstName string) []string{
	limitWord := v.LimitWord
	limitType := v.LimitType
	gender := v.Gender

	sql := `SELECT
			* 
			FROM
				t_suffix_name AS t1
				JOIN (
					SELECT
						ROUND(
							RAND() * (
							(	 
								SELECT MAX( id ) FROM t_suffix_name WHERE first_name_id = '`+firstNameId+`' AND sex = '`+gender+`' %s ) 
								- ( SELECT MIN( id ) FROM t_suffix_name WHERE first_name_id = '`+firstNameId+`' AND sex = '`+gender+`' %s ) 
							) + ( SELECT MIN( id ) FROM t_suffix_name WHERE first_name_id = '`+firstNameId+`' AND sex = '`+gender+`' %s ) 
						) AS tid 
					) AS t2 
				WHERE
				t1.id >= t2.tid 
				AND t1.first_name_id = '`+firstNameId+`' %s
				AND t1.sex = '`+gender+`' 
				ORDER BY t1.id LIMIT 1`
	if v.SingleName == LIMIT_IS_SINGLE{
		limitsql := " AND CHAR_LENGTH(suffix_name) <2 "
		limitTsql := " AND CHAR_LENGTH(t1.suffix_name) <2"
		sql = fmt.Sprintf(sql,limitsql,limitsql,limitsql,limitTsql)
	}else{
		if limitType != "" && limitWord != ""{
			limitsql := "AND suffix_name LIKE '%"+limitWord+"'"
			limitTsql := "AND t1.suffix_name LIKE '%"+limitWord+"'"
			if limitType == LIMIT_TYPE_CENTER{
				limitsql = "AND suffix_name LIKE '"+limitWord+"%'"
				limitTsql = "AND t1.suffix_name LIKE '"+limitWord+"%'"
			}
			sql = fmt.Sprintf(sql,limitsql,limitsql,limitsql,limitTsql)
		}else{
			charset := " "
			sql = fmt.Sprintf(sql,charset,charset,charset,charset)
		}
	}
	//每次随机查询出一条，遍历20次
	resultList := make([]string, 0)
	nameset := set.NewHashSet()
	for i:=0;i<40;i++{
		if result, _ := db.Conn.QueryFirst(sql);result != nil{
			nameset.Add(fmt.Sprintf("%s%s",firstName,result["suffix_name"]))
		}
	}
	for i,e := range nameset.Elements() {
		if i >= 20{
			break
		}
		resultList = append(resultList[:], fmt.Sprintf("%s",e))
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