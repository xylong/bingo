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
		(defaultResponder)(nil),
		(stringResponder)(nil),
		(jsonResponder)(nil),
		(apiResponder)(nil),
	}
}

type (
	apiResponder     func(*Context) (int, string, interface{})
	stringResponder  func(*Context) string
	jsonResponder    func(*Context) any
	defaultResponder func(*Context)
)

func (r apiResponder) Return() gin.HandlerFunc {
	return func(context *gin.Context) {
		if v, exists := context.Get(satellite); exists {
			context.JSON(http.StatusOK, v.(middlewares).handle(bingoContext(context), r).(gin.H))
		} else {
			code, message, data := r(bingoContext(context))
			context.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": message,
				"data":    data,
			})
		}
	}
}

func (r stringResponder) Return() gin.HandlerFunc {
	return func(context *gin.Context) {
		if v, exists := context.Get(satellite); exists {
			context.String(http.StatusOK, v.(middlewares).handle(bingoContext(context), r).(string))
		} else {
			context.String(http.StatusOK, r(bingoContext(context)))
		}
	}
}

func (r jsonResponder) Return() gin.HandlerFunc {
	return func(context *gin.Context) {
		if v, exists := context.Get(satellite); exists {
			context.JSON(http.StatusOK, v.(middlewares).handle(bingoContext(context), r).(any))
		} else {
			context.JSON(http.StatusOK, r(bingoContext(context)))
		}
	}
}

func (r defaultResponder) Return() gin.HandlerFunc {
	return func(context *gin.Context) {
		if v, exists := context.Get(satellite); exists {
			v.(middlewares).handle(bingoContext(context), r)
		} else {
			r(bingoContext(context))
		}
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
