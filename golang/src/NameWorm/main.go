package main

import (
	_ "NameWorm/restful/client"
	"NameWorm/restful/server"
)

func main() {
	server.StartRestfulApi{}.StartApi()
}