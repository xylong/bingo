package bingo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Group 路由组
type Group struct {
	*gin.RouterGroup
	group string
}

func NewGroup(routerGroup *gin.RouterGroup) *Group {
	return &Group{RouterGroup: routerGroup}
}

// Group 路由分组
func (g *Group) Group(group string, callback func(ig *Group)) {
	g.group = group + "/"
	callback(g)
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
	g.handle(http.MethodPut, relativePath, handler)
}

// DELETE delete请求
func (g *Group) DELETE(relativePath string, handler interface{}) {
	g.handle(http.MethodDelete, relativePath, handler)
}

func (g *Group) handle(httpMethod, relativePath string, handler interface{}) {
	if f := convert(handler); f != nil {
		g.Handle(httpMethod, strings.Trim(g.group+"/"+relativePath, "/"), f)
	}
}
