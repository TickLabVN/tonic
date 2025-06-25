package validator

import (
	"reflect"

	"github.com/TickLabVN/tonic/core/docs"
)

const VALIDATE = "validate"
const JSON = "json"
const QUERY = "query"
const PARAM = "param"

func ParseStruct(field reflect.Type) error {
	schema := docs.SchemaObject{}
	schema.Type = "object"
	return nil
}
