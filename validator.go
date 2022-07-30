package bingo

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

var (
	// Validator 验证器
	v *validator.Validate

	// ValidateMessage 验证消息
	ValidateMessage map[string]string
)

func init() {
	ValidateMessage = make(map[string]string)

	if engine, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v = engine
	} else {
		logrus.Fatal("error validator")
	}
}

// RegisterBindTag 注册自定义验证规则
func RegisterBindTag(tag string, ruleFunc validator.Func) error {
	return v.RegisterValidation(tag, ruleFunc)
}

// message 验证错误信息
func message(err error) {
	if v, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range v {
			if msg, exist := ValidateMessage[fieldError.Tag()]; exist {
				panic(msg)
			}
		}
	}
}
