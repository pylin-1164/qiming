package test

import (
	"NameWorm/restful/server"
	"encoding/xml"
	"fmt"
	"testing"
)

func TestMarshulXML(t *testing.T){
	menuXML := server.RequestClickMenuXML{}
	menuXML.FromUserName = server.FromUserName{Text:"abcd"}
	menuXML.ToUserName = server.ToUserName{Text:"123123"}
	marshal, _ := xml.Marshal(menuXML)
	fmt.Printf("%s",marshal)
}