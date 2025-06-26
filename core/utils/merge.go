package utils

import (
	"reflect"
)

func MergeStructs[D any](values ...D) D {
	var result D
	resVal := reflect.ValueOf(&result).Elem()

	for _, v := range values {
		val := reflect.ValueOf(v)
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			resField := resVal.Field(i)
			if !isZero(field) {
				// Deep merge for struct fields
				if field.Kind() == reflect.Struct && resField.CanSet() {
					merged := MergeStructs(resField.Interface(), field.Interface())
					resField.Set(reflect.ValueOf(merged))
				} else if resField.CanSet() {
					resField.Set(field)
				}
			}
		}
	}

	return result
}

func isZero(v reflect.Value) bool {
	zero := reflect.Zero(v.Type())
	return reflect.DeepEqual(v.Interface(), zero.Interface())
}
