package ioc

import "reflect"

type mapper map[reflect.Type]reflect.Value

func (m mapper) set(obj interface{}) {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Ptr {
		panic("required ptr")
	}

	m[t] = reflect.ValueOf(obj)
}

func (m mapper) get(key interface{}) reflect.Value {
	var t reflect.Type

	if v, ok := key.(reflect.Type); ok {
		t = v
	} else {
		t = reflect.TypeOf(key)
	}

	if v, ok := m[t]; ok {
		return v
	}

	// 接口处理
	for k, v := range m {
		if t.Kind() == reflect.Interface && k.Implements(t) {
			return v
		}
	}

	return reflect.Value{}
}
