package cnname

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type SuffixName struct {
	Sex		int	//性别
	Suffix 	string	//名
}

func DrillDownSexNames(firstData FirstNameData,sex int,sexpath string) func()([]SuffixName,error){
	suffixNames := make([]SuffixName, 0)
	nextUrl := sexpath
	return func() ([]SuffixName,error){
		c := colly.NewCollector()
		c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.108 Safari/537.36"
		c.OnRequest(func(r *colly.Request) {
			r.Headers.Set("Host", firstData.Path)
			r.Headers.Set("Connection", "keep-alive")
			r.Headers.Set("Accept", "*/*")
			r.Headers.Set("Origin", "http://www.resgain.net")
			r.Headers.Set("Referer", firstData.Path)
			r.Headers.Set("Accept-Encoding", "gzip, deflate")
			r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
		})

		c.OnHTML("body > div.main_ > div:nth-child(3) > div:nth-child(1) > div",func(e *colly.HTMLElement) {
			e.ForEach(".namelist", func(i int, namelist *colly.HTMLElement) {
				fullname := namelist.ChildText("div:first-child")
				suffixName := SuffixName{Sex: sex, Suffix: strings.Replace(fullname, firstData.First, "", 1)}
				suffixNames = append(suffixNames,suffixName)
			})
		})
		//TODO 可能没有.pagination属性 [http://ben.resgain.net/name/girls.html]
		c.OnHTML(".pagination", func(e *colly.HTMLElement) {
			e.ForEach("li > a", func(i int, alink *colly.HTMLElement) {
				if alink.Text == "下一页" {
					fmt.Println(fmt.Sprintf("%s%s",firstData.Path,alink.Attr("href")))
					nextUrl = fmt.Sprintf("%s%s",firstData.Path,alink.Attr("href"))
				}
				//TODO 只查询第一页
				//nextUrl = ""
			})
		})

		c.OnResponse(func(response *colly.Response) {
			nextUrl = ""
		})


		if nextUrl == "http://ben.resgain.net/name/boys_4.html"{
			fmt.Println("debuger...")
		}
		c.Visit(nextUrl)


		//延迟1s继续检索
		time.Sleep(time.Duration(1)*time.Second)

		if nextUrl == ""{
			return suffixNames,errors.New("End...")
		}
		return suffixNames,nil
	}
}

func DrillDownNames(firstData FirstNameData)[]SuffixName{
	boysPath := fmt.Sprintf("%s/name/boys.html", firstData.Path)
	girlsPath := fmt.Sprintf("%s/name/girls.html", firstData.Path)
	boyNamesFun := DrillDownSexNames(firstData, 1, boysPath)
	girlNamesFun := DrillDownSexNames(firstData, 2, girlsPath)
	var boyNames,grilNames []SuffixName
	var e error
	for {
		boyNames, e = boyNamesFun()
		if e != nil{
			break
		}
	}

	for {
		grilNames, e = girlNamesFun()
		if e != nil{
			break
		}
	}
	return append(boyNames,grilNames...)
}
