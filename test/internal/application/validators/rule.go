package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/xylong/bingo"
	"regexp"
)

var CheckPhone validator.Func = func(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(string)

	return ok && regexp.MustCompile("^1[345789]{1}\\d{9}$").MatchString(value)
}

var CheckID validator.Func = func(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(string)

	return ok && regexp.MustCompile("(^\\d{15}$)|(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)").MatchString(value)
}

type Nickname string

func (n Nickname) toFunc() validator.Func {
	bingo.ValidateMessage["nickname"] = "昵称必须在2-10位之间"

	return func(fl validator.FieldLevel) bool {
		value, ok := fl.Field().Interface().(string)
		return ok && n.validate(value)
	}
}

func (n Nickname) validate(nickname string) bool {
	return valid.Var(nickname, string(n)) == nil
}
