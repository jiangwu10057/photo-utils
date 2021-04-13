package model

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database(dsn string) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})

	// Error
	if err != nil {
		panic(err)
	}
	if gin.Mode() != "release" {
		db.Debug()
	}

	//设置连接池
	sqlDB, err := db.DB()

	//超时
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	//空闲
	sqlDB.SetMaxIdleConns(20)
	//打开
	sqlDB.SetMaxOpenConns(100)

	DB = db

	migration()
}

type Tabler interface {
	TableName() string
}
