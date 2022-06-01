package main

import (
	"github.com/xylong/bingo"
	v1 "github.com/xylong/bingo/test/internal/api/v1"
)

func main() {
	bingo.Init().
		Mount("v1", v1.NewUserController()).
		Lunch()
}
