package userLog

type (
	Attrs []Attr
	Attr  func(*UserLog)
)

func (a Attrs) apply(log *UserLog) {
	for _, attr := range a {
		attr(log)
	}
}

func WithUserID(uid int) Attr {
	return func(log *UserLog) {
		if uid > 0 {
			log.UserID = uid
		}
	}
}

func WithType(t uint8) Attr {
	return func(log *UserLog) {
		log.Type = t
	}
}

func WithRemark(remark string) Attr {
	return func(log *UserLog) {
		if remark != "" {
			log.Remark = remark
		}
	}
}
