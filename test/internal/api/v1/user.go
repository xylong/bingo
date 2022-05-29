package v1

import "github.com/xylong/bingo"

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) Register(ctx *bingo.Context) string {
	return "foo"
}
