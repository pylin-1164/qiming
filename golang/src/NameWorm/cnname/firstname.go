package cnname

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

//TODO 爬取数据地址：http://www.resgain.net/xsdq_%s.html

type FirstName struct {
	Datas []FirstNameData
}

type FirstNameData struct {
	Path	string	//汇总地址
	First 	string	//姓
}

func (f *FirstName)	FindUrls(){
	f.Datas = make([]FirstNameData,0)
	//firstchar := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","k","y","z"}
	firstchar := []string{"b"}
	for _,u := range firstchar {
		firstPath := fmt.Sprintf("http://www.resgain.net/xsdq_%s.html", u)
		c := colly.NewCollector()
		c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.108 Safari/537.36"
		c.OnRequest(func(r *colly.Request) {
			r.Headers.Set("Host", "query.sse.com.cn")
			r.Headers.Set("Connection", "keep-alive")
			r.Headers.Set("Accept", "*/*")
			r.Headers.Set("Origin", "http://www.resgain.net")
			r.Headers.Set("Referer", "http://www.resgain.net") //关键头 如果没有 则返回 错误
			r.Headers.Set("Accept-Encoding", "gzip, deflate")
			r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
		})

		c.OnHTML("body > div.main_ > div > div > div:nth-child(2)",func(e *colly.HTMLElement) {
			e.ForEach("a", func(i int, a *colly.HTMLElement) {
				path := fmt.Sprintf("http:%s", a.Attr("href"))
				firstName := ""
				if strings.Contains(a.Text,"姓之家"){
					firstName = strings.Replace(a.Text,"姓之家","",1)
				}else{
					firstName = strings.Replace(a.Text,"之家","",1)
				}
				firstNameData := FirstNameData{Path: path, First: firstName}
				f.Datas = append(f.Datas,firstNameData)
			})
		})



		c.Visit(firstPath)
	}
}