package bingo

import (
	"github.com/gin-gonic/gin"
)

// Context 上下文
type Context struct {
	*gin.Context
}

func NewContext(context *gin.Context) *Context {
	return &Context{Context: context}
}

// Token 获取token，默认token键为Authorization
func (c *Context) Token(key ...string) string {
	if len(key) > 0 {
		return c.Request.Header.Get(key[0])
	}
	return c.Request.Header.Get(authKey)
}

// Binding 参数绑定
func (c *Context) Binding(f func(interface{}) error, data interface{}) *ErrorResult {
	return NewErrorResult(data, f(data))
}
