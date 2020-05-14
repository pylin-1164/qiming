package main

import (
	"fmt"
	"NameWorm/cnname"
)

func main() {
	firstName := cnname.FirstName{}
	firstName.FindUrls()

	cnname.DeleteAll()//删除历史数据
	for _,data := range firstName.Datas {
		fmt.Printf("[%s] -> [%s] \n",data.First,data.Path)
		suffixNames := cnname.DrillDownNames(data)
		fmt.Printf("save first name [%s] and suffix name count[%d] \n",data.First, len(suffixNames))
		cnname.SaveName2Db(data.First,suffixNames)
	}
	fmt.Println("done ...")
}
