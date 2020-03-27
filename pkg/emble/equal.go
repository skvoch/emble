package emble

import (
	"reflect"
	"testing"
	"time"
)

// EqualWithoutTime - two structures without time fields
func EqualWithoutTime(t *testing.T, expect, actual interface{}) {
	if reflect.TypeOf(expect) != reflect.TypeOf(actual) {
		t.Fatalf("Not euqal types of input strcutures")
	}

	timeType := reflect.TypeOf(&time.Time{})

	valuesOfExpect := reflect.ValueOf(expect)
	valuesOfActual := reflect.ValueOf(actual)

	for i := 0; i < valuesOfExpect.NumField(); i++ {

		if valuesOfActual.Field(i).Type() != timeType {
			//if valuesOfActual.Field(i).Va != valuesOfExpect.Field(i) {
			//		t.Fatalf("")
			//		}
		}
	}
}
