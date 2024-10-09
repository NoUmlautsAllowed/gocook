package form

import (
	"fmt"
	"net/url"
	"reflect"
)

func Values(v any) (url.Values, error) {
	t := reflect.Indirect(reflect.ValueOf(v)).Type()
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("only support marshal for pointer to struct")
	}

	qs := make(url.Values)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("form") == "" {
			continue
		}
		value := reflect.Indirect(reflect.ValueOf(v)).Field(i).Interface()
		s := fmt.Sprintf("%v", value)
		if len(s) > 0 {
			qs.Set(field.Tag.Get("form"), s)
		}
	}

	return qs, nil
}

func Marshal(v any) ([]byte, error) {
	values, err := Values(v)
	if err != nil {
		return nil, err
	}

	return []byte(values.Encode()), nil
}
