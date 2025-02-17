package test

import (
	"testing"
)

func TestAdd(t *testing.T) {
	want := 10.0
	got, err := arithmetic.Add(5.0, 5.0)
	if err != nil {
		t.Errorf("Add() error = %v", err)
		return
	}
	if got != want {
		t.Errorf("Add() = %v, want %v", got, want)
	}
}

func TestSubtract(t *testing.T) {
	want := 5.0
	got, err := arithmetic.Subtract(10.0, 5.0)
	if err != nil {
		t.Errorf("Subtract() error = %v", err)
		return
	}
	if got != want {
		t.Errorf("Subtract() = %v, want %v", got, want)
	}
	