package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func GetSchemaName(t reflect.Type) string {
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	path := fmt.Sprintf("%s_%s", t.PkgPath(), t.Name())
	path = strings.ReplaceAll(path, "/", "_")
	path = strings.ReplaceAll(path, "*", "")
	return strings.ReplaceAll(path, ".", "_")
}

func GetSchemaPath(t reflect.Type) string {
	schemaName := GetSchemaName(t)
	return fmt.Sprintf("#/components/schemas/%s", schemaName)
}
