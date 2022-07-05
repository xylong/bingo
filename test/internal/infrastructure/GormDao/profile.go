package GormDao

import (
	"github.com/xylong/bingo/test/internal/domain/model/profile"
	"github.com/xylong/bingo/test/internal/domain/repository"
	"gorm.io/gorm"
)

// ProfileDao 用户信息
type ProfileDao struct {
	db *gorm.DB
}

func NewProfileDao(db *gorm.DB) *ProfileDao {
	return &ProfileDao{db: db}
}

// GetByUser 根据用户🆔获取
func (dao *ProfileDao) GetByUser(model repository.IModel) error {
	return dao.db.Where("user_id=?", model.(*profile.Profile).UserID).First(model).Error
}

// Create 创建用户信息
func (dao *ProfileDao) Create(model repository.IModel) error {
	return dao.db.Create(model).Error
}
