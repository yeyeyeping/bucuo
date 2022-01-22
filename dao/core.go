package dao

import (
	"bucuo/model/table"
	"bucuo/util/setting"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB  *gorm.DB
	err error
	cfg *ini.File
)

func initLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // 日志级别
		},
	)
}
func initDB() {
	DB, err = gorm.Open(mysql.Open(setting.DbString), &gorm.Config{
		Logger: initLogger(),
	})
	if err != nil {
		log.Fatalf("Fail to open database: %v", err)
		os.Exit(1)
	}
	//检查有没有表，没有
	if !DB.Migrator().HasTable(&table.User{}) {
		err := DB.Set("gorm:table_options", "ENGINE=InnoDB").
			Set("gorm:table_options", "default charset=UTF8").
			AutoMigrate(
				&table.Comment{},
				&table.ExprPost{},
				&table.Label{},
				&table.LocalPost{},
				&table.LostRegisteration{},
				&table.Reply{},
				&table.Resource{},
				&table.SkillPost{},
				&table.User{},
			)
		log.Printf("\033[0;31;47m%#v\033[0m\n", err)
	}
}

func initDBPool() {
	sqlDB, _ := DB.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
func init() {
	initDB()
	initDBPool()
}
func Ok() {
	log.Printf("ok")
}
