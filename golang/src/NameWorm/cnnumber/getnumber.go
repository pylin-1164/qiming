package cnnumber

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

//TODO 爬取地址【https://www.meimingteng.com/yy/】的打分
var REQUEST_NAME_NUMBER_URL="https://www.meimingteng.com/Naming/Default.aspx?Tag=4"
//var REQUEST_NAME_NUMBER_URL="http://172.16.138.90:20028/Naming/Default.aspx?Tag=4"


type GetNumPreDatas struct {
	Cookies 	map[string]string
	Inputs		map[string]string
}

type GetNumNameData struct {
	FirstName 		string 	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$tbXing"` 					//姓
	SuffixName 		string 	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$tbMingWords"` 				//名
	Genders   		string	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$ddlGenders"`				//性别 1[男]2[女]
	BirthDayMode 	string	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$SPECIFY_BIRHDAY"`			//生日类型默认[rbSpecifyBirthday]
	CalendarType 	string	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$CalendarType"`				//日历类型 默认【rbSolar】
	Year			string	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$ddlYear"`					//出生年
	Month 			string 	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$ddlMonth"`					//出生月
	Day 			string	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$ddlDay"`					//出生日
	Hour 			string 	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$ddlHour"`					//出生时 默认【0】
	Minute 			string 	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$ddlMinute"`				//出生分 默认【0】
	Country 		string 	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$tbCountry"`				//出生分 默认【中国】
	Province 		string 	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$tbProvince"`				//出生省 默认空
	City 			string 	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$tbCity"`					//出生省 默认空
	OtherHopes 		string 	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$tbOtherHopes"`				//默认空
	Career 			string 	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$ddlCareer"`				//默认【-2】
	Father 			string 	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$tbFather"`					//默认空
	Mother 			string 	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$tbMother"`					//默认空
	AvoidWords 		string 	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$tbAvoidWords"`				//默认空
	AvoidSimpParts 	string 	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$tbAvoidSimpParts"`			//默认空
	LoginUserName	string 	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$LoginAnywhere1$tbUserName"`//默认空
	LoginPwd 		string	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$LoginAnywhere1$tbPwd"`		//默认空
	LoginCode 		string	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$LoginAnywhere1$tbVCode"`	//默认空
	LoginParam		string	`json:"ctl00$ContentPlaceHolder1$InputBasicInfo1$LoginAnywhere1$loginParam"` //默认2
}

func BuildNumNameData(firstName string,suffixName string,gender string,year string,month string,day string) GetNumNameData{
	bean := GetNumNameData{}
	bean.FirstName = firstName
	bean.SuffixName = suffixName
	bean.Year = year
	bean.Month = month
	bean.Day = day
	bean.Genders = gender
	bean.Hour = "0"
	bean.Minute = "0"
	bean.Career = "-2"
	bean.Country = "中国"
	bean.LoginParam = "2"
	bean.BirthDayMode = "rbSpecifyBirthday"
	bean.CalendarType = "rbSolar"
	return bean
}

func GetNameIndexPage() GetNumPreDatas{
	cookies := make(map[string]string,0)
	inputs := make(map[string]string,0)

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.108 Safari/537.36"
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Host", "www.meimingteng.com")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Referer", "https://www.meimingteng.com")
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")

	})

	c.OnResponse(func(response *colly.Response) {
		setCookies := (*response.Headers)["Set-Cookie"]
		for _,setCookie := range setCookies {
			cookie := strings.SplitN(setCookie, "=",2)
			cookies[cookie[0]]=cookie[1]
		}
	})

	c.OnHTML("#aspnetForm", func(formElement *colly.HTMLElement) {
		formElement.ForEach("div > input", func(i int, inputElement *colly.HTMLElement) {
			inputKey := inputElement.Attr("name")
			if(inputKey == "ctl00$ContentPlaceHolder1$InputBasicInfo1$LoginAnywhere1$btImgAliLogin" ||
			 inputKey == "ctl00$ContentPlaceHolder1$InputBasicInfo1$LoginAnywhere1$btImgSinaLogin" ||
			 inputKey == "ctl00$ContentPlaceHolder1$InputBasicInfo1$LoginAnywhere1$btImgQQLogin"){
				return
			}
			inputValue := inputElement.Attr("value")
			inputs[inputKey] = inputValue
		})
	})


	c.Visit(REQUEST_NAME_NUMBER_URL)


	return GetNumPreDatas{Inputs:inputs,Cookies:cookies}
}


type NameNumberCal struct {
	ResultStatus 	string 		`json:"resultstatus"`
	ErrorInfo		string		`json:"errorinfo"`
	Scores 			[]string	`json:"numbers"`	//分数
	VerseArr 		[]string	`json:"verseArr"`	//诗词
	DetailArr 		[]string	`json:"detailArr"`	//详细
}

func GetNameNumber(preDatas GetNumPreDatas,getNumNameData GetNumNameData) NameNumberCal{
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.108 Safari/537.36"

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Origin", "https://www.meimingteng.com")
		r.Headers.Set("Content-Type","application/x-www-form-urlencoded; charset=UTF-8")
		r.Headers.Set("Host", "www.meimingteng.com")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
		r.Headers.Set("Referer", REQUEST_NAME_NUMBER_URL)
		r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
		r.Headers.Set("Content-Length", "6707")
		r.Headers.Set("Cache-Control", "no-cache")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
		cookie := "HELLO_USER=1;Hm_lvt_637e96da78d1c6c8f8a218c811dea5fb=1572836336; Hm_lpvt_637e96da78d1c6c8f8a218c811dea5fb=1572836336;"
		for key,value := range preDatas.Cookies {
			cookie = cookie + key+"="+strings.Split(value,"=")[0]+";"
		}
		cookie = strings.Replace(cookie,"domain;","",-1)
		r.Headers.Set("Cookie",cookie)

	})

	scores := make([]string,0) //分数
	fableText := ""	//寓意
	detailArr := make([]string,0) //详细
	verseArr := make([]string,0) //诗句
	var resp *colly.Response

	//分数
	c.OnHTML("#bdAppSummDiv > table:nth-child(7) > tbody > tr > td",func(e *colly.HTMLElement) {
		e.ForEach(".WordCategoryTitle", func(i int, wordTitle *colly.HTMLElement) {
			//scoreText = scoreText + fmt.Sprintf("%s	---		%s \n",wordTitle.Text,wordTitle.DOM.Next().Children().First().Text())
			scores = append(scores,wordTitle.DOM.Next().Children().First().Text())
		})
	})

	//寓意
	c.OnHTML("#ctl00_ContentPlaceHolder1_ShowNameDetails1_lbNameMeaning", func(nameMoralE *colly.HTMLElement) {
		fableText = fableText + fmt.Sprintf("寓意  %s ",nameMoralE.Text)
	})

	// 总评
	c.OnHTML("#ctl00_ContentPlaceHolder1_ShowNameDetails1_lbNLSummary", func(commentE *colly.HTMLElement) {
		charFormat := "%s  "
		detailText := ""
		for _,node := range commentE.DOM.Children().Nodes {
			if node.Data == "div" {
				continue
			}
			if node.FirstChild != nil && node.FirstChild.Data == "span" {
				if node.Attr[0].Val == "spanMeaningSentence"{
					detailText = detailText + commentE.ChildText("#spanMeaningSentence")
					/*meaningSentenceNode := node.FirstChild.FirstChild
					detailText = detailText + meaningSentenceNode.Data + meaningSentenceNode.NextSibling.FirstChild.Data +meaningSentenceNode.NextSibling.NextSibling.Data + meaningSentenceNode.NextSibling.NextSibling.NextSibling.FirstChild.Data


					if(meaningSentenceNode.NextSibling.NextSibling.NextSibling.NextSibling.Data == "font"){
						detailText = detailText + meaningSentenceNode.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.Data
					}else{
						detailText = detailText + meaningSentenceNode.NextSibling.NextSibling.NextSibling.NextSibling.Data
					}*/

				}
				if node.Attr[0].Val == "summaryIdiom"{
					detailText = detailText + commentE.ChildText("#summaryIdiom")
					/*changeyuNode := node.FirstChild.FirstChild

					detailText = detailText + changeyuNode.FirstChild.Data + changeyuNode.NextSibling.Data

					changeyuNode = node.FirstChild.NextSibling

					detailText = detailText + changeyuNode.NextSibling.FirstChild.Data + changeyuNode.NextSibling.FirstChild.NextSibling.FirstChild.Data*/
				}
			}

			if node.FirstChild != nil && node.FirstChild.Data != "span"{
				detailText = detailText + fmt.Sprintf(charFormat,node.FirstChild.Data)
				charFormat = "%s"
			}

			if node.NextSibling != nil && node.NextSibling.Data !="span" && node.NextSibling.Data !="div"{
				detailText = detailText + fmt.Sprintf(charFormat,node.NextSibling.Data)
				charFormat = "%s"
			}

			if node.Data == "br"{
				charFormat = "%s  "
				detailArr = append(detailArr,detailText)
				detailText = ""
			}
		}
	})


	//名人名句
	c.OnHTML("#ctl00_ContentPlaceHolder1_ShowNameDetails1_lbPoems", func(tableE *colly.HTMLElement) {
		verseText := ""
		tableE.ForEach("td > table", func(i int, tdE *colly.HTMLElement) {
			tdE.ForEach("td", func(j int, e *colly.HTMLElement) {
				verse := e.Text
				index := UnicodeIndex(e.Text,"。")
				unicodeLen, runes := UnicodeLen(e.Text)
				if index!=-1 && index != unicodeLen-1{
					verse = fmt.Sprintf("%s%s",string(runes[:index+1]),string(runes[index+1:]))
				}
				verseText = verseText + verse
				if  strings.Contains(verseText,"－"){
					verseArr = append(verseArr,verseText)
					verseText = ""
				}
			})
		})
	})
	c.OnResponse(func(response *colly.Response) {
		resp = response
	})
	postDatas := make(map[string]string)
	marshal, _ := json.Marshal(getNumNameData)
	json.Unmarshal(marshal,&postDatas)
	for key,value := range preDatas.Inputs {
		postDatas[key] = value
	}
	postDatas["__EVENTTARGET"]="ctl00$ContentPlaceHolder1$InputBasicInfo1$btNext"
	c.Post(REQUEST_NAME_NUMBER_URL,postDatas)



	return NameNumberCal{"1","",scores,verseArr,append(detailArr,fableText)}
}

func UnicodeIndex(str,substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str,substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}

	return result
}

func UnicodeLen(str string)(int,[]rune){
	s := []rune(string(str))
	return len(s),s
}