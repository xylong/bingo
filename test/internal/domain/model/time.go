package model

import (
	"encoding/json"
	"fmt"
)

// Time 时间
type Time struct {
	Second int64 `json:"second" gorm:"type:bigint(20);default:0;comment:秒"`
}

func NewTime() *Time {
	return &Time{}
}

func (t *Time) MarshalJSON() ([]byte, error) {
	hour := t.Second / 3600
	minute := (t.Second - hour*3600) / 60
	second := t.Second - hour*3600 - minute*60

	return json.Marshal(map[string]interface{}{
		"Second": t.Second,
		"Format": fmt.Sprintf("%d时%d分%d秒", hour, minute, second),
	})
}

// Add 增加时长
func (t *Time) Add(sec int64) {
	t.Second += sec
}
