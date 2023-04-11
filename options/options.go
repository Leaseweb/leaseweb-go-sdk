package options

import (
	"fmt"
	"net/url"
	"reflect"
)

func Encode(opts interface{}) string {
	values := extract(opts)
	return values.Encode()
}

// extract recursively extracts URL values from an interface{} and returns them as url.Values
func extract(opts interface{}) url.Values {
	v := url.Values{}
	ot := reflect.TypeOf(opts)
	ov := reflect.ValueOf(opts)
	extractValuesFromStruct(ot, ov, &v)
	return v
}

// extractValuesFromStruct recursively extracts URL values from a struct type and value and stores them in url.Values
func extractValuesFromStruct(ot reflect.Type, ov reflect.Value, v *url.Values) {
	for i := 0; i < ot.NumField(); i++ {
		otf := ot.Field(i)
		ovf := ov.Field(i)

		if otf.Type.Kind() == reflect.Struct {
			extractValuesFromStruct(otf.Type, ovf, v)
		} else if otf.Type.Kind() == reflect.Slice {
			// Handle slices
			for j := 0; j < ovf.Len(); j++ {
				param := otf.Tag.Get("param")
				if param == "" {
					param = otf.Name
				}
				val := ovf.Index(j).Interface()
				v.Add(param, fmt.Sprintf("%v", val))
			}
		} else if !isNilOrInvalid(ovf) {
			param := otf.Tag.Get("param")
			if param == "" {
				param = otf.Name
			}
			val := ovf.Elem().Interface()
			v.Add(param, fmt.Sprintf("%v", val))
		}
	}
}

func isNilOrInvalid(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		return v.IsNil()
	case reflect.Invalid:
		return true
	default:
		return false
	}
}
