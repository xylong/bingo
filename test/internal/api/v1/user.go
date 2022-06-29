package v1

import (
	"fmt"
	"github.com/xylong/bingo"
	"github.com/xylong/bingo/test/internal/domain/model/user"
	"github.com/xylong/bingo/test/internal/dto"
	"github.com/xylong/bingo/test/internal/middleware"
	"time"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) Route(group *bingo.Group) {
	group.Group("", func(users *bingo.Group) {
		users.GET("me", c.me)
		users.PUT("me", c.update)
		users.GET("friends", c.friend)
	}, middleware.NewAuthorization())

	group.POST("register", c.register)
	group.POST("login", c.login)
	group.DELETE("logout", c.logout)
}

func (c *UserController) register(ctx *bingo.Context) (int, string, interface{}) {
	form := &dto.RegisterForm{}

	err := ctx.ShouldBind(form)
	if err != nil {
		return 400, "参数错误", nil
	}

	return 0, "", "hello"
}

func (c *UserController) login(ctx *bingo.Context) string {
	return "login"
}

func (c *UserController) logout(ctx *bingo.Context) {
	fmt.Println("a")
}

func (c *UserController) me(ctx *bingo.Context) bingo.Json {
	return user.User{
		ID:        int(time.Now().Unix()),
		Nickname:  "summer",
		CreatedAt: time.Time{},
	}
}

func (c *UserController) update(ctx *bingo.Context) string {
	return "update"
}

func (c *UserController) friend(ctx *bingo.Context) string {
	return "friend"
}
