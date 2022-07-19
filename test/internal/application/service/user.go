package service

import (
	"github.com/xylong/bingo/test/internal/application/assembler"
	"github.com/xylong/bingo/test/internal/application/dto"
	"github.com/xylong/bingo/test/internal/domain/aggregation"
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
	//err := s.UserDao.Get(u)

	return u
}

func (s *UserService) GetSimpleUser(req *dto.SimpleUserReq) *dto.SimpleUser {
	u := s.Req.D2M_User(req)
	//_ = s.UserDao.Get(u)
	return s.Rep.M2D_SimpleUser(u)
}

func (s *UserService) Create(register *dto.SmsRegister) interface{} {
	tx := db.DB.Begin()
	u := user.New(user.WithPhone(register.Phone), user.WithNickName(register.Nickname))
	ud := GormDao.NewUserDao(tx)
	pd := GormDao.NewProfileDao(tx)

	member := aggregation.NewMember(aggregation.WithUser(u), aggregation.WithUserRepo(ud), aggregation.WIthProfileRepo(pd))
	if err := member.Create(); err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
		return nil
	}
}
