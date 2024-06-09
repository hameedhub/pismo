package model

import (
	"reflect"
	"testing"
)

var transaction = Transaction{
	ID:              1,
	AccountId:       1,
	OperationTypeId: 1,
}

func TestTransaction_NewTransaction(t *testing.T) {
	input := transaction
	expected := transaction
	got := input.NewTransaction()
	if reflect.TypeOf(got) != reflect.TypeOf(expected) {
		t.Errorf(" Type match failed")
	}
}
