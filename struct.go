package utilize

import (
	"reflect"
)

// SetValuesByTag 设置结构体指针中的指定字段值
func SetValuesByTag(obj any, tag string, values map[string]any) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		if value, ok := values[t.Field(i).Tag.Get(tag)]; ok {
			v.Field(i).Set(reflect.ValueOf(value))
		}
	}
}

// GetValueByTag 获取结构体指针中的指定字段值
func GetValueByTag(obj any, tag string, tagValue string) any {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get(tag) == tagValue {
			return v.Field(i).Interface()
		}
	}
	return nil
}
