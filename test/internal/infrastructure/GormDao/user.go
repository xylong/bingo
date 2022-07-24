package GormDao

import (
	"github.com/xylong/bingo/test/internal/domain/repository"
	"gorm.io/gorm"
)

// UserDao 用户
type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

// Create 创建用户
func (u *UserDao) Create(modeler repository.Modeler) error {
	return u.db.Create(modeler).Error
}

func (u *UserDao) Get(users interface{}) (total int64, err error) {
	err = u.db.Find(users).Count(&total).Error
	return
}
