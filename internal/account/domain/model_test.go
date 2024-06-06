package domain

import (
	"errors"
	"reflect"
	"testing"
)

var account = Account{
	AccountID:      1,
	DocumentNumber: "12345678900",
}

func TestIsValid(t *testing.T) {
	got := account.IsValid()
	if got != nil {
		t.Errorf("TestIsValid failed")
	}
}

func TestIsNotValid(t *testing.T) {
	account.DocumentNumber = ""
	got := account.IsValid()
	if !errors.Is(got, ErrDocumentNumberEmpty) {
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
