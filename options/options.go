package options

import (
	"fmt"
	"net/url"
	"reflect"
)

func Encode(opts interface{}) string {
	v := url.Values{}
	ot := reflect.TypeOf(opts)
	ov := reflect.ValueOf(opts)

	for i := 0; i < ot.NumField(); i++ {
		otf := ot.Field(i)
		ovf := ov.Field(i)

		if !ovf.IsNil() {
			p := otf.Tag.Get("param")
			if p == "" {
				p = otf.Name
			}
			val := ovf.Elem().Interface()
			v.Add(p, fmt.Sprintf("%v", val))
		}
	}

	return v.Encode()
}
