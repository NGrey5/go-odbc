package odbc

import (
	"reflect"
	"strings"
)

func TrimEndOfResults(result interface{}) {
	rv := reflect.ValueOf(result)
	trimrv(rv)
}

func trimrv(rv reflect.Value) {
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() == reflect.Struct {
		trimStruct(rv)
	}

	if rv.Kind() == reflect.Slice {
		trimSlice(rv)
	}
}

func trimSlice(rv reflect.Value) {
	ln := rv.Len()

	for i := 0; i < ln; i++ {
		trimStruct(rv.Index(i))
	}
}

func trimStruct(rv reflect.Value) {
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		if field.Kind() == reflect.String {
			field.SetString(strings.TrimSpace(field.String()))
		}
	}
}
