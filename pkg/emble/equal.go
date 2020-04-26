package emble

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

// For test
func PrintTestMessage() {
	fmt.Println("Test from Goland")
}

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

		expectValue := valuesOfActual.Field(i).Interface()
		actualValue := valuesOfExpect.Field(i).Interface()

		/*
			if valuesOfActual.Field(i).Kind() == reflect.Struct {
				equal(t, expectValue, actualValue)

				continue
			} */

		if valuesOfActual.Field(i).Type() != timeType {

			if expectValue != actualValue {

				val := reflect.Indirect(reflect.ValueOf(expect))
				varibleName := val.Type().Field(i).Name

				t.Fatalf("Not euqal fields, Name: %s, Expect: %#v, Actual: %#v", varibleName, expectValue, actualValue)

				return
			}
		} else {
			// Skip the time.Time field
		}
	}
}
