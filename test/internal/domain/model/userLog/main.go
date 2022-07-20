package userLog

import (
	"github.com/xylong/bingo/test/internal/domain"
	"github.com/xylong/bingo/test/internal/domain/repository"
	"time"
)

const (
	Register = iota // 注册
	Login           // 登录
	Logout          // 登出
)

// UserLog 用户日志
type UserLog struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement;"`
	UserID    int       `json:"user_id" gorm:"type:int(11);not null;comment:用户🆔"`
	Type      uint8     `json:"type" gorm:"type:tinyint(1);not null;comment:日志类型"`
	Remark    string    `json:"remark" gorm:"type:varchar(100);comment:描述"`
	CreatedAt time.Time `json:"created_at"`

	Dao repository.UserLogger `gorm:"-"`
}

func New(attr ...domain.Attr) *UserLog {
	log := &UserLog{}
	domain.Attrs(attr).Apply(log)

	return log
}

func (l *UserLog) TableName() string {
	return "user_logs"
}

func (l *UserLog) Create() error {
	return l.Dao.Create(l)
}
