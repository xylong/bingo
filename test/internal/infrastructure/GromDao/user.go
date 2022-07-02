package GromDao

import (
	"github.com/xylong/bingo/test/internal/domain/model/repository"
	"github.com/xylong/bingo/test/internal/domain/model/user"
	"gorm.io/gorm"
)

// UserDao 用户
type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

// GetByID 根据主键🆔获取
func (r *UserDao) GetByID(model repository.IModel) error {
	return r.db.First(model, model.(*user.User).ID).Error
}
