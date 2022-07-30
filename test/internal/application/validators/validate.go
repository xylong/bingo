package validators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

var (
	valid   *validator.Validate
	message map[string]string
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		valid = v
	} else {
		logrus.Fatal("error validator ")
	}

	message = make(map[string]string)

	register("phone", CheckPhone)
	register("ID", CheckID)
	register("nickname", Nickname("required,min=2,max=10").toFunc())
}

func register(rule string, fn validator.Func) {
	if err := valid.RegisterValidation(rule, fn); err != nil {
		logrus.Errorf("register %s tag error: %s", rule, err.Error())
	}
}

// ValidateMessage 绑定验证消息
func ValidateMessage(err error) {
	if v, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range v {
			if msg, exist := message[fieldError.Tag()]; exist {
				panic(msg)
			}
		}
	}
}
