package main

import (
	"github.com/xylong/bingo"
	v1 "github.com/xylong/bingo/test/internal/api/v1"
	"github.com/xylong/bingo/test/internal/middleware"
)

func main() {
	bingo.Init().
		Mount("v1", []bingo.Middleware{middleware.NewAuthorization()}, v1.NewUserController()).
		Lunch()
}
