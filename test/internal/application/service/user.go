package service

import (
	"github.com/xylong/bingo/test/internal/domain/model/user"
	"github.com/xylong/bingo/test/internal/infrastructure/GormDao"
)

// UserService 用户服务
type UserService struct {
	ud *GormDao.UserDao
}

func NewUserService() *UserService {
	return &UserService{
		ud: GormDao.NewUserDao(),
	}
}

func (s *UserService) Find() *user.User {
	u := user.New(user.WithID(1))
	err := s.ud.GetByID(u)
	if err == nil {
		return u
	}

	return nil
}

func (s *UserService) FindByID(id int64) {

}

func (s *UserService) FindByName(name string) {

}

func (s *UserService) FindByPhone(phone string) {

}

func (s UserService) name() {

}
