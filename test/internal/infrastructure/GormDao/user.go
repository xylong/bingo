package GormDao

import (
	. "github.com/xylong/bingo/test/internal/domain/model/user"
	"gorm.io/gorm"
)

// UserDao 用户
type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

// Get 获取用户
func (dao *UserDao) Get(user *User) error {
	return dao.db.First(user).Error
}

// GetByPhone 根据手机号获取用户
func (dao *UserDao) GetByPhone(phone string) *User {
	var u *User
	dao.db.Where("phone=?", phone).First(u)
	return u
}

// Create 创建用户
func (dao *UserDao) Create(user *User) error {
	return dao.db.Create(user).Error
}
