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

// Single 获取单个用户
func (u *UserDao) Single(modeler repository.Modeler) error {
	if err := u.db.First(modeler).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return nil
}

// Get 获取用户列表
func (u *UserDao) Get(users interface{}, comparator ...func(db *gorm.DB) *gorm.DB) error {
	if err := u.db.Scopes(comparator...).Find(users).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return nil
}

// Count 统计
func (u *UserDao) Count(modeler repository.Modeler, total *int64, comparator ...func(db *gorm.DB) *gorm.DB) error {
	if err := u.db.Model(modeler).Scopes(comparator...).Count(total).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return nil
}
