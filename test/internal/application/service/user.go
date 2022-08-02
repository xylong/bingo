package service

import (
	"fmt"
	"github.com/xylong/bingo"
	"github.com/xylong/bingo/test/internal/application/assembler"
	"github.com/xylong/bingo/test/internal/application/dto"
	"github.com/xylong/bingo/test/internal/domain/aggregation"
	"github.com/xylong/bingo/test/internal/domain/model/profile"
	"github.com/xylong/bingo/test/internal/domain/model/user"
	"github.com/xylong/bingo/test/internal/domain/model/userLog"
	GormDao2 "github.com/xylong/bingo/test/internal/infrastructure/dao/GormDao"
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
	err := aggregation.NewMember(aggregation.WithUser(u), aggregation.WithUserRepo(GormDao2.NewUserDao(db.DB))).GetUser()
	if err != nil {
		return nil
	}

	return s.rep.M2D_SimpleUser(u)
}

func (s *UserService) Create(register *dto.SmsRegister) interface{} {
	tx := db.DB.Begin()
	u, ud := user.New(user.WithPhone(register.Phone), user.WithNickName(register.Nickname)), GormDao2.NewUserDao(tx)
	p, pd := profile.New(), GormDao2.NewProfileDao(tx)

	member := aggregation.NewMember(
		aggregation.WithUser(u), aggregation.WithUserRepo(ud),
		aggregation.WithProfile(p), aggregation.WithProfileRepo(pd),
	)
	if err := member.Create(); err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()

		bingo.Task(s.Log, func() {
			fmt.Printf("user:%d 通过短信注册成功", u.ID)
		}, u, userLog.Register, "短信注册")

		return nil
	}
}

// Log 记录用户日志
func (s *UserService) Log(param ...interface{}) {
	fmt.Println(param[0], param[1], param[2], "aaa")
	_ = aggregation.NewMember(
		aggregation.WithLogRepo(GormDao2.NewUserLogDao(db.DB)),
		aggregation.WithLog(
			s.req.D2M_Log(
				param[0].(*user.User),
				param[1].(int),
				param[2].(string)),
		)).AddLog()
}

func (s *UserService) GetList(req *dto.UserReq) (int, string, []*dto.SimpleUser) {
	users, err := aggregation.NewMember(
		aggregation.WithUser(user.New()),
		aggregation.WithUserRepo(GormDao2.NewUserDao(db.DB)),
		aggregation.WithProfileRepo(GormDao2.NewProfileDao(db.DB)),
	).GetUsers(req)

	if err != nil {
		return 0, "", nil
	}

	return 0, "", s.rep.M2D_SimpleList(users)
}

// GetLog 用户日志
func (s *UserService) GetLog(req *dto.UserLogReq) *dto.UserInfo {
	return s.rep.M2D_UserInfo(req,
		aggregation.NewMember(
			aggregation.WithUser(user.New()),
			aggregation.WithUserRepo(GormDao2.NewUserDao(db.DB))))
}
