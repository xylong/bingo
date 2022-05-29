package iface

import "github.com/gin-gonic/gin"

// Responder 响应器
type Responder interface {
	Return() gin.HandlerFunc
}
