package bingo

import "reflect"

// Bean 注入对象
type Bean interface {
	Name() string
}

type BeanFactory struct {
	beans []Bean
}

func NewBeanFactory() *BeanFactory {
	bf := &BeanFactory{
		beans: make([]Bean, 0),
	}

	bf.setBean(bf)

	return bf
}

func (b *BeanFactory) setBean(beans ...Bean) {
	b.beans = append(b.beans, beans...)
}

func (b *BeanFactory) getBean(p reflect.Type) interface{} {
	for _, bean := range b.beans {
		if p == reflect.TypeOf(bean) {
			return bean
		}
	}

	return nil
}

// GetBean 获取bean
func (b *BeanFactory) GetBean(bean interface{}) interface{} {
	return b.getBean(reflect.TypeOf(bean))
}

// 将bean注入控制器
func (b *BeanFactory) inject(controller Controller) {
	t, v := reflect.TypeOf(controller).Elem(), reflect.ValueOf(controller).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() != reflect.Ptr || !field.IsNil() {
			continue
		}

		// 注解的处理
		if IsAnnotation(field.Type()) {
			field.Set(reflect.New(field.Type().Elem()))
			field.Interface().(Annotation).SetTag(t.Field(i).Tag)
			b.Inject(field.Interface())
			continue
		}

		if P := b.getBean(field.Type()); P != nil {
			field.Set(reflect.New(field.Type().Elem()))
			field.Elem().Set(reflect.ValueOf(P).Elem())
		}
	}
}

// Inject 注入
func (b *BeanFactory) Inject(obj interface{}) {
	v := reflect.ValueOf(obj)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() != reflect.Ptr || !field.IsNil() {
			continue
		}

		if p := b.getBean(field.Type()); p != nil && field.CanInterface() {
			field.Set(reflect.New(field.Type().Elem()))
			field.Elem().Set(reflect.ValueOf(p).Elem())
		}
	}
}

func (b *BeanFactory) Name() string {
	return "BeanFactory"
}
