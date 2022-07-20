package profile

import (
	"github.com/xylong/bingo/test/internal/domain"
	"time"
)

func WithUserID(id int) domain.Attr {
	return func(v interface{}) {
		if id > 0 {
			v.(*Profile).UserID = id
		}
	}
}

func WithPassword(password string) domain.Attr {
	return func(v interface{}) {
		if len(password) > 0 {
			v.(*Profile).Password = password
		}
	}
}

func WithBirthday(birthday string) domain.Attr {
	return func(v interface{}) {
		if len(birthday) > 0 {
			if t, err := time.Parse("2006-01-02 15:04:05", birthday); err == nil {
				v.(*Profile).Birthday = t
			}
		}
	}
}

func WithGender(gender int) domain.Attr {
	return func(v interface{}) {
		v.(*Profile).Gender = int8(gender)
	}
}

func WithLevel(level int) domain.Attr {
	return func(v interface{}) {
		if level >= 0 {
			v.(*Profile).Level = int8(level)
		}
	}
}

func WithSignature(signature string) domain.Attr {
	return func(v interface{}) {
		if signature != "" {
			v.(*Profile).Signature = signature
		}
	}
}
