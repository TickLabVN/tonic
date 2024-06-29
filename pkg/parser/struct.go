package parser

import (
	"reflect"

	"github.com/TickLabVN/tonic/docs"
)

const VALIDATE = "validate"
const JSON = "json"
const QUERY = "query"
const PARAM = "param"

func ParseStruct(field reflect.Type) error {
	schema := docs.Schema{}
	schema.Type = "object"
	return nil
}
