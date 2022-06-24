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
		(anyResponder)(nil),
		(defaultResponder)(nil),
	}
}

type (
	Json interface{}

	stringResponder  func(*Context) string
	errorResponder   func(*Context) error
	anyResponder     func(*Context) any
	defaultResponder func(*Context)
)

func (r stringResponder) Return() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(http.StatusOK, r(NewContext(context)))
	}
}

func (r errorResponder) Return() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := r(NewContext(context))
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (r anyResponder) Return() gin.HandlerFunc {
	return func(context *gin.Context) {
		data := r(NewContext(context))

		switch value := data.(type) {
		case string:
			context.String(http.StatusOK, value)
		case error:
			context.JSON(http.StatusBadRequest, gin.H{
				"error": value.Error(),
			})
		default:
			context.JSON(http.StatusOK, value)
		}
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
