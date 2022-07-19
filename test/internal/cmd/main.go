package main

import (
	"github.com/xylong/bingo"
	v1 "github.com/xylong/bingo/test/internal/api/v1"
	"github.com/xylong/bingo/test/internal/middleware"
)

func main() {
	bingo.Init().
		Mount("v1", v1.Controllers...)(middleware.NewLogger()).
		Lunch()
}
