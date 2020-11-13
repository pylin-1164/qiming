package config

import (
	"flag"
	"github.com/larspensjo/config"
	"log"
)


//topic list
var WECHART = make(map[string]string)
var UI_SERVER = make(map[string]string)
var API_SERVER_PORT = 8099
var DataBaseInfo = struct {
	Host 		string
	Port		string
	Name 		string
	User 		string
	Password 	string
}{}

func init() {
	configFile := flag.String("configfile", "config.ini", "General configuration file")
	//set config file std
	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		log.Fatalf("Fail to find", *configFile, err)
	}
	//set config file std End

	//Initialized topic from the configuration
	if cfg.HasSection("wechart") {
		appId, _ := cfg.String("wechart", "appId")
		appsecret, _ := cfg.String("wechart", "appsecret")
		WECHART["appId"] = appId
		WECHART["appsecret"] = appsecret
	}else{
		panic("[ERROR]read config.ini wechart error")
	}

	if value, err := cfg.String("ui", "address");err != nil{
		panic("[ERROR]read config.ini ui error")
	}else {
		UI_SERVER["host"] = value
	}

	if !cfg.HasSection("database"){
		panic("[ERROR]read config.ini database error")
	}

	DataBaseInfo.Host, _ = cfg.String("database", "host")
	DataBaseInfo.Port, _ = cfg.String("database", "port")
	DataBaseInfo.Name, _ = cfg.String("database", "name")
	DataBaseInfo.User, _ = cfg.String("database", "user")
	DataBaseInfo.Password, _ = cfg.String("database", "password")

	//API 监听端口
	if cfg.HasSection("api"){
		API_SERVER_PORT,_ = cfg.Int("api","port")
	}

}

func GetWechartAppID()string{
	return WECHART["appId"]
}

func GetWechartSecret()string{
	return WECHART["appsecret"]
}