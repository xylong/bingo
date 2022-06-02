package middleware

import (
	"fmt"
	"github.com/xylong/bingo"
)

// Authorization 认证
type Authorization struct {
}

func NewAuthorization() *Authorization {
	return &Authorization{}
}

func (a *Authorization) Before(ctx *bingo.Context) error {
	token := ctx.Token()
	fmt.Println(token, "测试")
	return nil
}
