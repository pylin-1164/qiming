package main

import (
	_ "NameWorm/common"
	_ "NameWorm/restful/client"
	"NameWorm/restful/server"
)

func main() {
	server.StartRestfulApi{}.StartApi()
}