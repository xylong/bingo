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
	fmt.Println("before auth")
	token := ctx.Token()
	if token == "" {
		return fmt.Errorf("unauthorized")
	}
	return nil
}

func (a *Authorization) After(data interface{}) (interface{}, error) {
	fmt.Println("after auth")
	return nil, nil
}
