package cnnumber

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func HttpPost(preDatas GetNumPreDatas,getNumNameData GetNumNameData){
	values := &url.Values{}
	postDatas := make(map[string]string)
	marshal, _ := json.Marshal(getNumNameData)
	json.Unmarshal(marshal,&postDatas)
	i := 0
	for key,value := range postDatas {
		values.Set(key,value)
		i++
		fmt.Printf("%d -> %s : %s \n",i,key,value)
	}
	for key,value := range preDatas.Inputs {
		values.Set(key,value)
		i++
		fmt.Printf("%d -> %s : %s \n",i,key,value)
	}
	//postDatas["__EVENTTARGET"]="ctl00$ContentPlaceHolder1$InputBasicInfo1$btNext"
	values.Set("__EVENTTARGET","ctl00$ContentPlaceHolder1$InputBasicInfo1$btNext")
	rawValues := values.Encode()
	r, err := http.NewRequest(http.MethodPost, REQUEST_NAME_NUMBER_URL, bytes.NewBufferString(rawValues))
	if err != nil {
		panic(err)
	}
	r.Header.Set("Origin", "https://www.meimingteng.com")
	r.Header.Set("Content-Type","application/x-www-form-urlencoded; charset=UTF-8")
	r.Header.Set("Host", "www.meimingteng.com")
	r.Header.Set("Connection", "keep-alive")
	r.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	r.Header.Set("Referer", REQUEST_NAME_NUMBER_URL)
	r.Header.Set("Accept-Encoding", "gzip, deflate, br")
	r.Header.Set("Cache-Control", "no-cache")
	r.Header.Set("Connection", "keep-alive")
	r.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	cookie := "HELLO_USER=1;Hm_lvt_637e96da78d1c6c8f8a218c811dea5fb=1572836336; Hm_lpvt_637e96da78d1c6c8f8a218c811dea5fb=1572836336;"
	for key,value := range preDatas.Cookies {
		cookie = cookie + key+"="+strings.Split(value,"=")[0]+";"
	}
	cookie = strings.Replace(cookie,"domain;","",-1)
	r.Header.Set("Cookie",cookie)
	client := &http.Transport{}
	response, err := client.RoundTrip(r)
	fmt.Printf("response status[%s] \n",response.Status)
}
