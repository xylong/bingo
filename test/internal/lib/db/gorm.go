package db

import (
	"github.com/xylong/bingo/test/internal/domain/model/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitGorm() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName:                "mysql",
		ServerVersion:             "8.0.13",
		DSN:                       dsn,
		Conn:                      nil,
		SkipInitializeWithVersion: false,
		DefaultStringSize:         191,
	}), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&user.User{}, &user.UserInfo{})

	//u := &user.User{
	//	Phone:    "13888888888",
	//	Nickname: "琳琳",
	//	Password: "123456",
	//	Info: &user.UserInfo{
	//		WechatUnionid:        "aaa",
	//		WechatAppletOpenid:   "bbb",
	//		WechatOfficialOpenid: "ccc",
	//	},
	//}

	u := user.NewUser(user.WithName("露露"), user.WithPhone("13999999999"), user.WithUnionid("xxoo"))

	db.Create(u)

	return db
}
