package options

import (
	"fmt"
	"net/url"
	"reflect"
)

func Encode(opts interface{}) string {
	v := url.Values{}
	encodeStruct(reflect.TypeOf(opts), reflect.ValueOf(opts), &v)
	return v.Encode()
}

func encodeStruct(ot reflect.Type, ov reflect.Value, v *url.Values) {
	for i := 0; i < ot.NumField(); i++ {
		otf := ot.Field(i)
		ovf := ov.Field(i)

		if otf.Type.Kind() == reflect.Struct {
			encodeStruct(otf.Type, ovf, v)
		} else if !ovf.IsNil() {
			p := otf.Tag.Get("param")
			if p == "" {
				p = otf.Name
			}
			val := ovf.Elem().Interface()
			v.Add(p, fmt.Sprintf("%v", val))
		}
	}
}
