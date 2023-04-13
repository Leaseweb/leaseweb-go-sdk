package options

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func Encode(opts interface{}) string {
	return extract(opts).Encode()
}

func extract(opts interface{}) url.Values {
	v := url.Values{}
	ot := reflect.TypeOf(opts)
	ov := reflect.ValueOf(opts)

	for i := 0; i < ot.NumField(); i++ {
		ovf := ov.Field(i)

		if ovf.Kind() == reflect.Struct {
			v = merge(v, extract(ovf.Interface()))
			continue
		}

		otf := ot.Field(i)
		p := otf.Tag.Get("param")
		if p == "" {
			p = otf.Name
		}

		if ovf.Kind() == reflect.Slice && !ovf.IsNil() {
			s := []string{}
			for i := 0; i < ovf.Len(); i++ {
				val := ovf.Index(i).Interface()
				s = append(s, fmt.Sprintf("%v", val))
			}
			v.Add(p, strings.Join(s, ","))
		}

		if ovf.Kind() != reflect.Pointer {
			continue
		}

		if !ovf.IsNil() {
			val := ovf.Elem().Interface()
			v.Add(p, fmt.Sprintf("%v", val))
		}
	}

	return v
}

func merge(v1, v2 url.Values) url.Values {
	for k, ls := range v2 {
		for _, s := range ls {
			v1.Add(k, s)
		}
	}
	return v1
}
