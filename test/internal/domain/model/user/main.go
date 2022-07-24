package user

import (
	"github.com/xylong/bingo/test/internal/domain"
	"github.com/xylong/bingo/test/internal/domain/model"
	"github.com/xylong/bingo/test/internal/domain/repository"
	"gorm.io/gorm"
	"strings"
	"time"
)

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

	Dao repository.IUser `gorm:"-"`
}

func New(attr ...domain.Attr) *User {
	user := &User{
		Wechat: NewWechat(),
		Info:   NewInfo(),
	}

	domain.Attrs(attr).Apply(user)

	return user
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Create() error {
	return u.Dao.Create(u)
}

func (u *User) Get() ([]*User, int64, error) {
	var users []*User
	total, err := u.Dao.Get(&users)
	return users, total, err
}

// HidePhone 隐藏手机号
func (u *User) HidePhone() {
	if u.Phone != "" {
		u.Phone = strings.Join([]string{u.Phone[0:3], "****", u.Phone[7:]}, "")
	}
}
