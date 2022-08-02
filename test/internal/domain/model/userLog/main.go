package userLog

import (
	"github.com/xylong/bingo/test/internal/application/dto"
	"github.com/xylong/bingo/test/internal/domain"
	"github.com/xylong/bingo/test/internal/domain/model"
	"github.com/xylong/bingo/test/internal/domain/repository"
	"github.com/xylong/bingo/test/internal/infrastructure/dao/GormDao"
	"gorm.io/gorm"
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

	Dao          repository.UserLogger `gorm:"-" json:"-"`
	*model.Model `gorm:"-" json:"-"`
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

// Get 获取用户日志
func (l *UserLog) Get(req *dto.UserLogReq) (logs []*UserLog, err error) {
	s := []func(*gorm.DB) *gorm.DB{
		l.Order("id")(false),
	}

	if req.Page > 0 && req.PageSize > 0 {
		s = append(s, l.SimplePage(req.Page, req.PageSize))
	}

	if req.ID > 0 {
		s = append(s, l.CompareUser(req.ID))
	}

	err = l.Dao.Get(&logs, s...)
	return logs, err
}

// CompareUser 比较用户🆔
func (l *UserLog) CompareUser(id int) GormDao.Scope {
	return l.Compare("user_id", id, GormDao.Equal)
}
