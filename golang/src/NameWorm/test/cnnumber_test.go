package test

import (
	"NameWorm/utils/aesutil"
	"encoding/json"
	"fmt"
	"NameWorm/cnnumber"
	"testing"
)

func TestGetNumberIndexPage(t *testing.T){
	preDatas := cnnumber.GetNameIndexPage()
	numNameData := cnnumber.BuildNumNameData("蒲", "思", "2", "2019", "11", "4")
	numberCal := cnnumber.GetNameNumber(preDatas, numNameData)
	if data, e := json.Marshal(numberCal);e != nil{
		fmt.Println(e)
	}else{
		fmt.Printf("%s \n",data)
	}
	//cnnumber.HttpPost(preDatas,numNameData)
	}

func TestLisenceCode(t *testing.T){
	value := make(map[string]int64)
	fmt.Println(value["123"])
}


func TestAes(t *testing.T){
	text := `{"firstName":"白","suffixName":"","gender":"1","year":"2020","month":"11","day":"1","licenseCode":"2847"}`
	fmt.Println(aesutil.Encrypt(text))
}
