package model

import (
	"fmt"
	"gorm.io/gorm"
)

// mysql查询方法
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

// Scope 查询条件处理方法
type Scope func(*gorm.DB) *gorm.DB

// Compare 字段比较
func Compare(field string, value interface{}, comparator int) Scope {
	return func(db *gorm.DB) *gorm.DB {
		switch comparator {
		case Equal:
			return db.Where(fmt.Sprintf("%s = ?", field), value)
		case NotEqual:
			return db.Where(fmt.Sprintf("%s <> ?", field), value)
		case GreaterThan:
			return db.Where(fmt.Sprintf("%s > ?", field), value)
		case GreaterEqual:
			return db.Where(fmt.Sprintf("%s >= ?", field), value)
		case LessThan:
			return db.Where(fmt.Sprintf("%s < ?", field), value)
		case LessEqual:
			return db.Where(fmt.Sprintf("%s <= ?", field), value)
		case In:
			return db.Where(fmt.Sprintf("%s in ?", field), value)
		case NotIn:
			return db.Where(fmt.Sprintf("%s not in ?", field), value)
		case Like:
			return db.Where(fmt.Sprintf("%s like ?", field), value)
		case NotLike:
			return db.Where(fmt.Sprintf("%s not like?", field), value)
		default:
			return db
		}
	}
}
