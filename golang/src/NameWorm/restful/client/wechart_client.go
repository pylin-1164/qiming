package client

import (
	config "NameWorm/common"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

func init() {
	client := ApiWechartClient{
		"https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=${APPID}&secret=${APPSECRET}",
		"https://api.weixin.qq.com/cgi-bin/menu/create?access_token=${ACCESS_TOKEN}"}
	client.getAccessToken()
/*	订阅号不支持创建菜单功能
	client.createMenu()
	client.uploadImage()*/

}

var  accessToken AccessTokenResponse
var  mediaCache map[string]string = make(map[string]string)

const (
	MenuButtonType_CLICK = "click"
	MenuButtonType_VIEW = "view"
	/** 花生壳域名 */
	SERVER_HOST = "http://27953499sv.zicp.vip"
)

type ApiWechartClient struct {

	GetAccessTokenUrl 	string
	CreateMenuUrl 		string
}

type AccessTokenResponse struct {
	AccessToken 	string		`json:"access_token"`
	Expires 		string		`json:"expires_in"`
}

type CreateMenu struct {
	Button 		[]MenuButton		`json:"button"`
}

type MenuButton struct {
	Name		string				`json:"name,omitempty"`
	Type 		string 				`json:"type,omitempty"`
	Key 		string				`json:"key,omitempty"`
	Url			string				`json:"url,omitempty"`
	SubButton	[]MenuButton 		`json:"sub_button,omitempty"`
}

func (w ApiWechartClient) getAccessToken(){

	appid := config.GetWechartAppID()
	appsecret := config.GetWechartSecret()
	url := w.GetAccessTokenUrl
	url = strings.Replace(url,"${APPID}",appid,1)
	url = strings.Replace(url,"${APPSECRET}",appsecret,1)
	client := resty.New()
	response, err := client.R().Get(url)
	if err != nil {
		panic("GET ACCESS_TOKEN PANIC ... ")
	}
	if response.StatusCode() == http.StatusOK {
		tokenResponse := AccessTokenResponse{}
		fmt.Printf("%s \n",response.Body())
		json.Unmarshal(response.Body(),&tokenResponse)
		accessToken = tokenResponse
	}else {
		panic("GET ACCESS_TOKEN RESPONSE CODE " + fmt.Sprint("%d",response.StatusCode()))
	}
}

func (w ApiWechartClient) createMenu(){
	url := w.CreateMenuUrl
	url = strings.Replace(url,"${ACCESS_TOKEN}",accessToken.AccessToken,1)

	buttonA := MenuButton{Name:"起名字", Type:MenuButtonType_CLICK, Key:"clickA", Url:"",SubButton:nil }
	buttonB := MenuButton{ Name:"加入宝妈圈", Type:MenuButtonType_CLICK, Key:"clickB", Url:"",SubButton:nil}
	buttonC := MenuButton{ Name:"帮助", Type:MenuButtonType_CLICK, Key:"clickC", Url:"",SubButton:nil}

	//menuBtnC1 := MenuButton{ Name:"帮助",Type: MenuButtonType_CLICK,Key: "click_c_1", Url:"",SubButton:nil}
	//menuBtnC2 := MenuButton{ Name:"菜单C-2",Type: MenuButtonType_VIEW,Key: "", Url:"http://www.soso.com/",SubButton:nil}
	//menuButtonsC := append([]MenuButton{}, menuBtnC1, menuBtnC2)
	//buttonC := MenuButton{Name:"帮助",Type: MenuButtonType_CLICK, Key:"clickC", Url:SERVER_HOST+"/api/wechart/menuC",SubButton:menuButtonsC}

	buttons := append([]MenuButton{}, buttonA, buttonB, buttonC)
	createMenu := CreateMenu{buttons}

	menuBody, _ := json.Marshal(createMenu)
	fmt.Println(string(menuBody))
	response, err := resty.New().R().SetBody(menuBody).Post(url)
	if err != nil {
		panic("POST CREATE MENU PANIC ... ")
	}
	if response.StatusCode() == 200 {
		fmt.Println("POST CREATE MENU OK")
		fmt.Printf("%s \n",response.Body())
	}else{
		panic("POST CREATE MENU RESPONSE STATUS " + fmt.Sprint("%d",response.StatusCode()))
	}

}

type MediaResponse struct {
	Type 		string			`json:"type"`
	MediaId 	string			`json:"media_id"`
	CreateAt 	string 			`json:"create_at"`
}


func (w ApiWechartClient) uploadImage(){
	url := "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=${ACCESS_TOKEN}&type=image"
	url = strings.Replace(url,"${ACCESS_TOKEN}",accessToken.AccessToken,1)
	filePath := "/data/wechart/img/qrcode.png";
	if runtime.GOOS == "windows"{
		filePath = "D:/file_upload/lz/workInfo/2020/01/08/微信图片8232557-115603.jpg";
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	fileName := "微信图片8232557-115603.jpg"
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).R()
	response, err := client.SetFileReader("media", fileName, bytes.NewReader(file)).Post(url)
	if err == nil {
		fmt.Println(string(response.Body()))
		mediaResponse := MediaResponse{}
		json.Unmarshal(response.Body(),&mediaResponse)
		mediaCache["QR-CODE"] = mediaResponse.MediaId
	}

}

func GetQRMeidaId() string{
	return mediaCache["QR-CODE"]
}