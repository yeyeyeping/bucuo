package dao

import (
	"bucuo/model"
	"bucuo/util"
	"log"
	"os"
	"time"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)
}
func initDB() {
	DB, err = gorm.Open(mysql.Open(util.DbString), &gorm.Config{
		Logger: initLogger(),
	})
	if err != nil {
		log.Fatalf("Fail to open database: %v", err)
		os.Exit(1)
	}
	//检查有没有表，没有
	if !DB.Migrator().HasTable(&model.User{}) {
		DB.AutoMigrate(&model.User{})
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
