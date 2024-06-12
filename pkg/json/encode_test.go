package json

import (
	"fmt"
	"testing"
)

func TestMalshalInline(t *testing.T) {
	// check if v is a struct, if not return error

	type subStruct struct {
		SubField1 string `json:"subField1"`
		SubField2 string `json:"subField2"`
	}

	type subStruct2 struct {
		SubField3 string    `json:"subField3"`
		SubField4 subStruct `json:",inline,omitempty"`
	}

	type testStruct struct {
		Field1 string     `json:"field1"`
		Field2 subStruct  `json:",inline,omitempty"`
		Field3 string     `json:"field3"`
		Field4 subStruct2 `json:",inline,omitempty"`
	}

	operation := testStruct{
		Field1: "field1",
		Field2: subStruct{
			SubField1: "subField1",
			SubField2: "subField2",
		},
		Field3: "field3",
		Field4: subStruct2{
			SubField3: "subField3",
			SubField4: subStruct{
				SubField1: "subField4.1",
				SubField2: "subField4.2",
			},
		},
	}

	encodeData, err := MalshalInline(operation)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(encodeData))

}
