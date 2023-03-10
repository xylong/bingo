package iface

import "github.com/gin-gonic/gin"

// Middleware 中间件
type Middleware interface {
	Before(ctx *gin.Context) error
	After(interface{}) (interface{}, error)
}
