package main

import (
	"github.com/xylong/bingo"
	v1 "github.com/xylong/bingo/test/internal/api/v1"
)

var (
	userController *v1.UserController
)

func init() {
	userController = v1.NewUserController()
}

func main() {
	bingo.Init().
		Route("v1", func(group *bingo.Group) {
			group.POST("register", userController.Register)
		}).Lunch()
}
