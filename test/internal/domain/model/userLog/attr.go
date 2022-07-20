package userLog

import "github.com/xylong/bingo/test/internal/domain"

func WithUserID(uid int) domain.Attr {
	return func(m interface{}) {
		if uid > 0 {
			m.(*UserLog).UserID = uid
		}
	}
}

func WithType(t uint8) domain.Attr {
	return func(m interface{}) {
		m.(*UserLog).Type = t
	}
}

func WithRemark(remark string) domain.Attr {
	return func(m interface{}) {
		if remark != "" {
			m.(*UserLog).Remark = remark
		}
	}
}
