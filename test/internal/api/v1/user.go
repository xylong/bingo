package v1

import (
	"fmt"
	"github.com/xylong/bingo"
	"github.com/xylong/bingo/test/internal/application/dto"
	"github.com/xylong/bingo/test/internal/application/service"
	"github.com/xylong/bingo/test/internal/infrastructure/utils"
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
	form := &dto.SmsRegister{}
	if err := ctx.ShouldBind(form); err != nil {
		return err.Error()
	}

	return c.service.Create(form)
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
		utils.Exec(ctx.ShouldBindUri, &dto.SimpleUserReq{}).
			Unwrap().(*dto.SimpleUserReq))
}

func (c *UserController) index(ctx *bingo.Context) (int, string, interface{}) {
	req := &dto.UserReq{}
	if err := ctx.ShouldBind(req); err != nil {
		return 400, err.Error(), nil
	}

	return c.service.GetList(req)
}

func (c *UserController) log(ctx *bingo.Context) any {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return nil
	}

	req := &dto.UserLogReq{}
	req.ID = id

	return c.service.GetLog(
		utils.Exec(ctx.ShouldBind, req).
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
