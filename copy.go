package utilize

import "reflect"

// DeepCopy deep copy
func DeepCopy(src any) (any, error) {
	srcValue := reflect.ValueOf(src)

	// If src is not a pointer or interface, simply return the original value
	if srcValue.Kind() != reflect.Ptr && srcValue.Kind() != reflect.Interface {
		return src, nil
	}

	// If src is a nil pointer or interface, return a nil copy
	if srcValue.IsNil() {
		return nil, nil
	}

	// Insert a new value of the same type as src
	dest := reflect.New(srcValue.Elem().Type()).Interface()

	// If src is a slice, perform a deep copy of the slice elements
	if srcValue.Elem().Kind() == reflect.Slice {
		srcSlice := srcValue.Elem()
		destSlice := reflect.ValueOf(dest).Elem()

		for i := 0; i < srcSlice.Len(); i++ {
			elem, err := DeepCopy(srcSlice.Index(i).Interface())
			if err != nil {
				return nil, err
			}
			destSlice = reflect.Append(destSlice, reflect.ValueOf(elem))
		}
		return dest, nil
	}

	// If src is a struct, perform a deep copy of the struct fields
	if srcValue.Elem().Kind() == reflect.Struct {
		for i := 0; i < srcValue.Elem().NumField(); i++ {
			field := srcValue.Elem().Field(i)
			// Deep copy each struct field
			deepCopyField := reflect.New(field.Type()).Interface()
			DeepCopy(field.Interface())
			reflect.ValueOf(dest).Elem().Field(i).Set(reflect.ValueOf(deepCopyField).Elem())
		}
		return dest, nil
	}

	// For other types, return the original value
	return src, nil
}
