package aggregation

import (
	"github.com/xylong/bingo/test/internal/application/dto"
	"github.com/xylong/bingo/test/internal/domain"
	"github.com/xylong/bingo/test/internal/domain/model/profile"
	"github.com/xylong/bingo/test/internal/domain/model/user"
	"github.com/xylong/bingo/test/internal/domain/model/userLog"
	"github.com/xylong/bingo/test/internal/domain/repository"
	. "github.com/xylong/bingo/test/internal/infrastructure/error"
	"go.uber.org/zap"
)

// Member 会员
// 聚合
type Member struct {
	User    *user.User // 根
	Profile *profile.Profile
	Log     *userLog.UserLog

	UserRepo    repository.IUser
	ProfileRepo repository.Profiler
	LogRepo     repository.UserLogger
}

func NewMember(attr ...domain.Attr) *Member {
	member := &Member{}
	domain.Attrs(attr).Apply(member)

	if member.User != nil && member.UserRepo != nil {
		member.User.Dao = member.UserRepo
	}

	if member.Profile != nil && member.ProfileRepo != nil {
		member.Profile.Dao = member.ProfileRepo
	}

	if member.Log != nil && member.LogRepo != nil {
		member.Log.Dao = member.LogRepo
	}

	return member
}

func WithUser(u *user.User) domain.Attr {
	return func(i interface{}) {
		if u != nil {
			i.(*Member).User = u
		}
	}
}

func WithUserRepo(iUser repository.IUser) domain.Attr {
	return func(i interface{}) {
		if iUser != nil {
			i.(*Member).UserRepo = iUser
		}
	}
}

func WithProfile(p *profile.Profile) domain.Attr {
	return func(i interface{}) {
		if p != nil {
			i.(*Member).Profile = p
		}
	}
}

func WithProfileRepo(iProfile repository.Profiler) domain.Attr {
	return func(i interface{}) {
		i.(*Member).ProfileRepo = iProfile
	}
}

func WithLog(log *userLog.UserLog) domain.Attr {
	return func(i interface{}) {
		if log != nil {
			i.(*Member).Log = log
		}
	}
}

func WithLogRepo(logger repository.UserLogger) domain.Attr {
	return func(i interface{}) {
		if logger != nil {
			i.(*Member).LogRepo = logger
		}
	}
}

// Create 创建用户
func (m *Member) Create() error {
	if err := m.User.Create(); err != nil {
		return NewOperatorError(CreateUserError)
	}

	m.Profile.UserID = m.User.ID
	if err := m.Profile.Create(); err != nil {
		return NewOperatorError(CreateProfileError)
	}

	m.Log = userLog.New(userLog.WithUserID(m.User.ID), userLog.WithType(userLog.Register), userLog.WithRemark("新增用户"))

	return nil
}

func (m *Member) GetUsers(req *dto.UserReq) ([]*user.User, error) {
	return m.User.Get(req)
}

func (m *Member) GetUser() error {
	if err := m.User.Single(); err != nil {
		zap.L().Error("get user err", zap.Error(err))
		return NewNotFoundError(NotFoundData, "用户查询失败")
	}

	return nil
}

// GetLog 获取用户日志
func (m *Member) GetLog(req *dto.UserLogReq) []*userLog.UserLog {
	logs, err := m.Log.Get(req)
	if err != nil {
		zap.L().Error("get user log error", zap.Error(err))
	}

	return logs
}

// AddLog 添加用户日志
func (m *Member) AddLog() error {
	m.Log.Dao = m.LogRepo

	if err := m.Log.Create(); err != nil {
		return NewOperatorError(CreateUserLogError)
	}

	return nil
}
