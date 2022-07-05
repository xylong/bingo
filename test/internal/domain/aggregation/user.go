package aggregation

import (
	"github.com/xylong/bingo/test/internal/domain/model/profile"
	"github.com/xylong/bingo/test/internal/domain/model/repository"
	"github.com/xylong/bingo/test/internal/domain/model/user"
	error2 "github.com/xylong/bingo/test/internal/infrastructure/error"
)

// FrontUserAgg 前台展示
type FrontUserAgg struct {
	User    *user.User       // 用户基础信息(聚合根)
	Profile *profile.Profile // 用户资料信息

	UserRepo    repository.IUserRepo    // 仓储
	ProfileRepo repository.IProfileRepo // 仓储
}

func NewFrontUserAgg(user *user.User, profile *profile.Profile, userRepo repository.IUserRepo, profileRepo repository.IProfileRepo) *FrontUserAgg {
	if user == nil {
		panic("root error: user")
	}

	fu := &FrontUserAgg{User: user, Profile: profile, UserRepo: userRepo, ProfileRepo: profileRepo}
	fu.User.Repo, fu.Profile.Repo = userRepo, profileRepo

	return fu
}

func (u *FrontUserAgg) Get() error {
	if u.User.ID <= 0 {
		return error2.NewNoIDError("user")
	}

	if err := u.User.Get(); err != nil {
		return error2.NewNoDataError("user")
	}

	if err := u.Profile.Get(); err != nil {
		return error2.NewNoDataError("profile")
	}

	return nil
}

// CreateUser 创建用户
func (u *FrontUserAgg) CreateUser() error {
	// 创建用户基础
	if err := u.User.Create(); err != nil {
		return error2.NewCreateError("user", err.Error())
	}

	// 创建用户信息
	u.Profile.UserID = u.User.ID
	if err := u.Profile.Create(); err != nil {
		return error2.NewCreateError("profile", err.Error())
	}

	return nil
}
