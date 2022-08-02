package GormDao

import (
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

// Create 创建用户信息
func (dao *ProfileDao) Create(modeler repository.Modeler) error {
	return dao.db.Create(modeler).Error
}
