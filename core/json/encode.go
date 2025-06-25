package json

import (
	"encoding/json"
	"reflect"
	"strings"
)

func MarshalInline(v interface{}) ([]byte, error) {
	// if v is not a struct or pointer to struct, return json.Marshal error
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return json.Marshal(v)
		}
		rv = rv.Elem()
	}

	kind := rv.Kind()
	if isPrimitiveType(kind) {
		return json.Marshal(v)
	}

	if kind == reflect.Array || kind == reflect.Slice {
		// check if the value is nil
		if rv.IsNil() {
			return json.Marshal(v)
		}

		if rv.Len() == 0 {
			return json.Marshal(v)
		}

		// check if the array is stored the primitive type
		if isPrimitiveType(rv.Index(0).Kind()) {
			return json.Marshal(v)
		}

		// check if the array is stored the array of any type
		if rv.Index(0).Kind() == reflect.Array {
			array := make([]interface{}, rv.Len())
			for i := 0; i < rv.Len(); i++ {
				data, err := MarshalInline(rv.Index(i).Interface())
				if err != nil {
					return nil, err
				}
				array[i] = data
			}

			return json.Marshal(array)
		}

		// check if the array is stored the pointer to struct
		if rv.Index(0).Kind() == reflect.Ptr {
			// check if the value is nil
			if rv.Index(0).IsNil() {
				return json.Marshal(v)
			}

			if rv.Index(0).Elem().Kind() == reflect.Struct {
				// create a slice to store the fields of the struct
				s := make([]interface{}, rv.Len())
				for i := 0; i < rv.Len(); i++ {
					data, err := MarshalInline(rv.Index(i).Interface())
					if err != nil {
						return nil, err
					}

					decoded := map[string]interface{}{}
					err = json.Unmarshal(data, &decoded)
					if err != nil {
						return nil, err
					}
					s[i] = decoded
				}
				return json.Marshal(s)
			}
		}

		if rv.Index(0).Kind() == reflect.Struct {
			// create a slice to store the fields of the struct
			s := make([]interface{}, rv.Len())
			for i := 0; i < rv.Len(); i++ {
				data, err := MarshalInline(rv.Index(i).Interface())
				if err != nil {
					return nil, err
				}

				decoded := map[string]interface{}{}
				err = json.Unmarshal(data, &decoded)
				if err != nil {
					return nil, err
				}
				s[i] = decoded
			}
			return json.Marshal(s)
		}
	}

	if kind == reflect.Map {
		// check if the value is nil
		if rv.IsNil() {
			return json.Marshal(v)
		}

		if rv.Len() == 0 {
			return json.Marshal(v)
		}

		// check if the map is stored the primitive type
		if isPrimitiveType(rv.MapIndex(rv.MapKeys()[0]).Kind()) {
			return json.Marshal(v)
		}

		// create a map to store the fields of the struct
		m := make(map[string]interface{})
		for _, key := range rv.MapKeys() {
			value := rv.MapIndex(key)
			data, err := MarshalInline(value.Interface())
			if err != nil {
				return nil, err
			}

			decoded := map[string]interface{}{}
			err = json.Unmarshal(data, &decoded)
			if err != nil {
				return nil, err
			}
			m[key.String()] = decoded
		}
		return json.Marshal(m)
	}

	// create a map to store the fields of the struct
	m := make(map[string]interface{})
	t := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		f := t.Field(i)
		v := rv.Field(i)
		jsonTags := f.Tag.Get("json")
		if jsonTags == "" {
			data, err := MarshalInline(v.Interface())
			if err != nil {
				return nil, err
			}
			m[f.Name] = data
		} else {
			/**
			if the field has json tag, get the json tag value
			**/
			parts := strings.Split(jsonTags, ",")
			inlineFlag := false
			for i := 1; i < len(parts); i++ {
				if parts[i] == "inline" {
					inlineFlag = true
					break
				}
			}

			kindOfV := v.Kind()

			if v.IsZero() {
				continue
			} else if (kindOfV == reflect.Struct || kindOfV == reflect.Ptr) && parts[0] == "" && inlineFlag {
				data, err := MarshalInline(v.Interface())
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
			} else if isPrimitiveType(kindOfV) {
				m[parts[0]] = v.Interface()
			} else {
				data, err := MarshalInline(v.Interface())
				if err != nil {
					return nil, err
				}
				var decoded interface{}
				err = json.Unmarshal(data, &decoded)
				if err != nil {
					return nil, err
				}
				m[parts[0]] = decoded
			}
		}
	}
	// return the json.Marshal of the map
	return json.Marshal(m)
}

func isPrimitiveType(kind reflect.Kind) bool {
	return kind == reflect.Bool || kind == reflect.Int || kind == reflect.Uint || kind == reflect.Int8 || kind == reflect.Uint8 || kind == reflect.Int16 || kind == reflect.Uint16 || kind == reflect.Int32 || kind == reflect.Uint32 || kind == reflect.Int64 || kind == reflect.Uint64 || kind == reflect.Complex64 || kind == reflect.Complex128 || kind == reflect.String

}
