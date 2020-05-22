package server

import (
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

type Event struct {
	Event 			xml.Name			`xml:"Event"`
	Text			string				`xml:",cdata"`
}

type EventKey struct {
	EventKey 		xml.Name			`xml:"EventKey"`
	Text			string				`xml:",cdata"`
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
	articleItem.ItemDescription = ItemDescription{Text:fmt.Sprintf("启航人生,从名字开始。你的授权码为： %s",licenseCode)}
	articleItem.ItemPicUrl = ItemPicUrl{Text:"http://27953499sv.zicp.vip:25067/resources/wechart/qihang.png"}
	articleItem.ItemUrl = ItemUrl{Text:fmt.Sprintf("http://27953499sv.zicp.vip:25067/?lisence=%s",licenseCode)}

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


func (w WechartAPI) clickMenu(req *restful.Request, response *restful.Response){
	s, _ := ioutil.ReadAll(req.Request.Body)
	requestXml := RequestClickMenuXML{}
	xml.Unmarshal(s,&requestXml)
	eventKey := requestXml.EventKey.Text
	switch eventKey {
	case "clickA":
		w.clickMenuA(&requestXml,response)
	case "clickB":
		w.clickMenuB(&requestXml,response)

	default:
		panic("unrecognized value")
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
