package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xylong/bingo"
	"github.com/xylong/bingo/test/internal/domain/aggregation"
	"github.com/xylong/bingo/test/internal/domain/model/profile"
	"github.com/xylong/bingo/test/internal/domain/model/user"
	"github.com/xylong/bingo/test/internal/dto"
	"github.com/xylong/bingo/test/internal/infrastructure/GormDao"
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

	tx := db.DB.Begin()
	up, pp := GormDao.NewUserDao(tx), GormDao.NewProfileDao(tx)
	u := user.New(user.WithPhone(form.Phone))
	p := profile.New(profile.WithPassword(form.Password))

	agg := aggregation.NewFrontUserAgg(u, p, up, pp)
	if err := agg.CreateUser(); err == nil {
		return 0, "", agg.User
	} else {
		return 400, err.Error(), nil
	}
}

func (c *UserController) login(ctx *bingo.Context) string {
	return "login"
}

func (c *UserController) logout(ctx *bingo.Context) {
	fmt.Println("a")
}

func (c *UserController) me(ctx *bingo.Context) bingo.Json {
	up, pp := GormDao.NewUserDao(db.DB), GormDao.NewProfileDao(db.DB)
	u, p := user.New(), profile.New()

	u.ID = 1
	p.UserID = u.ID

	agg := aggregation.NewFrontUserAgg(u, p, up, pp)
	if err := agg.Get(); err == nil {
		agg.User.Profile = agg.Profile
		return agg.User
	} else {
		return gin.H{"error": err.Error()}
	}
}

func (c *UserController) update(ctx *bingo.Context) string {
	return "update"
}

func (c *UserController) profile(ctx *bingo.Context) bingo.Json {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return gin.H{"error": err.Error()}
	}

	up, pp := GormDao.NewUserDao(db.DB), GormDao.NewProfileDao(db.DB)
	u, p := user.New(user.WithID(id)), profile.New(profile.WithUserID(id))

	agg := aggregation.NewFrontUserAgg(u, p, up, pp)
	if err := agg.Get(); err == nil {
		agg.User.Profile = agg.Profile
		return agg.User
	} else {
		return err
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
