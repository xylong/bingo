package GormDao

import (
	"github.com/xylong/bingo/test/internal/domain/repository"
	"gorm.io/gorm"
)

type UserLogDao struct {
	db *gorm.DB
}

func (d *UserLogDao) Create(modeler repository.Modeler) error {
	return d.db.Create(modeler).Error
}
