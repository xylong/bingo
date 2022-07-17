package model

import (
	"github.com/xylong/bingo/test/internal/lib/db"
	"gorm.io/gorm"
)

const (
	Equal        = iota // =
	NotEqual            // <>
	GreaterThan         // >
	GreaterEqual        // >=
	LessThan            // <
	LessEqual           // <=
	In                  // in
	NotIn               // not in
	Like                // like
	NotLike             // not like
)

// Compare 查询条件比较
type Compare func(*gorm.DB) *gorm.DB

type Model struct {
}

// Filter 过滤
func (m *Model) Filter(comparator ...func(db *gorm.DB) *gorm.DB) {
	db.DB.Scopes(comparator...)
}
