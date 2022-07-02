package GromDao

import (
	"github.com/xylong/bingo/test/internal/domain/model/profile"
	"github.com/xylong/bingo/test/internal/domain/model/repository"
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
func (r *ProfileDao) GetByUser(model repository.IModel) error {
	return r.db.Where("user_id=?", model.(*profile.Profile).UserID).First(model).Error
}
