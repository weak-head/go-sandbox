package reflection

import "reflect"

func walk(st interface{}, fn func(s string)) {
	val := reflect.ValueOf(st)
	field := val.Field(0)
	fn(field.String())
}
