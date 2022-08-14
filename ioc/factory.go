package ioc

import (
	"github.com/shenyisyn/goft-expr/src/expr"
	"reflect"
)

// injectTag 注入标签
const injectTag = "inject"

func init() {
	Factory = NewMapperFactory()
}

// Factory ioc全局变量
var Factory *MapperFactory

// MapperFactory ioc容器工厂
type MapperFactory struct {
	mapper
	expr map[string]interface{} // 表达式
}

func NewMapperFactory() *MapperFactory {
	return &MapperFactory{
		mapper: make(map[reflect.Type]reflect.Value),
		expr:   make(map[string]interface{}),
	}
}

// Set 设置，值必须是指针
func (f *MapperFactory) Set(item ...interface{}) {
	if item == nil || len(item) == 0 {
		return
	}

	for _, i := range item {
		f.mapper.set(i)
	}
}

// Get 获取
func (f *MapperFactory) Get(key interface{}) interface{} {
	if key != nil {
		if v := f.mapper.get(key); v.IsValid() {
			return v.Interface()
		}
	}

	return nil
}

// Apply 处理依赖注入
// 接受struct的指针，并且注入的字段需设置inject标签，如`inject:"-"`或`inject:"Service.Order()"`
// -为单例模式，表达式为多例模式
func (f *MapperFactory) Apply(obj interface{}) {
	if obj == nil {
		return
	}

	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag.Get(injectTag)

		if v.Field(i).CanSet() && tag != "" {
			if tag != "-" {
				// 多例模式
				if result := expr.BeanExpr(tag, f.expr); result != nil && !result.IsEmpty() && result[0] != nil {
					v.Field(i).Set(reflect.ValueOf(result[0]))
					f.Apply(result[0])
				}
			} else {
				// 单例模式
				if val := f.Get(field.Type); val != nil {
					v.Field(i).Set(reflect.ValueOf(val))
					f.Apply(val)
				}
			}
		}
	}
}

// Unwrap 加载注入实体，并自动构建表达式
func (f *MapperFactory) Unwrap(bean ...interface{}) {
	for _, b := range bean {
		t := reflect.TypeOf(b)

		if t.Kind() != reflect.Ptr {
			panic("required ptr object")
		}

		if t.Elem().Kind() != reflect.Struct {
			continue
		}

		f.Set(b)                    // 将自身加入mapper
		f.expr[t.Elem().Name()] = b // 自动构建expr
		f.Apply(b)                  // 处理依赖注入(new)

		v := reflect.ValueOf(b)

		for i := 0; i < t.NumMethod(); i++ {
			method := v.Method(i)
			result := method.Call(nil)

			if result != nil && len(result) == 1 {
				f.Set(result[0].Interface())
			}
		}
	}
}

// GetMapper 获取容器
func (f *MapperFactory) GetMapper() map[reflect.Type]reflect.Value {
	return f.mapper
}
