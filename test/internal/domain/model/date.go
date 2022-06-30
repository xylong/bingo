package model

import (
	"encoding/json"
	"time"
)

// At 时刻
type At struct {
	Date time.Time
}

func NewAt() *At {
	return &At{Date: time.Now()}
}

func (a At) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Date    time.Time `json:"date"`
		Default string    `json:"default"`
		Unix    int64     `json:"unix"`
	}{
		Date:    a.Date,
		Default: a.Format(),
		Unix:    a.Date.Unix(),
	})
}

func (a *At) Format() string {
	return a.Date.Format("2006-01-02 15:04:05")
}
