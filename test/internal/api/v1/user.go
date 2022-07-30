package v1

import (
	"fmt"
	"github.com/xylong/bingo"
	"github.com/xylong/bingo/test/internal/application/dto"
	"github.com/xylong/bingo/test/internal/application/service"
	"github.com/xylong/bingo/test/internal/middleware"
	"strconv"
)

func init() {
	RegisterController(NewUserController())
}

type UserController struct {
	// todo 注入
	service *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		service: service.NewUserService(),
	}
}

func (c *UserController) smsRegister(ctx *bingo.Context) interface{} {
	return c.service.Create(
		ctx.Binding(ctx.ShouldBind, &dto.SmsRegister{}).
			Unwrap().(*dto.SmsRegister))
}

func (c *UserController) login(ctx *bingo.Context) string {
	return "login"
}

func (c *UserController) logout(ctx *bingo.Context) {
	fmt.Println("a")
}

func (c *UserController) me(ctx *bingo.Context) any {
	return "b"
}

func (c *UserController) update(ctx *bingo.Context) string {
	return "update"
}

func (c *UserController) show(ctx *bingo.Context) any {
	return c.service.GetSimpleUser(
		ctx.Binding(ctx.ShouldBindUri, &dto.SimpleUserReq{}).
			Unwrap().(*dto.SimpleUserReq))
}

func (c *UserController) index(ctx *bingo.Context) (int, string, interface{}) {
	return c.service.GetList(
		ctx.Binding(ctx.ShouldBind, &dto.UserReq{}).
			Unwrap().(*dto.UserReq))
}

func (c *UserController) log(ctx *bingo.Context) any {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return nil
	}

	req := &dto.UserLogReq{}
	req.ID = id

	return c.service.GetLog(
		ctx.Binding(ctx.ShouldBind, req).
			Unwrap().(*dto.UserLogReq))
}

func (c *UserController) Route(group *bingo.Group) {
	group.Group("", func(users *bingo.Group) {
		users.GET("me", c.me)
		users.PUT("me", c.update)
		users.GET("users/:id", c.show)
		users.GET("users/:id/logs", c.log)
	}, middleware.NewAuthorization())

	group.POST("register", c.smsRegister)
	group.POST("login", c.login)
	group.DELETE("logout", c.logout)

	group.GET("users", c.index)

}
