package repository

import (
	"gorm.io/gorm"
)

type UserLogger interface {
	Create(Modeler) error
	Get(interface{}, ...func(db *gorm.DB) *gorm.DB) error
	ContinuousLogin() []int
}
