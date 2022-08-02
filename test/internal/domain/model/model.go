package model

import (
	"github.com/xylong/bingo/test/internal/infrastructure/dao/GormDao"
	"github.com/xylong/bingo/test/internal/lib/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
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
func (*Model) Compare(field string, value interface{}, comparator int) GormDao.Scope {
	return GormDao.Compare(field, value, comparator)
}

// Filter 过滤
func (m *Model) Filter(comparator ...func(db *gorm.DB) *gorm.DB) {
	db.DB.Scopes(comparator...).Find(m.child)
}

// Order 字段排序
func (m *Model) Order(field string) func(bool) GormDao.Scope {
	return func(b bool) GormDao.Scope {
		builder := strings.Builder{}
		builder.WriteString(field)

		if b {
			builder.WriteString(" asc")
		} else {
			builder.WriteString(" desc")
		}

		return func(db *gorm.DB) *gorm.DB {
			return db.Order(builder.String())
		}
	}
}

// SimplePage 简单分页
func (m *Model) SimplePage(page, pageSize int) GormDao.Scope {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
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
