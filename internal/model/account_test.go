package model

import (
	"errors"
	"reflect"
	"testing"
)

var account = Account{
	ID:             1,
	DocumentNumber: "12345678900",
}

func TestAccount_IsValid(t *testing.T) {
	got := account.IsValid()
	if got != nil {
		t.Errorf("TestIsValid failed")
	}
}

func TestAccount_Is_Not_Valid(t *testing.T) {
	account.DocumentNumber = ""
	got := account.IsValid()
	if !errors.Is(got, ErrDocumentNumberInvalid) {
		t.Errorf("TestIsNotValid failed")
	}
}

func TestNewAccount(t *testing.T) {
	input := account
	expected := account
	got := input.NewAccount()
	if reflect.TypeOf(got) != reflect.TypeOf(expected) {
		t.Errorf("TestNewAccount failed")
	}
}
