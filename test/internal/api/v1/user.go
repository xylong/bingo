package v1

import (
	"fmt"
	"github.com/xylong/bingo"
	"github.com/xylong/bingo/test/internal/dto"
	"github.com/xylong/bingo/test/internal/middleware"
	"github.com/xylong/bingo/test/internal/model/user"
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

	group.POST("register", bingo.NewBind[dto.RegisterForm]().Try(c.register).Catch().Complete())
	group.POST("login", c.login)
	group.DELETE("logout", c.logout)
}

func (c *UserController) register(ctx *bingo.Context, form *dto.RegisterForm) bingo.Json {
	//ctx.JSON(http.StatusOK, form)
	return form
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
