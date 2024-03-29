package bingo

import (
	"github.com/gin-gonic/gin"
	"github.com/xylong/bingo/iface"
)

type middlewares []iface.Middleware

func (m middlewares) remove(me iface.Middleware) {
	index := 0

	for _, middleware := range m {
		if middleware != me {
			m[index] = middleware
			index++
		}
	}
}

//func (m middlewares) before(ctx *gin.Context) {
//	for _, f := range m {
//		if err := f.Before(ctx); err != nil {
//			panic(err)
//		}
//	}
//}

func (m middlewares) after(ctx *gin.Context, data interface{}) interface{} {
	for i := len(m) - 1; i >= 0; i-- {
		if result, err := m[i].After(data); err != nil {
			panic(err)
		} else {
			data = result
		}
	}

	return data
}

// handle 处理前置中间件和后置中间件
func (m middlewares) handle(ctx *gin.Context, responder iface.Responder) (result interface{}) {
	//m.before(ctx)

	switch responder.(type) {
	case apiResponder:
		code, message, data := responder.(apiResponder)(ctx)
		result = gin.H{"code": code, "message": message, "data": data}
	case stringResponder:
		result = responder.(stringResponder)(ctx)
	case jsonResponder:
		result = responder.(jsonResponder)(ctx)
	default:
		responder.(defaultResponder)(ctx)
	}

	return m.after(ctx, result)
}
