package schema_test

import (
	"encoding/json"
	"reflect"

	"github.com/TickLabVN/tonic/core/docs"
	"github.com/stretchr/testify/assert"
)

func AssertParse(assert *assert.Assertions, spec *docs.OpenApi, data any) (string, error) {
	dt := reflect.TypeOf(data)
	schema, err := spec.Components.AddSchema(dt, "json", "validate")
	assert.Nil(err, "AddSchema should not return an error")
	assert.NotNil(schema)

	b, err := json.Marshal(schema)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
