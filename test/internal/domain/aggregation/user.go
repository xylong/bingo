package aggregation

import (
	"fmt"
	"github.com/xylong/bingo/test/internal/domain/model/profile"
	"github.com/xylong/bingo/test/internal/domain/model/repository"
	"github.com/xylong/bingo/test/internal/domain/model/user"
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
	return fu
}

func (u *FrontUserAgg) Get() error {
	if u.User.ID <= 0 {
		return fmt.Errorf("root error: %s", "user model's is zero")
	}

	if err := u.User.Get(); err != nil {
		return fmt.Errorf("user data error: %s", err.Error())
	}

	if err := u.Profile.Get(); err != nil {
		return fmt.Errorf("profile data error:%s", err.Error())
	}

	return nil
}
