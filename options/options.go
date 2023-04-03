package options

import (
	"fmt"
	"net/url"
	"reflect"
)

func Encode(input interface{}) string {
	queryValues := url.Values{}
	structType := reflect.TypeOf(input)
	structValue := reflect.ValueOf(input)

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		value := structValue.Field(i)

		if !value.IsNil() {
			queryParam := field.Tag.Get("param")
			if queryParam == "" {
				queryParam = field.Name
			}
			val := value.Elem().Interface()
			queryValues.Add(queryParam, fmt.Sprintf("%v", val))
		}
	}

	return queryValues.Encode()
}
