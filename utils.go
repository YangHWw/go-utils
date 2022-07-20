package utils

import "reflect"

func TypeOf(e interface{}) reflect.Type {
	t := reflect.TypeOf(e)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func ValueOf(e interface{}) reflect.Value {
	v := reflect.ValueOf(e)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}
