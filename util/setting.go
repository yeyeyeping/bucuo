package util

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
)

var (
	DbString string
	LogLevel string
	AppName  string
)

func init() {
	cfg, err := ini.Load("./conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to read file: %v", err)
		os.Exit(1)
	}
	//读取数据库配置文件
	loadDB(cfg)
	loadApp(cfg)
}
func loadDB(cfg *ini.File) {
	DbString = cfg.Section("mysql").Key("dsn").String()
}
func loadApp(cfg *ini.File) {
	LogLevel = cfg.Section("app").Key("log_level").String()
	AppName = cfg.Section("app").Key("app_name").String()
}
