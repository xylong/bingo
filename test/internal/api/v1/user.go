package v1

import "github.com/xylong/bingo"

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) Route(group *bingo.Group) {
	group.POST("register", c.register)
	group.POST("login", c.login)
	group.DELETE("logout", c.logout)
	group.GET("me", c.me)
	group.PUT("me", c.update)
	group.GET("friends", c.friend)
}

func (c *UserController) register(ctx *bingo.Context) string {
	return "foo"
}

func (c *UserController) login(ctx *bingo.Context) string {
	return "login"
}

func (c *UserController) logout(ctx *bingo.Context) string {
	return "logout"
}

func (c *UserController) me(ctx *bingo.Context) string {
	return "me"
}

func (c *UserController) update(ctx *bingo.Context) string {
	return "update"
}

func (c *UserController) friend(ctx *bingo.Context) string {
	return "friend"
}
