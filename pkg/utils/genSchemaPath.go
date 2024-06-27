package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func GetSchemaPath(t reflect.Type) string {
	pkgPath := strings.ReplaceAll(t.PkgPath(), "/", "_")
	pkgPath = strings.ReplaceAll(pkgPath, ".", "_")
	return fmt.Sprintf("#/components/schemas/%s_%s", pkgPath, t.Name())
}
