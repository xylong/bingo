package main

import (
	"github.com/xylong/bingo"
	v1 "github.com/xylong/bingo/test/internal/api/v1"
	v2 "github.com/xylong/bingo/test/internal/api/v2"
	"github.com/xylong/bingo/test/internal/interface/config"
	"github.com/xylong/bingo/test/internal/middleware"
)

func main() {
	bingo.Init("conf", "config").
		Inject(config.NewAdapter(), config.NewService()).
		Mount("v1", v1.Controllers...)(middleware.NewLogger(), middleware.NewValidate()).
		Mount("v2", v2.Controllers...)().
		//Crontab("0/3 * * * * *", ".MockController.Foo").
		Lunch()

}
