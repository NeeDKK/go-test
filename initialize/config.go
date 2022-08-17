package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var Gorm = new(_gorm)

type _gorm struct{}

type writer struct {
	logger.Writer
}

func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

func (w *writer) Printf(message string, data ...interface{}) {
	w.Writer.Printf(message, data...)
}

// Config gorm 自定义配置
func (g *_gorm) Config() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	config.Logger = _default.LogMode(logger.Info)
	return config
}

func GormMysql() *gorm.DB {
	dbConfigString := "root:root@tcp(10.25.10.125:3306)/audit?charset=utf8"
	mysqlConfig := mysql.Config{
		DSN:                       dbConfigString, // DSN data source name
		DefaultStringSize:         191,            // string 类型字段的默认长度
		SkipInitializeWithVersion: false,          // 根据版本自动配置
	}
RETRY:
	if db, err := gorm.Open(mysql.New(mysqlConfig), Gorm.Config()); err != nil {
		//mysql重连
		fmt.Println("==========================mysql重连")
		goto RETRY
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(0)
		sqlDB.SetMaxOpenConns(0)
		return db
	}
}
