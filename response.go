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
	apiResponder     func(*gin.Context) (int, string, interface{})
	stringResponder  func(*gin.Context) string
	jsonResponder    func(*gin.Context) any
	defaultResponder func(*gin.Context)
)

func (r apiResponder) Return() gin.HandlerFunc {
	return func(context *gin.Context) {
		if v, exists := context.Get("middleware"); exists {
			context.JSON(http.StatusOK, v.(middlewares).handle(context, r).(gin.H))
		} else {
			code, message, data := r(context)
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
		if v, exists := context.Get("middleware"); exists {
			context.String(http.StatusOK, v.(middlewares).handle(context, r).(string))
		} else {
			context.String(http.StatusOK, r(context))
		}
	}
}

func (r jsonResponder) Return() gin.HandlerFunc {
	return func(context *gin.Context) {
		if v, exists := context.Get("middleware"); exists {
			context.JSON(http.StatusOK, v.(middlewares).handle(context, r).(any))
		} else {
			context.JSON(http.StatusOK, r(context))
		}
	}
}

func (r defaultResponder) Return() gin.HandlerFunc {
	return func(context *gin.Context) {
		if v, exists := context.Get("middleware"); exists {
			v.(middlewares).handle(context, r)
		} else {
			r(context)
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
