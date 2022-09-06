package validators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/xylong/bingo"
	"go.uber.org/zap"
)

var (
	valid *validator.Validate
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		valid = v
	} else {
		zap.L().Fatal("error validator")
	}

	bingo.RegisterBindTag("phone", CheckPhone)
	bingo.RegisterBindTag("ID", CheckID)
	bingo.RegisterBindTag("nickname", Nickname("required,min=2,max=10").toFunc())
}
