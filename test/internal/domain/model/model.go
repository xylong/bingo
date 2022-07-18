package model

import (
	"github.com/xylong/bingo/test/internal/lib/db"
	"gorm.io/gorm"
)

// Model 基础模型
type Model struct {
	child interface{}
}

func NewModel(data interface{}) *Model {
	return &Model{child: data}
}

// Compare 字段过滤
func (*Model) Compare(field string, value interface{}, comparator int) Scope {
	return Compare(field, value, comparator)
}

// Filter 过滤
func (m *Model) Filter(comparator ...func(db *gorm.DB) *gorm.DB) {
	db.DB.Scopes(comparator...).Find(m.child)
}

// One 根据过滤条件查单个记录
func (m *Model) One(comparator ...func(db *gorm.DB) *gorm.DB) {
	db.DB.Scopes(comparator...).First(m.child)
}

// GetByID 根据🆔获取
func (m *Model) GetByID(id int) interface{} {
	db.DB.First(m.child, id)
	return m.child
}

// Save 指定字段保存
func (m *Model) Save(column ...string) error {
	return db.DB.Select(column).Save(m.child).Error
}
