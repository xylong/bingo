package bingo

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en2 "github.com/go-playground/validator/v10/translations/en"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
)

const (
	local = "zh"
)

var (
	// Validator 验证器
	v *validator.Validate

	// trans 翻译器
	trans ut.Translator

	// ValidateMessage 验证消息
	ValidateMessage map[string]string
)

func init() {
	var err error

	ValidateMessage = make(map[string]string)

	// 修改gin框架中的Validator引擎属性，实现自定制
	if engine, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT, enT := zh.New(), en.New()
		uni := ut.New(enT, zhT) // 第一个备用

		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		if trans, ok = uni.GetTranslator(local); !ok {
			panic(fmt.Errorf("uni.GetTranslator(%s) failed", local))
		}

		switch local {
		case "zh":
			err = zh2.RegisterDefaultTranslations(engine, trans)
		default:
			err = en2.RegisterDefaultTranslations(engine, trans)
		}

		if err != nil {
			zap.L().Error("register default translation error")
		}

		v = engine
	} else {
		zap.L().Error("error validator")
	}
}

// RegisterBindTag 注册自定义验证规则
func RegisterBindTag(tag string, ruleFunc validator.Func) error {
	return v.RegisterValidation(tag, ruleFunc)
}

// message 验证错误信息
func message(err error) string {
	if v, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range v {
			if msg, exist := ValidateMessage[fieldError.Tag()]; exist {
				return msg
			} else {
				return fieldError.Translate(trans)
			}
		}
	}

	return err.Error()
}
