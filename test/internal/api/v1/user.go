package v1

import (
	"fmt"
	"github.com/xylong/bingo"
	"github.com/xylong/bingo/test/internal/application/dto"
	"github.com/xylong/bingo/test/internal/application/service"
	"github.com/xylong/bingo/test/internal/middleware"
)

func init() {
	RegisterController(NewUserController())
}

type UserController struct {
	// todo 注入
	service *service.UserService
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) register(ctx *bingo.Context) (int, string, interface{}) {
	return 0, "", nil
}

func (c *UserController) login(ctx *bingo.Context) string {
	return "login"
}

func (c *UserController) logout(ctx *bingo.Context) {
	fmt.Println("a")
}

func (c *UserController) me(ctx *bingo.Context) any {
	return nil
}

func (c *UserController) update(ctx *bingo.Context) string {
	return "update"
}

func (c *UserController) show(ctx *bingo.Context) any {
	param := &dto.SimpleUserReq{}
	if err := ctx.ShouldBindUri(param); err != nil {
		return nil
	}

	return c.service.GetSimpleUser(param)
}

func (c *UserController) Route(group *bingo.Group) {
	group.Group("", func(users *bingo.Group) {
		users.GET("me", c.me)
		users.PUT("me", c.update)
		users.GET("users/:id", c.show)
	}, middleware.NewAuthorization())

	group.POST("register", c.register)
	group.POST("login", c.login)
	group.DELETE("logout", c.logout)
}
