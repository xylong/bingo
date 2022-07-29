package validators

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

var valid *validator.Validate

func init() {
	fmt.Println("aa")
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		valid = v
	} else {
		logrus.Fatal("error validator ")
	}

	if err := valid.RegisterValidation("phone", CheckPhone); err != nil {
		logrus.Errorf("validator phone error: %s", err.Error())
	}
}
