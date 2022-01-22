package setting

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"strings"
)

var (
	// DbString connection string
	DbString string

	// LogLevel app config
	IdPrefix string
	LogLevel string
	AppName  string
	Port     string
	// Issuer jwt config
	Issuer       string
	JwtKey       string
	ExpiresAt    string
	Extension    []string
	ResourcePath string

	//expr
	ExprColumns []string
	//local
	LocalColumns []string
	//skill
	SkillColumns []string
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
	loadColumns(cfg)
}
func loadColumns(cfg *ini.File) {
	ExprColumns = strings.Split(cfg.Section("expr").Key("columns").String(), " ")
	LocalColumns = strings.Split(cfg.Section("local").Key("columns").String(), " ")
	SkillColumns = strings.Split(cfg.Section("skill").Key("columns").String(), " ")
}
func loadDB(cfg *ini.File) {
	DbString = cfg.Section("mysql").Key("dsn").String()
}
func loadApp(cfg *ini.File) {
	LogLevel = cfg.Section("app").Key("log_level").String()
	AppName = cfg.Section("app").Key("app_name").String()
	Port = cfg.Section("app").Key("port").String()
	IdPrefix = cfg.Section("app").Key("id_prefix").String()
	ResourcePath = cfg.Section("app").Key("resource_path").String()
	s := cfg.Section("app").Key("exts").String()
	Extension = strings.Split(s, " ")
	if _, err := os.Stat(ResourcePath); err != nil {
		os.MkdirAll(ResourcePath, 460)
	}
}
func loadJwt(cfg *ini.File) {
	Issuer = cfg.Section("jwt").Key("issuer").String()
	JwtKey = cfg.Section("jwt").Key("jwt_key").String()
	ExpiresAt = cfg.Section("jwt").Key("expires_at").String()
}
