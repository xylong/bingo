package v1

import (
	"fmt"
	"github.com/xylong/bingo"
	"github.com/xylong/bingo/test/internal/application/service"
	"github.com/xylong/bingo/test/internal/middleware"
	"strconv"
)

func init() {
	RegisterController(NewUserController())
}

type UserController struct {
	service *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		service: service.NewUserService(),
	}
}

func (c *UserController) register(ctx *bingo.Context) (int, string, interface{}) {
	//form := &dto.UserRegister{}
	//
	//err := ctx.ShouldBind(form)
	//if err != nil {
	//	return 400, "参数错误", nil
	//}
	//
	//tx := db.DB.Begin()
	//up, pp := GormDao.NewUserDao(tx), GormDao.NewProfileDao(tx)
	//u := user.New(user.WithPhone(form.Phone), user.WithNickName(form.Nickname))
	//p := profile.New(profile.WithPassword(form.Password))
	//
	//agg := aggregation.NewFrontUserAgg(u, p, up, pp)
	//if err := agg.CreateUser(); err == nil {
	//	tx.Commit()
	//	return 0, "", agg.User
	//} else {
	//	tx.Rollback()
	//	return 400, err.Error(), nil
	//}
}

func (c *UserController) login(ctx *bingo.Context) string {
	return "login"
}

func (c *UserController) logout(ctx *bingo.Context) {
	fmt.Println("a")
}

func (c *UserController) me(ctx *bingo.Context) bingo.Json {
	//up, pp := GormDao.NewUserDao(db.DB), GormDao.NewProfileDao(db.DB)
	//u, p := user.New(), profile.New()
	//
	//u.ID = 1
	//p.UserID = u.ID
	//
	//agg := aggregation.NewFrontUserAgg(u, p, up, pp)
	//if err := agg.Get(); err == nil {
	//	return agg.User
	//} else {
	//	return gin.H{"error": err.Error()}
	//}
}

func (c *UserController) update(ctx *bingo.Context) string {
	return "update"
}

func (c *UserController) profile(ctx *bingo.Context) bingo.Json {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return nil
	}

	return nil
}

func (c *UserController) Route(group *bingo.Group) {
	group.Group("", func(users *bingo.Group) {
		users.GET("me", c.me)
		users.PUT("me", c.update)
		users.GET("users/:id", c.profile)
	}, middleware.NewAuthorization())

	group.POST("register", c.register)
	group.POST("login", c.login)
	group.DELETE("logout", c.logout)
}
