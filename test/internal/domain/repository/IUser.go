package repository

import "gorm.io/gorm"

// IUser 用户接口
type IUser interface {
	Single(Modeler) error
	Get(interface{}, ...func(db *gorm.DB) *gorm.DB) error
	Count(Modeler, *int64, ...func(db *gorm.DB) *gorm.DB) error
	Create(Modeler) error
}
