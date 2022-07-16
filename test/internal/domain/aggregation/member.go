package aggregation

import (
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

func NewMember(user *user.User, profile *profile.Profile, userRepo repository.IUser, profileRepo repository.IProfile, logRepo repository.IUserLog) *Member {
	return &Member{User: user, Profile: profile, UserRepo: userRepo, ProfileRepo: profileRepo, LogRepo: logRepo}
}

func NewMemberByPhone(phone string, userRepo repository.IUser, profileRepo repository.IProfile) *Member {
	u := userRepo.GetByPhone(phone)

	return &Member{
		User:        u,
		UserRepo:    userRepo,
		ProfileRepo: profileRepo,
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
