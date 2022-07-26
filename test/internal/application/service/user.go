package service

import (
	"github.com/xylong/bingo/test/internal/application/assembler"
	"github.com/xylong/bingo/test/internal/application/dto"
	"github.com/xylong/bingo/test/internal/domain/aggregation"
	"github.com/xylong/bingo/test/internal/domain/model/profile"
	"github.com/xylong/bingo/test/internal/domain/model/user"
	"github.com/xylong/bingo/test/internal/infrastructure/GormDao"
	"github.com/xylong/bingo/test/internal/lib/db"
)

// UserService 用户服务
type UserService struct {
	req *assembler.UserReq
	rep *assembler.UserRep
}

func NewUserService() *UserService {
	return &UserService{req: &assembler.UserReq{}, rep: &assembler.UserRep{}}
}

func (s *UserService) Find() *user.User {
	u := user.New(user.WithID(1))

	return u
}

func (s *UserService) GetSimpleUser(req *dto.SimpleUserReq) *dto.SimpleUser {
	u := s.req.D2M_User(req)

	return s.rep.M2D_SimpleUser(u)
}

func (s *UserService) Create(register *dto.SmsRegister) interface{} {
	tx := db.DB.Begin()
	u, ud := user.New(user.WithPhone(register.Phone), user.WithNickName(register.Nickname)), GormDao.NewUserDao(tx)
	p, pd := profile.New(), GormDao.NewProfileDao(tx)

	member := aggregation.NewMember(
		aggregation.WithUser(u), aggregation.WithUserRepo(ud),
		aggregation.WithProfile(p), aggregation.WithProfileRepo(pd),
		aggregation.WithLogRepo(GormDao.NewUserLogDao(tx)),
	)
	if err := member.Create(); err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
		return nil
	}
}

func (s *UserService) GetList(req *dto.UserReq) (int, string, []*dto.SimpleUser) {
	users, err := aggregation.NewMember(
		aggregation.WithUser(user.New()),
		aggregation.WithUserRepo(GormDao.NewUserDao(db.DB)),
		aggregation.WithProfileRepo(GormDao.NewProfileDao(db.DB)),
	).GetUsers(req)

	if err != nil {
		return 0, "", nil
	}

	return 0, "", s.rep.M2D_SimpleList(users)
}
