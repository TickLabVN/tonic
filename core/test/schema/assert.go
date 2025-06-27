package schema_test

import (
	"encoding/json"
	"reflect"

	"github.com/TickLabVN/tonic/core/docs"
	"github.com/TickLabVN/tonic/core/utils"
	"github.com/stretchr/testify/assert"
)

func AssertParse(assert *assert.Assertions, spec *docs.OpenApi, data any) (string, error) {
	dt := reflect.TypeOf(data)
	spec.Components.AddSchema(dt, "validate")
	schemaName := utils.GetSchemaName(dt)

	schema, ok := spec.Components.Schemas[schemaName]
	assert.True(ok)
	assert.NotNil(schema)

	b, err := json.Marshal(schema)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
