package schema_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/TickLabVN/tonic/core"
	"github.com/TickLabVN/tonic/core/docs"
	"github.com/TickLabVN/tonic/core/utils"
	"github.com/stretchr/testify/assert"
)

// Comparision validate tags: https://github.com/go-playground/validator?tab=readme-ov-file#comparisons
func TestValidate_Comparision_Eq(t *testing.T) {
	assert := assert.New(t)
	type TestEq struct {
		Age           int      `json:"age" validate:"eq=21"`
		Name          string   `json:"name" validate:"eq=John Doe"`
		IsAdmin       bool     `json:"isAdmin" validate:"eq=true"`
		RecoverEmails []string `json:"recoverEmails" validate:"eq=1"`
		Score         float64  `json:"score" validate:"eq=3.14"`
	}

	dt := reflect.TypeOf(TestEq{})
	_, err := docs.ParseSchema(dt)
	if err != nil {
		t.Fatal(err)
	}

	spec := core.Init()
	schemaName := utils.GetSchemaPath(dt)

	schema, ok := spec.Components.Schemas[schemaName]
	assert.True(ok)
	assert.NotNil(schema)

	b, err := json.Marshal(schema)
	if err != nil {
		t.Fatal(err)
	}
	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"age": {
				"type": "integer",
				"format": "int32",
				"minimum": 21,
				"maximum": 21
			},
			"name": {
				"type": "string",
				"enum": ["John Doe"]
			},
			"isAdmin": {
				"type": "boolean",
				"enum": [true]
			},
			"score": {
				"type": "number",
				"format": "double",
				"enum": [3.14]
			},
			"recoverEmails": {
				"type": "array",
				"items": {
					"type": "string"
				},
				"minItems": 1,
				"maxItems": 1
			}
		}
	}`, string(b))
}

func TestValidate_Comparision_EqIgnoreCase(t *testing.T) {
	assert := assert.New(t)
	type TestEqIgnoreCase struct {
		Name string `json:"name" validate:"eq_ignore_case=John Doe"`
	}

	dt := reflect.TypeOf(TestEqIgnoreCase{})
	_, err := docs.ParseSchema(dt)
	if err != nil {
		t.Fatal(err)
	}

	spec := core.Init()
	schemaName := utils.GetSchemaPath(dt)

	schema, ok := spec.Components.Schemas[schemaName]
	assert.True(ok)
	assert.NotNil(schema)

	b, err := json.Marshal(schema)
	if err != nil {
		t.Fatal(err)
	}
	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"name": {
				"type": "string",
				"enum": ["John Doe"]
				"description": "Equal ignore case"
			}
		}
	}`, string(b))
}

func TestValidate_Comparision_Gt(t *testing.T) {
	assert := assert.New(t)
	type TestGt struct {
		Age            int      `json:"age" validate:"gt=21"`
		Name           string   `json:"name" validate:"gt=5"`
		RecoverEmails  []string `json:"recoverEmails" validate:"gt=1.2"`
		InterestedList []string `json:"interestedList" validate:"gt=2"`
	}
	dt := reflect.TypeOf(TestGt{})
	_, err := docs.ParseSchema(dt)
	if err != nil {
		t.Fatal(err)
	}

	schemaPath := utils.GetSchemaPath(dt)
	spec := core.Init()
	schema, ok := spec.Components.Schemas[schemaPath]

	assert.True(ok)
	assert.NotNil(schema)

	b, err := json.Marshal(schema)
	if err != nil {
		t.Fatal(err)
	}

	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"age": {
				"type": "integer",
				"format": "int32",
				"minimum": 21,
				"exclusiveMinimum": true
			},
			"name": {
				"type": "string",
				"minimum": 6,
				"exclusiveMinimum": true
			},
			"recoverEmails": {
				"type": "array",
				"items": { "type": "string" },
				"minItems": 2
			},
			"interestedList": {
				"type": "array",
				"items": { "type": "string" },
				"minItems": 3
			}
		}
	}`, string(b))
}

func TestValidate_Comparision_Gte(t *testing.T) {
	assert := assert.New(t)
	type TestGt struct {
		Age           int      `json:"age" validate:"gte=21"`
		Name          string   `json:"name" validate:"gte=5"`
		RecoverEmails []string `json:"recoverEmails" validate:"gte=1"`
	}
	dt := reflect.TypeOf(TestGt{})
	_, err := docs.ParseSchema(dt)
	if err != nil {
		t.Fatal(err)
	}

	schemaPath := utils.GetSchemaPath(dt)
	spec := core.Init()
	schema, ok := spec.Components.Schemas[schemaPath]

	assert.True(ok)
	assert.NotNil(schema)

	b, err := json.Marshal(schema)
	if err != nil {
		t.Fatal(err)
	}

	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"age": {
				"type": "integer",
				"format": "int32",
				"minimum": 21,
				"exclusiveMinimum": true
			},
			"name": {
				"type": "string",
				"minimum": 6,
				"exclusiveMinimum": true
			},
			"recoverEmails": {
				"type": "array",
				"items": { "type": "string" },
				"minItems": 1
			}
		}
	}`, string(b))
}

func TestValidate_Comparision_Lt(t *testing.T) {
	assert := assert.New(t)
	type TestLt struct {
		Age            int      `json:"age" validate:"lt=21"`
		Name           string   `json:"name" validate:"lt=5"`
		RecoverEmails  []string `json:"recoverEmails" validate:"lt=1.2"`
		InterestedList []string `json:"interestedList" validate:"lt=2"`
	}
	dt := reflect.TypeOf(TestLt{})
	_, err := docs.ParseSchema(dt)
	if err != nil {
		t.Fatal(err)
	}

	schemaPath := utils.GetSchemaPath(dt)
	spec := core.Init()
	schema, ok := spec.Components.Schemas[schemaPath]
	assert.True(ok)
	assert.NotNil(schema)

	b, err := json.Marshal(schema)
	if err != nil {
		t.Fatal(err)
	}

	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"age": {
				"type": "integer",
				"format": "int32",
				"maximum": 21,
				"exclusiveMaximum": true
			},
			"name": {
				"type": "string",
				"maximum": 6,
				"exclusiveMaximum": true
			},
			"recoverEmails": {
				"type": "array",
				"items": { "type": "string" },
				"maxItems": 1
			},
			"interestedList": {
				"type": "array",
				"items": { "type": "string" },
				"maxItems": 1
			}
		}
	}`, string(b))
}

func TestValidate_Comparision_Lte(t *testing.T) {
	assert := assert.New(t)
	type TestLt struct {
		Age            int      `json:"age" validate:"lte=21"`
		Name           string   `json:"name" validate:"lte=5"`
		RecoverEmails  []string `json:"recoverEmails" validate:"lte=1.2"`
		InterestedList []string `json:"interestedList" validate:"lte=2"`
	}
	dt := reflect.TypeOf(TestLt{})
	_, err := docs.ParseSchema(dt)
	if err != nil {
		t.Fatal(err)
	}

	schemaPath := utils.GetSchemaPath(dt)
	spec := core.Init()
	schema, ok := spec.Components.Schemas[schemaPath]
	assert.True(ok)
	assert.NotNil(schema)

	b, err := json.Marshal(schema)
	if err != nil {
		t.Fatal(err)
	}

	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"age": {
				"type": "integer",
				"format": "int32",
				"maximum": 21
			},
			"name": {
				"type": "string",
				"maximum": 6
			},
			"recoverEmails": {
				"type": "array",
				"items": { "type": "string" },
				"maxItems": 1
			},
			"interestedList": {
				"type": "array",
				"items": { "type": "string" },
				"maxItems": 2
			}
		}
	}`, string(b))
}

func TestValidate_Comparision_Ne(t *testing.T) {
	assert := assert.New(t)
	type TestEq struct {
		Age           int      `json:"age" validate:"ne=21"`
		Name          string   `json:"name" validate:"ne=John Doe"`
		IsAdmin       bool     `json:"isAdmin" validate:"ne=true"`
		RecoverEmails []string `json:"recoverEmails" validate:"ne=1"`
		Score         float64  `json:"score" validate:"ne=3.14"`
	}

	dt := reflect.TypeOf(TestEq{})
	_, err := docs.ParseSchema(dt)
	if err != nil {
		t.Fatal(err)
	}

	spec := core.Init()
	schemaName := utils.GetSchemaPath(dt)

	schema, ok := spec.Components.Schemas[schemaName]
	assert.True(ok)
	assert.NotNil(schema)

	b, err := json.Marshal(schema)
	if err != nil {
		t.Fatal(err)
	}
	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"age": {
				"type": "integer",
				"format": "int32",
				"not": { "enum": [21] }
				"description": "Not equal to 21"
			},
			"name": {
				"type": "string",
				"not": { "enum": ["John Doe"] }
				"description": "Not equal to John Doe"
			},
			"isAdmin": {
				"type": "boolean",
				"not": { "enum": [true] },
				"description": "Not be true"
			},
			"score": {
				"type": "number",
				"format": "double",
				"not": { "enum": [3.14] }
				"description": "Not equal to 3.14"
			},
			"recoverEmails": {
				"type": "array",
				"items": {
					"type": "string"
				},
				"not": {
					"minItems": 1,
					"maxItems": 1
				},
				"description": "Length not equal to 1"
			}
		}
	}`, string(b))
}

func TestValidate_Comparision_NeIgnoreCase(t *testing.T) {
	assert := assert.New(t)
	type TestEq struct {
		Name string `json:"name" validate:"ne_ignore_case=John Doe"`
	}

	dt := reflect.TypeOf(TestEq{})
	_, err := docs.ParseSchema(dt)
	if err != nil {
		t.Fatal(err)
	}

	spec := core.Init()
	schemaName := utils.GetSchemaPath(dt)

	schema, ok := spec.Components.Schemas[schemaName]
	assert.True(ok)
	assert.NotNil(schema)

	b, err := json.Marshal(schema)
	if err != nil {
		t.Fatal(err)
	}
	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"name": {
				"type": "string",
				"not": { "enum": ["John Doe"] },
				"description": "Not equal ignore case"
			}
		}
	}`, string(b))
}
