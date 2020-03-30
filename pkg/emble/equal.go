package emble

import (
	"reflect"
	"testing"
	"time"
)

// EqualWithoutTime - two structures without time fields
func EqualWithoutTime(t *testing.T, expect, actual interface{}) {
	equal(t, expect, actual)
}

func equal(t *testing.T, expect, actual interface{}) {
	if reflect.TypeOf(expect) != reflect.TypeOf(actual) {
		t.Fatalf("Not euqal types of input strcutures")
	}

	timeType := reflect.TypeOf(time.Time{})

	valuesOfExpect := reflect.ValueOf(expect).Elem()
	valuesOfActual := reflect.ValueOf(actual).Elem()

	for i := 0; i < valuesOfExpect.NumField(); i++ {

		if valuesOfActual.Field(i).Type() != timeType {
			if valuesOfActual.Field(i).Interface() != valuesOfExpect.Field(i).Interface() {
				t.Fail()

				return
			}
		} else {
			// Skip the time.Time field
		}
	}
}
