package helper

import (
	"reflect"
	"strings"
)

func ImplodeStructTag(u interface{}, tag string) string {
	var pieces []string

	ut := reflect.TypeOf(u)
	for i := 0; i < ut.NumField(); i++ {
		field := ut.Field(i)
		if v, ok := field.Tag.Lookup(tag); ok {
			if v != "" {
				pieces = append(pieces, v)
			}
		}
	}

	return strings.Join(pieces, ",")
}
