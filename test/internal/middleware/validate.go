package middleware

import (
	"fmt"
	"github.com/xylong/bingo"
)

// Validate 参数验证中间件
type Validate struct {
}

func NewValidate() *Validate {
	return &Validate{}
}

func (v *Validate) Before(ctx *bingo.Context) error {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
	return nil
}
