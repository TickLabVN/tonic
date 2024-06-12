package json

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func MalshalInline(v interface{}) ([]byte, error) {
	// if v is not a struct or pointer to struct, return json.Marshal error
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Struct && rv.Kind() != reflect.Ptr {
		return json.Marshal(v)
	}

	// if v is a pointer to struct, get the value of the pointer
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	// create a map to store the fields of the struct
	m := make(map[string]interface{})
	t := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		f := t.Field(i)
		v := rv.Field(i)
		jsonTags := f.Tag.Get("json")
		fmt.Println("jsonTags: ", jsonTags)
		if jsonTags == "" {
			data, err := MalshalInline(v.Interface())
			if err != nil {
				return nil, err
			}
			m[f.Name] = data
		} else {
			parts := strings.Split(jsonTags, ",")
			// check if the field is struct or pointer to struct
			if v.Kind() == reflect.Struct || v.Kind() == reflect.Ptr {
				if parts[0] == "" && parts[1] == "inline" {
					data, err := MalshalInline(v.Interface())
					if err != nil {
						return nil, err
					}
					// the data is a map, so we need to merge it with the current map
					// convert the data to a map
					var dataMap map[string]interface{}
					err = json.Unmarshal(data, &dataMap)
					if err != nil {
						return nil, err
					}
					// merge the dataMap with the current map
					for k, v := range dataMap {
						m[k] = v
					}
				} else {
					m[parts[0]] = v.Interface()
				}
			} else {
				m[parts[0]] = v.Interface()
			}

		}
	}

	// return the json.Marshal of the map
	return json.Marshal(m)
}
