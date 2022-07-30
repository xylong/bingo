package middleware

import (
	"github.com/xylong/bingo"
	"net/http"
)

// Validate 参数验证中间件
type Validate struct {
}

func NewValidate() *Validate {
	return &Validate{}
}

func (v *Validate) Before(ctx *bingo.Context) error {
	defer func() {
		if err := recover(); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": err,
				"data":    nil,
			})
		}
	}()

	ctx.Next()
	return nil
}
