package main

import (
	"os"

	"server/conf"
	"server/router"
)

func main() {
	dir, _ := os.Getwd()
	var bashDir = dir + "/../"

	var saveDir = bashDir + "upload/"
	// var logDir = bashDir + "logs/"
	var confDir = bashDir + "conf/"

	os.Setenv("SAVE_DIR", saveDir)
	// os.Setenv("LOG_DIR", logDir)
	os.Setenv("CONF_DIR", confDir)

	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := router.NewRouter()
	r.Run()
}
