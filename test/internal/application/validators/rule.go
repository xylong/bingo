package validators

import (
	"github.com/go-playground/validator/v10"
)

var CheckPhone validator.Func = func(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(string)
	return ok && len(value) == 11
}
