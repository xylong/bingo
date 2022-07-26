package user

import (
	"github.com/xylong/bingo/test/internal/application/dto"
	"github.com/xylong/bingo/test/internal/domain"
	"github.com/xylong/bingo/test/internal/domain/model"
	"github.com/xylong/bingo/test/internal/domain/repository"
	"github.com/xylong/bingo/test/internal/infrastructure/GormDao"
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

	Dao repository.IUser `gorm:"-" json:"-"`
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

func (u *User) Get(req *dto.UserReq) (users []*User, err error) {
	var s []func(*gorm.DB) *gorm.DB

	if req.Nickname != "" {
		s = append(s, u.CompareNickname(req.Nickname, GormDao.Like))
	}

	if req.Phone != "" {
		s = append(s, u.ComparePhone(req.Phone))
	}

	if err = u.Dao.Get(&users, s...); err != nil {
		return nil, err
	}

	return
}

// HidePhone 隐藏手机号
func (u *User) HidePhone() {
	if u.Phone != "" {
		u.Phone = strings.Join([]string{u.Phone[0:3], "****", u.Phone[7:]}, "")
	}
}

// CompareNickname 比较昵称
func (u *User) CompareNickname(name string, comparator int) GormDao.Scope {
	if comparator == GormDao.Equal || comparator == GormDao.NotEqual {
		return u.Compare("nickname", name, comparator)
	} else if comparator == GormDao.Like || comparator == GormDao.NotLike {
		return u.Compare("nickname", "%"+name+"%", comparator)
	} else {
		return nil
	}
}

// ComparePhone 比较手机号
func (u *User) ComparePhone(phone string) GormDao.Scope {
	return u.Compare("phone", phone, GormDao.Equal)
}
