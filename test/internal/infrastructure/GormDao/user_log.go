package GormDao

import (
	"github.com/xylong/bingo/test/internal/domain/repository"
	"gorm.io/gorm"
)

type UserLogDao struct {
	db *gorm.DB
}

func NewUserLogDao(db *gorm.DB) *UserLogDao {
	return &UserLogDao{db: db}
}

func (d *UserLogDao) Create(modeler repository.Modeler) error {
	return d.db.Create(modeler).Error
}

// Get 获取用户日志
func (d *UserLogDao) Get(logs interface{}, comparator ...func(db *gorm.DB) *gorm.DB) error {
	if err := d.db.Scopes(comparator...).Find(logs).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return nil
}
