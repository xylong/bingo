package service

import (
	"github.com/xylong/bingo/test/internal/application/assembler"
	"github.com/xylong/bingo/test/internal/application/dto"
	"github.com/xylong/bingo/test/internal/domain/model/user"
	"github.com/xylong/bingo/test/internal/domain/repository"
	"github.com/xylong/bingo/test/internal/infrastructure/GormDao"
	"github.com/xylong/bingo/test/internal/lib/db"
)

// UserService 用户服务
type UserService struct {
	Req *assembler.UserReq
	Rep *assembler.UserRep

	UserDao repository.IUser
}

func NewUserService() *UserService {
	return &UserService{Req: &assembler.UserReq{}, Rep: &assembler.UserRep{}, UserDao: GormDao.NewUserDao(db.DB)}
}

func (s *UserService) Find() *user.User {
	u := user.New(user.WithID(1))
	err := s.UserDao.Get(u)
	if err == nil {
		return u
	}

	return nil
}

func (s *UserService) GetSimpleUser(req *dto.SimpleUserReq) *dto.SimpleUser {
	u := s.Req.D2M_User(req)
	_ = s.UserDao.Get(u)
	return s.Rep.M2D_SimpleUser(u)
}
