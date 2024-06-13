package parser_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/TickLabVN/tonic/parser"
	"github.com/stretchr/testify/assert"
)

func TestParseUser(t *testing.T) {
	assert := assert.New(t)
	type User struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}

	schema, err := parser.ParseStruct(reflect.TypeOf(User{}))
	if err != nil {
		t.Fatal(err)
	}
	
	b, err := json.Marshal(schema)
	if err != nil {
		t.Fatal(err)
	}
	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"id": {
				"type": "integer",
				"format": "int32"
			},
			"username": {
				"type": "string"
			}
		}
	}`, string(b))
}
