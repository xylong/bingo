package aggregation

import (
	"github.com/xylong/bingo/test/internal/domain"
	"github.com/xylong/bingo/test/internal/domain/model/profile"
	"github.com/xylong/bingo/test/internal/domain/model/user"
	"github.com/xylong/bingo/test/internal/domain/model/userLog"
	"github.com/xylong/bingo/test/internal/domain/repository"
	. "github.com/xylong/bingo/test/internal/infrastructure/error"
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

	if member.User != nil && member.Profile != nil {
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
	m.Log.Dao = m.LogRepo
	if err := m.Log.Create(); err != nil {
		return NewOperatorError(CreateUserLogError)
	}

	return nil
}

func (m *Member) GetLog() []*userLog.UserLog {
	return nil
}
