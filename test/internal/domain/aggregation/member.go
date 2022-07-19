package aggregation

import (
	"github.com/xylong/bingo/test/internal/domain"
	"github.com/xylong/bingo/test/internal/domain/model/profile"
	"github.com/xylong/bingo/test/internal/domain/model/user"
	"github.com/xylong/bingo/test/internal/domain/model/userLog"
	"github.com/xylong/bingo/test/internal/domain/repository"
)

// Member 会员
// 聚合
type Member struct {
	User    *user.User // 根
	Profile *profile.Profile
	Log     *userLog.UserLog

	UserRepo    repository.IUser
	ProfileRepo repository.IProfile
	LogRepo     repository.IUserLog
}

func NewMember(attr ...domain.Attr) *Member {
	member := &Member{}
	domain.Attrs(attr).Apply(member)

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

func WIthProfileRepo(iProfile repository.IProfile) domain.Attr {
	return func(i interface{}) {
		i.(*Member).ProfileRepo = iProfile
	}
}

// Create 创建用户
func (m *Member) Create() error {
	err := m.UserRepo.Create(m.User)
	if err != nil {
		return err
	}

	m.Log = userLog.New(userLog.WithUserID(m.User.ID), userLog.WithType(userLog.Register), userLog.WithRemark("新增用户"))
	return m.LogRepo.Save(m.Log)
}

func (m *Member) GetLog() []*userLog.UserLog {
	return nil
}
