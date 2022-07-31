package main

import (
	"fmt"
	"github.com/xylong/bingo"
	v1 "github.com/xylong/bingo/test/internal/api/v1"
	"github.com/xylong/bingo/test/internal/middleware"
)

func main() {
	bingo.Init().
		Mount("v1", v1.Controllers...)(middleware.NewLogger(), middleware.NewValidate()).
		Crontab("0/3 * * * * *", func() {
			fmt.Println("ping")
		}).
		Lunch()
}
