package service

import (
	"fmt"
	"github.com/xylong/bingo/test/internal/domain/repository"
	"github.com/xylong/bingo/test/internal/infrastructure/utils"
)

// LoginService 登录
type LoginService struct {
	userDao    repository.IUser
	profileDao repository.IProfile
}

// Login 登录
func (s *LoginService) Login(phone, password string) (string, error) {
	user := s.userDao.GetByPhone(phone)

	if user.ID == 0 {
		return "10000", fmt.Errorf("用户不存在")
	}

	profile, err := s.profileDao.GetByUser(user)
	if err != nil {
		return "10001", fmt.Errorf("用户信息不存在")
	}

	if false == utils.CheckPassword(profile.Password, password) {
		return "10002", fmt.Errorf("密码错误")

	}

	return "token", nil
}
