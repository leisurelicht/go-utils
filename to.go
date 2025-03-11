package utilize

import (
	"reflect"
	"strconv"
)

// IntS2StrS convert int slice to string slice
func IntS2StrS(s []int) []string {
	result := make([]string, len(s))
	for i, v := range s {
		result[i] = strconv.Itoa(v)
	}
	return result
}

// Int64S2StrS convert int64 slice to string slice
func Int64S2StrS(s []int64) []string {
	result := make([]string, len(s))
	for i, v := range s {
		result[i] = strconv.FormatInt(v, 10)
	}
	return result
}

// StrS2Map convert string slice to map
func StrS2Map(s []string) (res map[string]struct{}) {
	res = make(map[string]struct{})
	for _, e := range s {
		res[e] = struct{}{}
	}
	return res
}

// Struct2Map convert struct to map, tag is the tag name of struct
func Struct2Map(obj any, tag string) map[string]any {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	data := make(map[string]any, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get(tag) == "" {
			continue
		}
		data[t.Field(i).Tag.Get(tag)] = v.Field(i).Interface()
	}
	return data
}

// MapToStruct map to struct
func MapToStruct(mapVal map[string]any, val any, tag string) (ok bool) {
	t := reflect.TypeOf(val)
	v := reflect.ValueOf(val)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		if value, ok := mapVal[t.Field(i).Tag.Get(tag)]; ok {
			v.Field(i).Set(reflect.ValueOf(value))
		}
	}
	return true
}
