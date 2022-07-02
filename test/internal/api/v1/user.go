package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xylong/bingo"
	"github.com/xylong/bingo/test/internal/domain/model/user"
	"github.com/xylong/bingo/test/internal/dto"
	"github.com/xylong/bingo/test/internal/infrastructure/GromDao"
	"github.com/xylong/bingo/test/internal/lib/db"
	"github.com/xylong/bingo/test/internal/middleware"
	"strconv"
)

func init() {
	RegisterController(NewUserController())
}

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
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
	return user.New(
		user.WithNickName("静静"),
		user.WithPhone("19987654320"),
		user.WithUnionid("abcd1234"),
	)
}

func (c *UserController) update(ctx *bingo.Context) string {
	return "update"
}

func (c *UserController) profile(ctx *bingo.Context) bingo.Json {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return gin.H{"error": err.Error()}
	}

	repo := GromDao.NewUserRepo(db.DB)
	u := user.New(user.WithRepo(repo))
	u.ID = id

	if err := u.Get(); err != nil {
		return gin.H{"error": err.Error()}
	} else {
		return u
	}
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
