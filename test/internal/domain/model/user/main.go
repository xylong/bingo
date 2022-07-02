package user

import (
	"github.com/xylong/bingo/test/internal/domain/model/profile"
	"github.com/xylong/bingo/test/internal/domain/model/repository"
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

func New(attr ...Attr) *User {
	u := &User{
		Wechat:  NewWechat(),
		Info:    NewInfo(),
		Profile: profile.New(),
	}

	Attrs(attr).apply(u)
	return u
}

func WithRepo(repo repository.IUserRepo) Attr {
	return func(user *User) {
		user.repo = repo
	}
}

// User 用户
type User struct {
	ID      int `gorm:"primaryKey;autoIncrement;" xorm:"'id' int(11) pk autoincr notnull" json:"id"`
	*Wechat     // 微信信息
	*Info       // 用户信息

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;->:false;"`

	Profile *profile.Profile // has one

	repo repository.IUserRepo `gorm:"-"`
}

func (u *User) Get() error {
	return u.repo.GetByID(u)
}
