package bingo

import (
	"fmt"
	"reflect"
	"strings"
)

func init() {
	Annotations = make([]Annotation, 0)
	Annotations = append(Annotations, new(Value))
}

var Annotations []Annotation

// Annotation 注解
type Annotation interface {
	SetTag(tag reflect.StructTag)
	String() string
}

// IsAnnotation 判断注入对象是否为注解
func IsAnnotation(p reflect.Type) bool {
	for _, a := range Annotations {
		if reflect.TypeOf(a) == p {
			return true
		}
	}

	return false
}

// Value 注解的值
// 满足`prefix:"user.age"`格式的用户自定义配置
type Value struct {
	tag reflect.StructTag
	*BeanFactory
}

func (v *Value) SetTag(tag reflect.StructTag) {
	v.tag = tag
}

// String 将注解的值转为字符串
func (v *Value) String() string {
	prefix := v.tag.Get("prefix")
	if prefix == "" {
		return ""
	}

	arr := strings.Split(prefix, ".")
	if config := v.BeanFactory.GetBean(new(Config)); config != nil {
		if value := GetConfig(config.(*Config).Custom, arr, 0); value != nil {
			return fmt.Sprintf("%v", value)
		}
	}

	return ""
}
