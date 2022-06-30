package user

import (
	"gorm.io/gorm"
	"time"
)

type Attr func(user *User)
type Attrs []Attr

func (a Attrs) apply(user *User) {
	for _, attr := range a {
		attr(user)
	}
}

// User 用户
type User struct {
	ID        int            `gorm:"primaryKey;autoIncrement;" xorm:"'id' int(11) pk autoincr notnull" json:"id"`
	Phone     string         `gorm:"type:char(11);uniqueIndex;comment:手机号" json:"phone"`
	Email     string         `gorm:"type:varchar(50);default:null;uniqueIndex;comment:邮件" json:"email"`
	Avatar    string         `gorm:"type:varchar(100);comment:头像" json:"avatar"`
	Nickname  string         `gorm:"type:varchar(20);not null;comment:昵称" json:"nickname"`
	Password  string         `gorm:"type:varchar(32);comment:密码" json:"password"`
	Birthday  time.Time      `gorm:"type:date;default:null;comment:出生日期" json:"birthday"`
	Gender    int            `gorm:"type:tinyint(1);default:-1;comment:-1保密 0女 1男" json:"gender"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;->:false;"`

	Info *UserInfo // has one
}

func New(attr ...Attr) *User {
	u := &User{
		Info: NewUserInfo(),
	}

	Attrs(attr).apply(u)
	return u
}

func WithPhone(phone string) Attr {
	return func(user *User) {
		user.Phone = phone
	}
}

func WithName(name string) Attr {
	return func(user *User) {
		user.Nickname = name
	}
}

func WithPassword(password string) Attr {
	return func(user *User) {
		user.Password = password
	}
}
