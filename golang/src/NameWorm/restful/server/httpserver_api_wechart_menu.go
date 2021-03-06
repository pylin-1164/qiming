package server

import (
	config "NameWorm/common"
	"NameWorm/restful/client"
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"github.com/emicklei/go-restful"
	"io/ioutil"
	"math/rand"
	"sort"
	"strings"
	"time"
)

//授权码有效期
var LicenseCache = make(map[string]int64)

const ErrorMSG  = "抱歉，未检索到相关信息。\r ^_^ 您可以回复\"名字\"获取授权访问链接"
type WechartAPI struct {

}


func init() {
	WechartAPI{}.RegistRoute(WS)
}


func (w WechartAPI) RegistRoute(server *restful.WebService) {

	/**微信接口Token校验*/
	server.Route(server.GET("/api/wechart/token").To(w.wechartToken).Doc("get wechart token"))
	fmt.Printf("listener join post api : %s \n","/api/wechart/token")

	server.Route(server.POST("/api/wechart/token").To(w.clickMenu).Doc("POST Click Button"))
	fmt.Printf("listener join post api : %s \n","/api/wechart/menuA")

}


/**请求报文*/
type RequestClickMenuXML struct {
	XMLName			xml.Name 		`xml:"xml"`
	ToUserName		ToUserName		`xml:"ToUserName"`
	FromUserName	FromUserName	`xml:"FromUserName"`
	CreateTime		string			`xml:"CreateTime"`
	MsgType			MsgType			`xml:"MsgType"`
	Event			Event			`xml:"Event"`
	EventKey		EventKey		`xml:"EventKey"`
	Content			string			`xml:"Content"`
	MsgId			string			`xml:"MsgId"`

}

type ToUserName struct {
	XMLName			xml.Name			`xml:"ToUserName"`
	Text 			string				`xml:",cdata"`
}

type FromUserName struct {
	XMLName 		xml.Name			`xml:"FromUserName"`
	Text 			string				`xml:",cdata"`
}

type MsgType struct {
	XMLName 		xml.Name			`xml:"MsgType"`
	Text			string				`xml:",cdata"`
}

type Content struct {
	XMLName 		xml.Name			`xml:"Content"`
	Text			string				`xml:",cdata"`
}

type Event struct {
	Event 			xml.Name			`xml:"Event"`
	Text			string				`xml:",cdata"`
}

type EventKey struct {
	EventKey 		xml.Name			`xml:"EventKey"`
	Text			string				`xml:",cdata"`
}


type ResponseTextXML struct {
	XMLName			xml.Name 		`xml:"xml"`
	ToUserName		ToUserName		`xml:"ToUserName"`
	FromUserName	FromUserName	`xml:"FromUserName"`
	CreateTime		string			`xml:"CreateTime"`
	MsgType			MsgType			`xml:"MsgType"`
	Content 		Content 		`xml:"Content"`
}


type ResponseNewsXML struct {
	XMLName			xml.Name 		`xml:"xml"`
	ToUserName		ToUserName		`xml:"ToUserName"`
	FromUserName	FromUserName	`xml:"FromUserName"`
	CreateTime		string			`xml:"CreateTime"`
	MsgType			MsgType			`xml:"MsgType"`
	ArticleCount	int				`xml:"ArticleCount,omitempty"`
	Articles		*Articles		`xml:"Articles,omitempty"`
	Image			*Image			`xml:"Image,omitempty"`
}


type ArticleCount struct {
	XMLName 	xml.Name				`xml:"ArticleCount"`
	Text			string				`xml:",cdata"`
}
type Articles struct {
	XMLName 		xml.Name				`xml:"Articles"`
	Items			[]ArticleItem			`xml:"item"`
}

type Image struct {
	XMLName 		xml.Name 			`xml:"Image"`
	MediaId			MediaId				`xml:"MediaId"`
}

type MediaId struct {
	XMLName 		xml.Name			`xml:"MediaId"`
	Text 			string				`xml:",cdata"`
}

type ArticleItem struct {
	XMLName				xml.Name			`xml:"item"`
	ItemTitle 			ItemTitle 			`xml:"Title,omitempty"`
	ItemDescription		ItemDescription		`xml:"Description,omitempty"`
	ItemPicUrl			ItemPicUrl 			`xml:"PicUrl,omitempty"`
	ItemUrl				ItemUrl 			`xml:"Url,omitempty"`
}


type ItemTitle struct {
	XMLName 		xml.Name			`xml:"Title"`
	Text  			string				`xml:",cdata"`
}
type ItemDescription struct {
	XMLName 	xml.Name			`xml:"Description"`
	Text  			string				`xml:",cdata"`
}
type ItemPicUrl struct {
	XMLName 			xml.Name			`xml:"PicUrl"`
	Text  				string				`xml:",cdata"`
}

type ItemUrl struct {
	XMLName 			xml.Name			`xml:"Url"`
	Text  			string				`xml:",cdata"`
}

/**起名字，提供token*/
func (w WechartAPI) clickMenuA(requestXml *RequestClickMenuXML, response *restful.Response){

	responseNewsXML := ResponseNewsXML{}
	responseNewsXML.FromUserName = FromUserName{Text:requestXml.ToUserName.Text}
	responseNewsXML.ToUserName = ToUserName{Text:requestXml.FromUserName.Text}
	responseNewsXML.ArticleCount = 1
	responseNewsXML.CreateTime = fmt.Sprintf("%d",time.Now().Unix())
	responseNewsXML.MsgType = MsgType{Text:"news"}

	licenseCode := fmt.Sprintf("%d",1000+rand.Intn(9000))
	expire := time.Now().Unix()+30*60
	LicenseCache[licenseCode]=expire

	articleItem := ArticleItem{}
	articleItem.ItemTitle = ItemTitle{Text:"启名"}
	articleItem.ItemDescription = ItemDescription{Text:fmt.Sprintf("以名之始，启航人生。你的授权码为： %s",licenseCode)}
	articleItem.ItemPicUrl = ItemPicUrl{Text:fmt.Sprintf("%s/resources/wechart/qihang.png",config.UI_SERVER["host"])}
	articleItem.ItemUrl = ItemUrl{Text:fmt.Sprintf("%s/?lisence=%s",config.UI_SERVER["host"],licenseCode)}

	//qrItem := ArticleItem{}
	//qrItem.ItemTitle = ItemTitle{Text:"关注孩子"}
	//qrItem.ItemDescription = ItemDescription{Text:"加入我们，一起交流孩子成长的烦恼"}
	//qrItem.ItemPicUrl = ItemPicUrl{Text:"http://27953499sv.zicp.vip:25067/resources/wechart/qrcode.png"}
	//qrItem.ItemUrl = ItemUrl{Text:"http://27953499sv.zicp.vip:25067/resources/wechart/qrcode.png"}
	responseNewsXML.Articles = &Articles{Items:[]ArticleItem{articleItem}}


	responseBody, _ := xml.Marshal(responseNewsXML)
	response.Write(responseBody)
}

/**加入朋友圈*/
func (w WechartAPI) clickMenuB(requestXml *RequestClickMenuXML, response *restful.Response){
	responseNewsXML := ResponseNewsXML{}
	responseNewsXML.FromUserName = FromUserName{Text:requestXml.ToUserName.Text}
	responseNewsXML.ToUserName = ToUserName{Text:requestXml.FromUserName.Text}
	responseNewsXML.CreateTime = fmt.Sprintf("%d",time.Now().Unix())
	responseNewsXML.MsgType = MsgType{Text:"image"}
	responseNewsXML.Articles = nil
	mediaId := MediaId{Text: client.GetQRMeidaId()}
	responseNewsXML.Image = &Image{MediaId:mediaId}
	responseBody, _ := xml.Marshal(responseNewsXML)
	fmt.Println(string(responseBody))
	response.Write(responseBody)
}

func (w WechartAPI) errorMsg(requestXml *RequestClickMenuXML, response *restful.Response){
	responseTextXML := ResponseTextXML{}
	responseTextXML.FromUserName = FromUserName{Text:requestXml.ToUserName.Text}
	responseTextXML.ToUserName = ToUserName{Text:requestXml.FromUserName.Text}
	responseTextXML.CreateTime = fmt.Sprintf("%d",time.Now().Unix())
	responseTextXML.MsgType = MsgType{Text:"text"}
	responseTextXML.Content =  Content{Text:ErrorMSG}
	responseBody, _ := xml.Marshal(responseTextXML)
	fmt.Println(string(responseBody))
	response.Write(responseBody)
}


func (w WechartAPI) clickMenu(req *restful.Request, response *restful.Response){
	s, _ := ioutil.ReadAll(req.Request.Body)
	requestXml := RequestClickMenuXML{}
	xml.Unmarshal(s,&requestXml)
	eventKey := requestXml.EventKey.Text
	content := requestXml.Content
	if "名字" == strings.Trim(content," ")  {
		eventKey = "msg"
	}
	switch eventKey {
	case "clickA":
		w.clickMenuA(&requestXml,response)
		break
	case "clickB":
		w.clickMenuB(&requestXml,response)
		break
	case "msg" :
		w.clickMenuA(&requestXml,response)
		break
	default:
		w.errorMsg(&requestXml,response)

	}

}

/***
 * 微信Token校验
 */
func (w WechartAPI) wechartToken(req *restful.Request, response *restful.Response) {
	signature := req.Request.FormValue("signature")
	timestamp := req.Request.FormValue("timestamp")
	nonce := req.Request.FormValue("nonce")
	echostr := req.Request.FormValue("echostr")
	token := "pyl1164"
	mytokenarr := append([]string{},token,timestamp,nonce)
	sort.Strings(mytokenarr)
	mysign := sha1Encrypt(strings.Join(mytokenarr, ""))
	if signature == mysign {

		response.Write([]byte(echostr))
	}else{
		response.Write([]byte("false"))
	}

}

func sha1Encrypt(s string) string{
	h := sha1.New() // md5加密类似md5.New()
	//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(s))
	//这个用来得到最终的散列值的字符切片。Sum 的参数可以用来对现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)
	//SHA1 值经常以 16 进制输出，使用%x 来将散列结果格式化为 16 进制字符串。
	return fmt.Sprintf("%x", bs)

}
