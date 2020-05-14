package cnnumber

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net/url"
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
		//fmt.Printf("%s \n",response.Body)
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

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type","application/x-www-form-urlencoded")
		r.Headers.Set("cache-control", "no-cache")
		//cookie := "HELLO_USER=1;Hm_lvt_637e96da78d1c6c8f8a218c811dea5fb=1572836336; Hm_lpvt_637e96da78d1c6c8f8a218c811dea5fb=1572836336;"
		cookie := `Params=%26Xing%3d%e4%bd%95%26Gender%3d2%26Year%3d2020%26Month%3d4%26Day%3d27%26Hour%3d0%26Minute%3d0%26IsSolarCalendar%3d1%26IsLeapMonth%3d0%26NameType%3d1%26ReiterativeLocution%3d0%26Location%3d%e4%b8%ad%e5%9b%bd++%26Career%3d-2%26Personality%3d%26Father%3d%26Mother%3d%26SpecifiedName%3d%26SpecifiedNameIndex%3d0%26OtherHopes%3d%26AvoidWords%3d%26AvoidSimpParts%3d%26SpecifiedMing1SimpParts%3d%26SpecifiedMing2SimpParts%3d%26SpecifiedMing1Stroke%3d%26SpecifiedMing2Stroke%3d%26Tag%3d4%7c2%26LinChanQi%3dFalse%26NamingByCategoryCategoryID%3d-1%26SM1S%3d%26SM2S%3d%26SM1T%3d%26SM2T%3d%26SM1M%3d%26SM2M%3d%26RN%3d%26SpecifiedMing1Spell%3d%26SpecifiedMing2Spell%3d%26SM1Y%3d%26SM2Y%3d%26FA%3d%e5%87%8c%e6%99%a8++%e8%81%94%e8%b0%8a%e5%9f%8e%e6%97%a5++%e6%98%a5%e5%ad%a3++%e5%9b%9b%e6%9c%88%26LOCATION_COUNTY%3d%e4%b8%ad%e5%9b%bd%26LOCATION_PROVINCE%3d%26LOCATION_CITY%3d%26MING_WORDS%3d%e4%ba%ba; mmtsuser=1; HELLO_USER=1; ASP.NET_SessionId=komk32555sjkjr45skwedimv; ckcookie=chcookie; Hm_lvt_637e96da78d1c6c8f8a218c811dea5fb=1587980894,1588044567,1588337686; Hm_lpvt_637e96da78d1c6c8f8a218c811dea5fb=1588338448`

		/*for key,value := range preDatas.Cookies {
			cookie = cookie + key+"="+strings.Split(value,"=")[0]+";"
		}*/
		fmt.Printf("%s \n",cookie)
		r.Headers.Set("Cookie",cookie)

	})

	scores := make([]string,0) //分数
	fableText := ""	//寓意
	detailArr := make([]string,0) //详细
	verseArr := make([]string,0) //诗句

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

	/*c.OnResponse(func(response *colly.Response){
		fmt.Printf("%s",response.Body)
	})*/
	postDatas := make(map[string]string)
	marshal, _ := json.Marshal(getNumNameData)
	json.Unmarshal(marshal,&postDatas)
	for key,value := range preDatas.Inputs {
		postDatas[key] = value
	}
	postDatas["__EVENTTARGET"]="ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24btNext"
	postDatas["__VIEWSTATE"]=url.QueryEscape(postDatas["__VIEWSTATE"])
	postDatas["__EVENTVALIDATION"]=url.QueryEscape(postDatas["__EVENTVALIDATION"])
	datas := "1=1"
	for key,value := range postDatas {
		datas = fmt.Sprintf("%s&%s=%s",datas,key,value)
	}

	//c.PostRaw(REQUEST_NAME_NUMBER_URL,[]byte("__EVENTTARGET=ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24btNext&__EVENTARGUMENT=&__VIEWSTATE=%2FwEPDwULLTEyNjU5OTUwOTBkGAEFHl9fQ29udHJvbHNSZXF1aXJlUG9zdEJhY2tLZXlfXxYeBTtjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRyYlNwZWNpZnlCaXJ0aGRheQU9Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkcmJTcGVjaWZ5TGluQ2hhblFpbgU9Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkcmJTcGVjaWZ5TGluQ2hhblFpbgU%2BY3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkcmJOb3RTcGVjaWZ5QmlydGhkYXkFPmN0bDAwJENvbnRlbnRQbGFjZUhvbGRlcjEkSW5wdXRCYXNpY0luZm8xJHJiTm90U3BlY2lmeUJpcnRoZGF5BTFjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRyYlNvbGFyBTFjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRyYkx1bmFyBTFjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRyYkx1bmFyBTdjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRjYklzTGVhcE1vbnRoBTFjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRsYnROb25lBTFjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRsYnROb25lBTJjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRyYnRMdW5ZdQUyY3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkcmJ0THVuWXUFNGN0bDAwJENvbnRlbnRQbGFjZUhvbGRlcjEkSW5wdXRCYXNpY0luZm8xJHJidFNoaUppbmcFNGN0bDAwJENvbnRlbnRQbGFjZUhvbGRlcjEkSW5wdXRCYXNpY0luZm8xJHJidFNoaUppbmcFMWN0bDAwJENvbnRlbnRQbGFjZUhvbGRlcjEkSW5wdXRCYXNpY0luZm8xJHJidFBvZW0FMWN0bDAwJENvbnRlbnRQbGFjZUhvbGRlcjEkSW5wdXRCYXNpY0luZm8xJHJidFBvZW0FMmN0bDAwJENvbnRlbnRQbGFjZUhvbGRlcjEkSW5wdXRCYXNpY0luZm8xJHJidElkaW9tBTJjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRyYnRJZGlvbQU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkMAU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkMQU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkMgU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkMwU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkNAU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkNQU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkNgU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkNwU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkOAU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkOQU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkOc%2BRBppac3%2FCXaY8AJwjzwaxBvxk&__VIEWSTATEGENERATOR=9F5AD4C7&__EVENTVALIDATION=%2FwEWvQICkof8EwK19uvYAgKlkY7lBAKBmOYKAoCY5goCqcPP0AcCqMPP0AcCq8PP0AcCnd3BDAKVuaPyDgKV99WYCAL9kPD0BgKmsv%2F0AQKL2c%2BpBQKMlpPbCAKMloemAQKMluuBCgKMlt%2FsAgKMlsO3CwKMlreTBAKMlpv%2BDAKMls%2BWCgKMlrPyAgLhr8HqBQLhr7W2DgLhr5mRBwLhr438DwLhr%2FHHCALhr%2BWiAQLhr8mNCgLhr73pAgLhr%2BEBAuGv1ewIAvq448ULAvq416AEAvq4u4wNAvq4r9cFAvq4k7IOAvq4h50HAvq46%2FgPAvq438MIAvq4g%2FwFAvq498cOAt%2FRhLABAt%2FR6JsKAt%2FR3OYCAt%2FRwMELAt%2FRtK0EAt%2FRmIgNAt%2FRjNMFAt%2FR8L4OAt%2FRpNcLAt%2FRiLIEArDrpqsHArDrivYPArDr%2FtEIArDr4rwBArDr1ocKArDruuMCArDrrs4LArDrkqkEArDrxsEBArDrqq0KApWEuIYNApWErOEFApWEkMwOApWEhJcHApWE6PIPApWE3N0IApWEwLgBApWEtIQKApWE2LwHApWEzAcC7p3a8AIC7p3O2wsC7p2ypwQC7p2mgg0C7p2K7QUC7p3%2ByA4C7p3ikwcC7p3W%2Fg8C7p36lw0C7p3u8gUCw7b86wgCw7bgtgECw7bUkQoCw7a4%2FQICw7as2AsCw7aQowQCw7aEjg0Cw7bo6QUCw7acggMCw7aA7QsC9Ny8qgEC9Nyg9QkC9NyU0AIC9Nz4uwsC9NzshgQC9NzQ4QwC9NzEzAUC9NyoqA4C9NzcwAsC9NzAqwQCyfXehAcCyfXC7w8CyfW2ywgCyfWalgECyfWO8QkCyfXy3AICyfXmpwsCyfXKggQCyfX%2BuwECyfXihgoCjZaL8A8CjZb%2F2wgCjZbjpgECjZbXgQoCjZa77QICjZavyAsCjZaTkwQCjZaH%2FgwCjZarlwoCjZaf8gIC5q%2Bt6wUC5q%2BRtg4C5q%2BFkQcC5q%2Fp%2FA8C5q%2FdxwgC5q%2FBogEC5q%2B1jgoC5q%2BZ6QIC5q%2FNAQLmr7HtCAL7uM%2FFCwL7uLOhBAL7uKeMDQL7uIvXBQL7uP%2ByDgLyr4XtDALzr4XtDALwr4XtDALxr4XtDAL2r4XtDAL3r4XtDAL0r4XtDALlr4XtDALqr4XtDALyr8XuDALyr8nuDALyr83uDAKVwsetDgKUwsetDgKXwsetDgKWwsetDgKRwsetDgKQwsetDgKTwsetDgKCwsetDgKNwsetDgKVwoeuDgKVwouuDgKVwo%2BuDgKVwrOuDgKVwreuDgKVwruuDgKVwr%2BuDgKVwqOuDgKVwuetDgKVwuutDgKUwoeuDgKUwouuDgKUwo%2BuDgKUwrOuDgKUwreuDgKUwruuDgKUwr%2BuDgKUwqOuDgKUwuetDgKUwuutDgKXwoeuDgKXwouuDgKArtmFCAKTrpWGCAKMrpWGCAKNrpWGCAKOrpWGCAKPrpWGCAKIrpWGCAKJrpWGCAKKrpWGCAKbrpWGCAKUrpWGCAKMrtWFCAKMrtmFCAKMrt2FCAKMruGFCAKMruWFCAKMrumFCAKMru2FCAKMrvGFCAKMrrWGCAKMrrmGCAKNrtWFCAKNrtmFCAKNrt2FCAKNruGFCALCsOLeDALdsOLeDALcsOLeDALfsOLeDALesOLeDALZsOLeDALYsOLeDALbsOLeDALKsOLeDALFsOLeDALdsKLdDALdsK7dDALdsKrdDALdsJbdDALdsJLdDALdsJ7dDALdsJrdDALdsIbdDALdsMLeDALdsM7eDALcsKLdDALcsK7dDALcsKrdDALcsJbdDALcsJLdDALcsJ7dDALcsJrdDALcsIbdDALcsMLeDALcsM7eDALfsKLdDALfsK7dDALfsKrdDALfsJbdDALfsJLdDALfsJ7dDALfsJrdDALfsIbdDALfsMLeDALfsM7eDALesKLdDALesK7dDALesKrdDALesJbdDALesJLdDALesJ7dDALesJrdDALesIbdDALesMLeDALesM7eDALZsKLdDALZsK7dDALZsKrdDALZsJbdDALZsJLdDALZsJ7dDALZsJrdDALZsIbdDALZsMLeDALZsM7eDALq3rerBALCuMjbDwKi2vn0BQK1gMHLAgKy5bvtBQK236G8AQLV9%2FnaAgKgqtmeCgLVnOXeBwL62evxBgL22aPyBgL32aPyBgL02aPyBgL12aPyBgLy2aPyBgLz2aPyBgLw2aPyBgLh2aPyBgLu2aPyBgL22ePxBgL22e%2FxBgL22evxBgL22dfxBgL22dPxBgL22dvxBgL22cfxBgL62e%2FxBgL%2Fl8bNAwL%2Fl8rNAwL%2Fl77NAwL%2Fl8LNAwL%2Fl9bNAwL%2Fl9rNAwL%2Fl87NAwL%2Fl9LNAwL%2Fl6bNAwL%2Fl6rNAwL9lKKdCgKWk7qbCgL1g6aRDwLXw%2FfkCALvyL%2FmDALev%2FOZDgLL5ZPJCAKi%2FN2EAQKD5eu3CwKgl5eyCgKeoZaXAwKd8qmzBALwr6%2FEDwLS6rKCB%2BZNztJhPsAvOO8hoz19BQD7mwTF&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24tbXing=%E4%BD%95&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24tbMingWords=%E4%BA%BA&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24ddlGenders=2&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24SPECIFY_BIRHDAY=rbSpecifyBirthday&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24CalendarType=rbSolar&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24ddlYear=2020&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24ddlMonth=4&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24ddlDay=27&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24ddlHour=0&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24ddlMinute=0&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24tbCountry=%E4%B8%AD%E5%9B%BD&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24tbProvince=&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24tbCity=&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24tbOtherHopes=&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24ddlCareer=-2&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24tbFather=&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24tbMother=&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24tbAvoidWords=&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24tbAvoidSimpParts=&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24LoginAnywhere1%24tbUserName=&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24LoginAnywhere1%24tbPwd=&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24LoginAnywhere1%24tbVCode=&ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24LoginAnywhere1%24loginParam=2"))
	c.PostRaw(REQUEST_NAME_NUMBER_URL,[]byte(datas))

	//已失效
	//c.Post(REQUEST_NAME_NUMBER_URL,postDatas)



	return NameNumberCal{"1","",scores,verseArr,append(detailArr,fableText)}
}

// 测试参数
func testFormData() map[string]string{

	data := make(map[string]string)
	data["__EVENTTARGET"] = "ctl00%24ContentPlaceHolder1%24InputBasicInfo1%24btNext"
	data["__EVENTARGUMENT"] = ""
	data["__VIEWSTATE"] = url.QueryEscape("/wEPDwULLTEyNjU5OTUwOTBkGAEFHl9fQ29udHJvbHNSZXF1aXJlUG9zdEJhY2tLZXlfXxYeBTtjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRyYlNwZWNpZnlCaXJ0aGRheQU9Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkcmJTcGVjaWZ5TGluQ2hhblFpbgU9Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkcmJTcGVjaWZ5TGluQ2hhblFpbgU+Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkcmJOb3RTcGVjaWZ5QmlydGhkYXkFPmN0bDAwJENvbnRlbnRQbGFjZUhvbGRlcjEkSW5wdXRCYXNpY0luZm8xJHJiTm90U3BlY2lmeUJpcnRoZGF5BTFjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRyYlNvbGFyBTFjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRyYkx1bmFyBTFjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRyYkx1bmFyBTdjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRjYklzTGVhcE1vbnRoBTFjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRsYnROb25lBTFjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRsYnROb25lBTJjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRyYnRMdW5ZdQUyY3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkcmJ0THVuWXUFNGN0bDAwJENvbnRlbnRQbGFjZUhvbGRlcjEkSW5wdXRCYXNpY0luZm8xJHJidFNoaUppbmcFNGN0bDAwJENvbnRlbnRQbGFjZUhvbGRlcjEkSW5wdXRCYXNpY0luZm8xJHJidFNoaUppbmcFMWN0bDAwJENvbnRlbnRQbGFjZUhvbGRlcjEkSW5wdXRCYXNpY0luZm8xJHJidFBvZW0FMWN0bDAwJENvbnRlbnRQbGFjZUhvbGRlcjEkSW5wdXRCYXNpY0luZm8xJHJidFBvZW0FMmN0bDAwJENvbnRlbnRQbGFjZUhvbGRlcjEkSW5wdXRCYXNpY0luZm8xJHJidElkaW9tBTJjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRyYnRJZGlvbQU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkMAU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkMQU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkMgU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkMwU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkNAU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkNQU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkNgU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkNwU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkOAU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkOQU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkOc+RBppac3/CXaY8AJwjzwaxBvxk")
	data["__VIEWSTATEGENERATOR"] = "9F5AD4C7"
	data["__EVENTVALIDATION"] = url.QueryEscape("/wEWvQICkof8EwK19uvYAgKlkY7lBAKBmOYKAoCY5goCqcPP0AcCqMPP0AcCq8PP0AcCnd3BDAKVuaPyDgKV99WYCAL9kPD0BgKmsv/0AQKL2c+pBQKMlpPbCAKMloemAQKMluuBCgKMlt/sAgKMlsO3CwKMlreTBAKMlpv+DAKMls+WCgKMlrPyAgLhr8HqBQLhr7W2DgLhr5mRBwLhr438DwLhr/HHCALhr+WiAQLhr8mNCgLhr73pAgLhr+EBAuGv1ewIAvq448ULAvq416AEAvq4u4wNAvq4r9cFAvq4k7IOAvq4h50HAvq46/gPAvq438MIAvq4g/wFAvq498cOAt/RhLABAt/R6JsKAt/R3OYCAt/RwMELAt/RtK0EAt/RmIgNAt/RjNMFAt/R8L4OAt/RpNcLAt/RiLIEArDrpqsHArDrivYPArDr/tEIArDr4rwBArDr1ocKArDruuMCArDrrs4LArDrkqkEArDrxsEBArDrqq0KApWEuIYNApWErOEFApWEkMwOApWEhJcHApWE6PIPApWE3N0IApWEwLgBApWEtIQKApWE2LwHApWEzAcC7p3a8AIC7p3O2wsC7p2ypwQC7p2mgg0C7p2K7QUC7p3+yA4C7p3ikwcC7p3W/g8C7p36lw0C7p3u8gUCw7b86wgCw7bgtgECw7bUkQoCw7a4/QICw7as2AsCw7aQowQCw7aEjg0Cw7bo6QUCw7acggMCw7aA7QsC9Ny8qgEC9Nyg9QkC9NyU0AIC9Nz4uwsC9NzshgQC9NzQ4QwC9NzEzAUC9NyoqA4C9NzcwAsC9NzAqwQCyfXehAcCyfXC7w8CyfW2ywgCyfWalgECyfWO8QkCyfXy3AICyfXmpwsCyfXKggQCyfX+uwECyfXihgoCjZaL8A8CjZb/2wgCjZbjpgECjZbXgQoCjZa77QICjZavyAsCjZaTkwQCjZaH/gwCjZarlwoCjZaf8gIC5q+t6wUC5q+Rtg4C5q+FkQcC5q/p/A8C5q/dxwgC5q/BogEC5q+1jgoC5q+Z6QIC5q/NAQLmr7HtCAL7uM/FCwL7uLOhBAL7uKeMDQL7uIvXBQL7uP+yDgLyr4XtDALzr4XtDALwr4XtDALxr4XtDAL2r4XtDAL3r4XtDAL0r4XtDALlr4XtDALqr4XtDALyr8XuDALyr8nuDALyr83uDAKVwsetDgKUwsetDgKXwsetDgKWwsetDgKRwsetDgKQwsetDgKTwsetDgKCwsetDgKNwsetDgKVwoeuDgKVwouuDgKVwo+uDgKVwrOuDgKVwreuDgKVwruuDgKVwr+uDgKVwqOuDgKVwuetDgKVwuutDgKUwoeuDgKUwouuDgKUwo+uDgKUwrOuDgKUwreuDgKUwruuDgKUwr+uDgKUwqOuDgKUwuetDgKUwuutDgKXwoeuDgKXwouuDgKArtmFCAKTrpWGCAKMrpWGCAKNrpWGCAKOrpWGCAKPrpWGCAKIrpWGCAKJrpWGCAKKrpWGCAKbrpWGCAKUrpWGCAKMrtWFCAKMrtmFCAKMrt2FCAKMruGFCAKMruWFCAKMrumFCAKMru2FCAKMrvGFCAKMrrWGCAKMrrmGCAKNrtWFCAKNrtmFCAKNrt2FCAKNruGFCALCsOLeDALdsOLeDALcsOLeDALfsOLeDALesOLeDALZsOLeDALYsOLeDALbsOLeDALKsOLeDALFsOLeDALdsKLdDALdsK7dDALdsKrdDALdsJbdDALdsJLdDALdsJ7dDALdsJrdDALdsIbdDALdsMLeDALdsM7eDALcsKLdDALcsK7dDALcsKrdDALcsJbdDALcsJLdDALcsJ7dDALcsJrdDALcsIbdDALcsMLeDALcsM7eDALfsKLdDALfsK7dDALfsKrdDALfsJbdDALfsJLdDALfsJ7dDALfsJrdDALfsIbdDALfsMLeDALfsM7eDALesKLdDALesK7dDALesKrdDALesJbdDALesJLdDALesJ7dDALesJrdDALesIbdDALesMLeDALesM7eDALZsKLdDALZsK7dDALZsKrdDALZsJbdDALZsJLdDALZsJ7dDALZsJrdDALZsIbdDALZsMLeDALZsM7eDALq3rerBALCuMjbDwKi2vn0BQK1gMHLAgKy5bvtBQK236G8AQLV9/naAgKgqtmeCgLVnOXeBwL62evxBgL22aPyBgL32aPyBgL02aPyBgL12aPyBgLy2aPyBgLz2aPyBgLw2aPyBgLh2aPyBgLu2aPyBgL22ePxBgL22e/xBgL22evxBgL22dfxBgL22dPxBgL22dvxBgL22cfxBgL62e/xBgL/l8bNAwL/l8rNAwL/l77NAwL/l8LNAwL/l9bNAwL/l9rNAwL/l87NAwL/l9LNAwL/l6bNAwL/l6rNAwL9lKKdCgKWk7qbCgL1g6aRDwLXw/fkCALvyL/mDALev/OZDgLL5ZPJCAKi/N2EAQKD5eu3CwKgl5eyCgKeoZaXAwKd8qmzBALwr6/EDwLS6rKCB+ZNztJhPsAvOO8hoz19BQD7mwTF")
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$tbXing"] = "何"
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$tbMingWords"] = "人"
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$ddlGenders"] = "2"
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$SPECIFY_BIRHDAY"] = "rbSpecifyBirthday"
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$CalendarType"] = "rbSolar"
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$ddlYear"] = "2020"
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$ddlMonth"] = "4"
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$ddlDay"] = "27"
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$ddlHour"] = "0"
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$ddlMinute"] = "0"
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$tbCountry"] = "中国"
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$tbProvince"] = ""
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$tbCity"] = ""
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$tbOtherHopes"] = ""
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$ddlCareer"] = "-2"
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$tbFather"] = ""
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$tbMother"] = ""
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$tbAvoidWords"] = ""
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$tbAvoidSimpParts"] = ""
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$LoginAnywhere1$tbUserName"] = ""
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$LoginAnywhere1$tbPwd"] = ""
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$LoginAnywhere1$tbVCode"] = ""
	data["ctl00$ContentPlaceHolder1$InputBasicInfo1$LoginAnywhere1$loginParam"] = "2"
	return data
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