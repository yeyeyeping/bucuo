package util

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
)

var (
	//connection string
	DbString string

	//app config
	LogLevel string
	AppName  string

	// jwt config
	Issuer    string
	JwtKey    string
	ExpiresAt string
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
	loadJwt(cfg)
}
func loadDB(cfg *ini.File) {
	DbString = cfg.Section("mysql").Key("dsn").String()
}
func loadApp(cfg *ini.File) {
	LogLevel = cfg.Section("app").Key("log_level").String()
	AppName = cfg.Section("app").Key("app_name").String()
}
func loadJwt(cfg *ini.File) {
	Issuer = cfg.Section("jwt").Key("issuer").String()
	JwtKey = cfg.Section("jwt").Key("jwt_key").String()
	ExpiresAt = cfg.Section("jwt").Key("expires_at").String()
}
