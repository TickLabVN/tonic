package schema_test

import (
	"testing"
	"time"

	"github.com/TickLabVN/tonic/core"
	"github.com/stretchr/testify/assert"
)

func TestStructWithPrimitiveField(t *testing.T) {
	assert := assert.New(t)
	type User struct {
		ID        int     `json:"id"`
		Age       int32   `json:"age"`
		CreatedAt int64   `json:"created_at"`
		Username  string  `json:"username"`
		Email     string  `json:"email"`
		Balance   float64 `json:"balance"`
		Score     float32 `json:"score"`
		Disabled  bool    `json:"disabled"`
		Remarks   string  // Default json tag is field name
	}

	spec := core.Init()

	schema, err := AssertParse(assert, spec, User{})
	assert.Nil(err)

	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"id": {
				"type": "integer",
				"format": "int32"
			},
			"age": { type: "integer", "format": "int32" },
			"created_at": { type: "integer", "format": "int64" },
			"username": { type: "string" },
			"email": { type: "string", "format": "email" },
			"balance": { type: "number", "format": "double" },	
			"score": { type: "number", "format": "float" },
			"disabled": { type: "boolean" },
			"Remarks": { type: "string" }
		}
	}`, schema)
}

func TestStructWithCompoundField(t *testing.T) {
	assert := assert.New(t)

	type Information struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Age       int    `json:"age"`
	}
	type User struct {
		Addresses []string    `json:"addresses"`
		Info      Information `json:"info"`
	}

	spec := core.Init()
	schema, err := AssertParse(assert, spec, User{})
	assert.Nil(err)
	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"addresses": {
				"type": "array",
				"items": { "type": "string" }
			},
			"info": {
				"type": "object",
				"properties": {
					"firstName": { "type": "string" },
					"lastName": { "type": "string" },
					"age": { "type": "integer" }
				}
			}
		}
	}`, schema)
}

func TestStructWithStructPointerField(t *testing.T) {
	assert := assert.New(t)

	type Information struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Age       int    `json:"age"`
	}
	type User struct {
		Info *Information `json:"info"`
	}

	spec := core.Init()
	schema, err := AssertParse(assert, spec, User{})
	assert.Nil(err)

	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"info": {
				"type": "object",
				"properties": {
					"firstName": { "type": "string" },
					"lastName": { "type": "string" },
					"age": { "type": "integer" }
				}
			}
		}
	}`, schema)
}
func TestStructWithFieldIsArrayOfStruct(t *testing.T) {
	assert := assert.New(t)

	type Information struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Age       int    `json:"age"`
	}
	type User struct {
		Infos []Information `json:"infos"`
	}

	spec := core.Init()
	schema, err := AssertParse(assert, spec, User{})
	assert.Nil(err)
	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"infos": {
				"type": "array",
				"items": {
					"type": "object",
					"properties": {
						"firstName": { "type": "string" },
						"lastName": { "type": "string" },
						"age": { "type": "integer" }
					}
				}
			}
		}
	}`, schema)
}
func TestStructWithPrimitivePointerType(t *testing.T) {
	assert := assert.New(t)

	type User struct {
		ID        *int     `json:"id"`
		Age       *int32   `json:"age"`
		CreatedAt *int64   `json:"created_at"`
		Username  *string  `json:"username"`
		Email     *string  `json:"email"`
		Balance   *float64 `json:"balance"`
		Score     *float32 `json:"score"`
		Disabled  *bool    `json:"disabled"`
		Remarks   *string  `json:"remarks"`
	}

	spec := core.Init()
	schema, err := AssertParse(assert, spec, User{})
	assert.Nil(err)
	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"id": {
				"type": "integer",
				"format": "int32"
			},
			"age": { type: "integer", "format": "int32" },
			"created_at": { type: "integer", "format": "int64" },
			"username": { type: "string" },
			"email": { type: "string", "format": "email" },
			"balance": { type: "number", "format": "double" },	
			"score": { type: "number", "format": "float" },
			"disabled": { type: "boolean" },
			"remarks": { type: "string" }
		}
	}`, schema)
}
func TestStructWithTimeField(t *testing.T) {
	assert := assert.New(t)

	type User struct {
		CreatedAt time.Time `json:"created_at"`
	}

	spec := core.Init()
	schema, err := AssertParse(assert, spec, User{})
	assert.Nil(err)
	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"created_at": {
				"type": "string",
				"format": "date"
			}
		}
	}`, schema)
}
func TestEmbeddedField(t *testing.T) {
	assert := assert.New(t)

	type Information struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Age       int    `json:"age"`
	}
	type User struct {
		Information `json:",inline"`
		Username    string `json:"username"`
	}

	spec := core.Init()
	schema, err := AssertParse(assert, spec, User{})
	assert.Nil(err)
	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"firstName": { "type": "string" },
			"lastName": { "type": "string" },
			"age": { "type": "integer" },
			"username": { "type": "string" }
		}
	}`, schema)
}
func TestMultiEmbeddedField(t *testing.T) {
	assert := assert.New(t)

	type Information struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Age       int    `json:"age"`
	}
	type Address struct {
		Street string `json:"street"`
		City   string `json:"city"`
	}
	type User struct {
		Information `json:",inline"`
		Address     `json:",inline"`
		Username    string `json:"username"`
	}

	spec := core.Init()
	schema, err := AssertParse(assert, spec, User{})
	assert.Nil(err)
	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"firstName": { "type": "string" },
			"lastName": { "type": "string" },
			"age": { "type": "integer" },
			"street": { "type": "string" },
			"city": { "type": "string" },
			"username": { "type": "string" }
		}
	}`, schema)
}
func TestEmbeddedFieldWithOverlappedProperties(t *testing.T) {
	assert := assert.New(t)

	type Information struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Age       int    `json:"age"`
	}
	type User struct {
		Information `json:",inline"`
		FirstName   string `json:"firstName"`
		Username    string `json:"username"`
	}

	spec := core.Init()
	schema, err := AssertParse(assert, spec, User{})
	assert.Nil(err)
	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"firstName": { "type": "string" },
			"lastName": { "type": "string" },
			"age": { "type": "integer" },
			"username": { "type": "string" }
		}
	}`, schema)
}
func TestStructWithMapField(t *testing.T) {
	assert := assert.New(t)

	type User struct {
		Properties map[string]string `json:"properties"`
	}

	spec := core.Init()
	schema, err := AssertParse(assert, spec, User{})
	assert.Nil(err)
	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"properties": {
				"type": "object",
				"additionalProperties": {
					"type": "string"
				}
			}
		}
	}`, schema)
}
