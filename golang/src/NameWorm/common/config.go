package config

import (
	"flag"
	"github.com/larspensjo/config"
	"log"
)


//topic list
var WECHART = make(map[string]string)

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
}

func GetWechartAppID()string{
	return WECHART["appId"]
}

func GetWechartSecret()string{
	return WECHART["appsecret"]
}