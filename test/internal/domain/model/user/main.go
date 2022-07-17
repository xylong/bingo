package user

import (
	"github.com/xylong/bingo/test/internal/domain/model"
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
	*model.Model `gorm:"-"`

	ID      int               `gorm:"primaryKey;autoIncrement;" xorm:"'id' int(11) pk autoincr notnull" json:"id"`
	*Wechat `gorm:"embedded"` // 微信信息
	*Info   `gorm:"embedded"` // 用户信息

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;->:false;"`

	//Profile *profile.Profile // has one

	//Dao repository.IUser `gorm:"-"`
}

func (u *User) Get() error {
	return nil
}

func (u *User) Create() error {
	return nil
}
