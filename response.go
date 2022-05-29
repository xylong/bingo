package bingo

import (
	"github.com/gin-gonic/gin"
	"github.com/xylong/bingo/iface"
	"net/http"
	"reflect"
)

var responderList []iface.Responder

func init() {
	responderList = []iface.Responder{
		(stringResponder)(nil),
		(jsonResponder)(nil),
		(defaultResponder)(nil),
	}
}

type (
	Json interface{}

	stringResponder  func(*Context) string
	jsonResponder    func(*Context) Json
	defaultResponder func(*Context)
)

func (r stringResponder) Return() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(http.StatusOK, r(NewContext(context)))
	}
}

func (r jsonResponder) Return() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, r(NewContext(context)))
	}
}

func (r defaultResponder) Return() gin.HandlerFunc {
	return func(context *gin.Context) {
		r(NewContext(context))
	}
}

// convert 将路由函数转为gin的HandlerFunc
func convert(handler interface{}) gin.HandlerFunc {
	value := reflect.ValueOf(handler)

	for _, responder := range responderList {
		t := reflect.TypeOf(responder)
		if value.Type().ConvertibleTo(t) {
			return value.Convert(t).Interface().(iface.Responder).Return()
		}
	}

	return nil
}
