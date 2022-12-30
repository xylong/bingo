package bingo

import (
	"github.com/gin-gonic/gin"
	"sync"
)

const (
	satellite = "middleware"
)

var ctxPool *sync.Pool

func init() {
	ctxPool = &sync.Pool{
		New: func() interface{} {
			return &Context{}
		},
	}
}

// Context 上下文
type Context struct {
	*gin.Context
}

func bingoContext(context *gin.Context) *Context {
	ctx := ctxPool.Get().(*Context)
	defer ctxPool.Put(ctx)

	ctx.Context = context
	return ctx
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
