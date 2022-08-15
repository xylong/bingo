package config

import (
	"fmt"
	"github.com/xylong/bingo"
	"github.com/xylong/bingo/ioc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// Adapter db配置
type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

// Gorm 创建gorm
func (a *Adapter) Gorm() *gorm.DB {
	conf := ioc.Factory.Get((*bingo.Config)(nil))
	if conf == nil {
		return nil
	}

	mysqlConf := conf.(*bingo.Config).Mysql

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.DB, mysqlConf.Charset)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName:                "mysql",
		ServerVersion:             "8.0.13",
		DSN:                       dsn,
		Conn:                      nil,
		SkipInitializeWithVersion: true,
		DefaultStringSize:         191,
	}), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logger.Info,
		}),
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Second * 10)

	return db
}
