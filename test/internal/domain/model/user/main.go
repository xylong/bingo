package user

import (
	"gorm.io/gorm"
	"time"
)

func New(attr ...Attr) *User {
	u := &User{
		Wechat: NewWechat(),
		Info:   NewInfo(),
		//Profile: profile.New(),
	}

	Attrs(attr).apply(u)
	return u
}

// User 用户
type User struct {
	ID      int `gorm:"primaryKey;autoIncrement;" xorm:"'id' int(11) pk autoincr notnull" json:"id"`
	*Wechat     // 微信信息
	*Info       // 用户信息

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;->:false;"`

	//Profile *profile.Profile // has one

	//Repo repository.IUser `gorm:"-"`
}

func (u *User) Get() error {
	return nil
}

func (u *User) Create() error {
	return nil
}
