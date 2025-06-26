package utils

import (
	"reflect"
)

// a := TestStruct{Field1: "Hello", Field2: 42}
//
//	b := TestStruct{Field1: "World", Field2: 0}
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
	if v.Kind() >= reflect.Int && v.Kind() <= reflect.Int64 && v.Int() == 0 ||
		v.Kind() >= reflect.Uint && v.Kind() <= reflect.Uintptr && v.Uint() == 0 ||
		v.Kind() == reflect.Float32 || v.Kind() == reflect.Float64 && v.Float() == 0 ||
		v.Kind() == reflect.Bool && !v.Bool() {
		return false
	}
	zero := reflect.Zero(v.Type())
	return reflect.DeepEqual(v.Interface(), zero.Interface())
}
