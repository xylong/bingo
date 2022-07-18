package model

import (
	"github.com/xylong/bingo/test/internal/lib/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Model 基础模型
type Model struct {
	child    interface{}
	conflict func() clause.OnConflict
}

func NewModel(data interface{}) *Model {
	return &Model{
		child: data,
		conflict: func() clause.OnConflict {
			return clause.OnConflict{DoNothing: true}
		},
	}
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
	return db.DB.Clauses(m.conflict()).Select(column).Save(m.child).Error
}

// Conflict 冲突
// 唯一约束冲突时的执行方案
func (m *Model) Conflict(column []string, assignment map[string]interface{}) {
	columns := make([]clause.Column, 0)
	for _, item := range column {
		columns = append(columns, clause.Column{
			Name: item,
		})
	}

	m.conflict = func() clause.OnConflict {
		return clause.OnConflict{
			Columns:   columns,
			DoUpdates: clause.Assignments(assignment),
		}
	}
}
