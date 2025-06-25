package utils

import (
	"encoding/json"
	"reflect"

	tonic "github.com/TickLabVN/tonic/core"
	"github.com/TickLabVN/tonic/core/parser"
	"github.com/stretchr/testify/assert"
)

func AssertParse(assert *assert.Assertions, data interface{}) (string, error) {
	dt := reflect.TypeOf(data)
	err := parser.ParseStruct(dt)
	if err != nil {
		return "", err
	}

	spec := tonic.GetSpec()
	schemaName := GetSchemaPath(dt)

	schema, ok := spec.Components.Schemas[schemaName]
	assert.True(ok)
	assert.NotNil(schema)

	b, err := json.Marshal(schema)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
