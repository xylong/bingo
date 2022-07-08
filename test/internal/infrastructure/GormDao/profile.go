package GormDao

import (
	"github.com/xylong/bingo/test/internal/domain/model/profile"
	"github.com/xylong/bingo/test/internal/domain/model/user"
	"gorm.io/gorm"
)

// ProfileDao 用户信息
type ProfileDao struct {
	db *gorm.DB
}

func NewProfileDao(db *gorm.DB) *ProfileDao {
	return &ProfileDao{db: db}
}

// GetByUser 根据用户
func (dao *ProfileDao) GetByUser(user *user.User) (p *profile.Profile, err error) {
	err = dao.db.Where("user_id=?", user.ID).First(p).Error
	return
}

// Create 创建用户信息
func (dao *ProfileDao) Create(profile *profile.Profile) error {
	return dao.db.Create(profile).Error
}
