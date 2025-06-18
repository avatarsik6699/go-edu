package walker

import (
	"errors"
	"log"
	"reflect"
)

// golang challenge: write a function walk(x interface{}, fn func(string))
// which takes a struct x and calls fn for all strings fields found inside.
// difficulty level: recursively.

var (
	ErrInvalidXParam = errors.New("x should be struct")
)

func Walker(x any, fn func(string)) {
	// Должен убедиться в том, что x это структура / map.
	// После того как убедился нужно каким-то способом итеративно пройтись
	// по каждому key,value и вызвать fn с аргументом value
	// Завершить функцию

	v := reflect.ValueOf(x)

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for idx := range v.Len() {
			Walker(v.Index(idx).Interface(), fn)
		}

	}

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	log.Println(v.Kind(), v.IsValid())

	if v.Kind() == reflect.Map {
		iter := v.MapRange()
		for iter.Next() {
			mv := iter.Value()
			fn(mv.String())
		}
	}

	if v.Kind() == reflect.Struct {

		for idx := range v.NumField() {
			value := v.Field(idx)

			switch value.Kind() {
			case reflect.String:
				fn(value.String())
			case reflect.Struct:
				Walker(value.Interface(), fn)
			}

		}
	}

	if v.Kind() == reflect.Chan {
		for {
			if vch, ok := v.Recv(); ok {
				Walker(vch.Interface(), fn)
			} else {
				break
			}
		}
	}

}
