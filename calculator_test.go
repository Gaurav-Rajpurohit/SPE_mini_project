package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5 but got %f", result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	if result != 2 {
		t.Errorf("Expected 2 but got %f", result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(2, 3)
	if result != 6 {
		t.Errorf("Expected 6 but got %f", result)
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(6, 3)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != 2 {
		t.Errorf("Expected 2 but got %f", result)
	}
}
