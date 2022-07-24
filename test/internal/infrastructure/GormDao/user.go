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

func (u *UserDao) Get(users interface{}, comparator ...func(db *gorm.DB) *gorm.DB) error {
	return u.db.Scopes(comparator...).Find(users).Error
}

func (u *UserDao) Count(modeler repository.Modeler, total *int64, comparator ...func(db *gorm.DB) *gorm.DB) error {
	return u.db.Model(modeler).Scopes(comparator...).Count(total).Error
}
