package guerystring

import (
	"fmt"
	"reflect"
	"strings"
)

func Encode(x interface{}) (string, error) {
	var sb strings.Builder
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	numField := t.NumField()

	for i := 0; i < numField; i++ {
		name, ok := t.Field(i).Tag.Lookup("query")
		if !ok {
			continue
		}
		value := v.Field(i).Interface()

		sb.WriteString(fmt.Sprintf("%s=%v", name, value))
		if i < numField-1 {
			sb.WriteRune('&')
		}
	}

	return sb.String(), nil
}
