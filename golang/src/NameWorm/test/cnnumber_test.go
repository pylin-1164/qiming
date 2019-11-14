package test

import (
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

