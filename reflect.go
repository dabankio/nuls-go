/*
 * Copyright (c) 2019. Dabank Authors
 * All rights reserved.
 */

package nuls

import (
	"fmt"
	"reflect"
	"strconv"
)

// extractValue obtains string-formed slice representation via
// input. It only handles string, int, and their slice form, and
// panics otherwise.
func extractValue(input interface{}) (output []string) {
	v := direct(reflect.ValueOf(input))

	if v.Kind() == reflect.Slice {
		length := v.Len()
		output = make([]string, length)

		for i := 0; i < length; i++ {
			output[i] = valueToStr(v.Index(i))
		}
	} else {
		output = make([]string, 1, 1)
		output[0] = valueToStr(v)
	}

	return
}

// valueToStr convert v into proper string representation
// Only handles int and string, panic otherwise.
func valueToStr(v reflect.Value) (str string) {
	switch v.Kind() {
	case reflect.String:
		str = v.String()
	case reflect.Bool:
		if v.Bool() {
			str = "true"
		} else {
			str = "false"
		}
	case reflect.Int, reflect.Int64:
		str = strconv.FormatInt(v.Int(), 10)
	case reflect.Uint8:
		str = strconv.FormatUint(v.Uint(), 10)
	default:
		panic(fmt.Sprintf("valueToStr: %v is of unexpected kind %q", v, v.Kind()))
	}
	return
}

// direct traverses the pointer chain to fetch
// the solid value
func direct(v reflect.Value) reflect.Value {
	for ; v.Kind() == reflect.Ptr; v = v.Elem() {
		// relax
	}
	return v
}
