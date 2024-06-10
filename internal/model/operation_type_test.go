package model

import (
	"reflect"
	"testing"
)

var operation_type = OperationType{ID: 1, Description: "PURCHASE"}

func TestOperationType_NewOperationType(t *testing.T) {
	input := operation_type
	expected := operation_type
	got := input.NewOperationType()
	if reflect.TypeOf(got) != reflect.TypeOf(expected) {
		t.Errorf(" Type match failed")
	}
}
