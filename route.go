package bingo

import (
	"github.com/gin-gonic/gin"
	"github.com/xylong/bingo/iface"
	"net/http"
	"strings"
)

// Group 路由组
type Group struct {
	*gin.RouterGroup
	group string
	middlewares
}

func NewGroup(routerGroup *gin.RouterGroup) *Group {
	return &Group{RouterGroup: routerGroup}
}

// Without 分组路由排除指定中间件
func (g Group) Without(callback func(*Group), m ...iface.Middleware) {
	var arr middlewares

	for _, middleware := range m {
		for _, m2 := range g.middlewares {
			if middleware != m2 {
				arr = append(arr, m2)
			}
		}
	}

	g.middlewares = arr
	callback(&g)
}

// Group 路由分组
func (g Group) Group(group string, callback func(*Group), middleware ...iface.Middleware) {
	g.attach(middleware...)
	g.group += group + "/"
	callback(&g)
}

// 挂载中间件
// 级联路由由外向内依次挂载中间件
func (g *Group) attach(middlewares ...iface.Middleware) {
loop:
	for _, middleware := range middlewares {
		// 跳过重复中间件
		for _, m := range g.middlewares {
			if middleware == m {
				continue loop
			}
		}

		middleware := middleware // 防止中间件覆盖
		g.Use(func(context *gin.Context) {
			_ = middleware.Before(context)
			context.Next()
		})

		// 记录使用过的中间件
		g.middlewares = append(g.middlewares, middleware)
	}
}

// HEAD head请求
func (g *Group) HEAD(relativePath string, handler interface{}) {
	g.handle(http.MethodHead, relativePath, handler)
}

// GET get请求
func (g *Group) GET(relativePath string, handler interface{}) {
	g.handle(http.MethodGet, relativePath, handler)
}

// POST post请求
func (g *Group) POST(relativePath string, handler interface{}) {
	g.handle(http.MethodPost, relativePath, handler)
}

// PUT put请求
func (g *Group) PUT(relativePath string, handler interface{}) {
	g.handle(http.MethodPut, relativePath, handler)
}

// PATCH patch请求
func (g *Group) PATCH(relativePath string, handler interface{}) {
	g.handle(http.MethodPatch, relativePath, handler)
}

// DELETE delete请求
func (g *Group) DELETE(relativePath string, handler interface{}) {
	g.handle(http.MethodDelete, relativePath, handler)
}

// OPTIONS option请求
func (g *Group) OPTIONS(relativePath string, handler interface{}) {
	g.handle(http.MethodOptions, relativePath, handler)
}

// ANY 任意请求方式
func (g *Group) ANY(relativePath string, handler interface{}) {
	g.HEAD(relativePath, handler)
	g.GET(relativePath, handler)
	g.POST(relativePath, handler)
	g.PUT(relativePath, handler)
	g.PATCH(relativePath, handler)
	g.DELETE(relativePath, handler)
	g.OPTIONS(relativePath, handler)
}

func (g *Group) handle(httpMethod, relativePath string, handler interface{}) {
	if f := convert(handler); f != nil {
		g.Handle(httpMethod, strings.Trim(g.group+"/"+relativePath, "/"), f)
	}
}
