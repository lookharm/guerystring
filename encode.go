package guerystring

import (
	"fmt"
	"reflect"
	"strings"
)

var errNotStruct = fmt.Errorf("guerystring: x is not struct")

func Encode(x interface{}) (string, error) {
	v := reflect.ValueOf(x)

	if v.Kind() != reflect.Struct {
		return "", errNotStruct
	}

	var (
		sb       strings.Builder
		numField = v.NumField()
	)

	for i := 0; i < numField; i++ {
		name, ok := v.Type().Field(i).Tag.Lookup("query")
		if !ok {
			continue
		}

		value, ok := v.Field(i).Interface().(string)
		if !ok {
			continue
		}

		sb.WriteString(name)
		sb.WriteByte('=')
		sb.WriteString(value)
		if i < numField-1 {
			sb.WriteByte('&')
		}
	}

	return sb.String(), nil
}
