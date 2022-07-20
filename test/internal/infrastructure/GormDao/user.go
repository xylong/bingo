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
func (dao *UserDao) Create(modeler repository.Modeler) error {
	return dao.db.Create(modeler).Error
}
