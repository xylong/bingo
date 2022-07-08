package service

import (
	"github.com/xylong/bingo/test/internal/application/assembler"
	"github.com/xylong/bingo/test/internal/application/dto"
	"github.com/xylong/bingo/test/internal/domain/model/user"
	"github.com/xylong/bingo/test/internal/infrastructure/GormDao"
)

// UserService 用户服务
type UserService struct {
	ud  *GormDao.UserDao
	Req *assembler.UserReq
	Rep *assembler.UserRep
}

func (s *UserService) Find() *user.User {
	u := user.New(user.WithID(1))
	err := s.ud.Get(u)
	if err == nil {
		return u
	}

	return nil
}

func (s *UserService) GetSimpleUser(req *dto.SimpleUserReq) *dto.SimpleUser {
	u := s.Req.D2M_User(req)
	_ = s.ud.Get(u)
	return s.Rep.M2D_SimpleUser(u)
}
