package userLog

import "time"

const (
	UserLogCreate = 1
)

// UserLog 用户日志
type UserLog struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement;"`
	UserID    int       `json:"user_id" gorm:"type:int(11);not null;comment:用户🆔"`
	Type      uint8     `json:"type" gorm:"type:tinyint(1);not null;comment:日志类型"`
	Remark    string    `json:"remark" gorm:"type:varchar(100);comment:描述"`
	CreatedAt time.Time `json:"created_at"`
}

func New(attr ...Attr) *UserLog {
	log := &UserLog{}
	Attrs(attr).apply(log)

	return log
}
