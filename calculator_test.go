package main

import (
	"math"
	"testing"
)

// ============================ Addition Testcases =============================

func TestAdd1(t *testing.T) {
	result := add(9, 10)
	if result != 19 {
		t.Errorf("Expected 19, got %g", result)
	}
}

func TestAdd2(t *testing.T) {
	result := add(-5, 4)
	if result != -1 {
		t.Errorf("Expected -1, got %g", result)
	}
}

func TestAdd3(t *testing.T) {
	result := add(-3, -4)
	if result != -7 {
		t.Errorf("Expected -7, got %g", result)
	}
}

func TestAdd4(t *testing.T) {
	result := add(math.MaxFloat64, math.MaxFloat64)
	if !math.IsInf(result, 1) {
		t.Errorf("Expected +Inf, got %g", result)
	}
}

func TestAdd5(t *testing.T) {
	result := add(-math.MaxFloat64, -math.MaxFloat64)
	if !math.IsInf(result, -1) {
		t.Errorf("Expected -Inf, got %g", result)
	}
}

func TestAdd6(t *testing.T) {
	result := add(3, 0)
	if result != 3 {
		t.Errorf("Expected 3, got %g", result)
	}
}

func TestAdd7(t *testing.T) {
	result := add(1e16, 1)
	if result != 10000000000000001 {
		t.Errorf("Expected 10000000000000001, got %g", result)
	}
}

func TestAdd8(t *testing.T) {
	result := add(math.Inf(1), math.Inf(-1))
	if !math.IsNaN(result) {
		t.Errorf("Expected NaN, got %g", result)
	}
}

func TestAdd9(t *testing.T) {
	result := add(math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64)
	if result == 0 {
		t.Errorf("Expected non-zero result, got %g", result)
	}
}

// ============================ Subtraction =============================

func TestSubtraction1(t *testing.T) {
	result := subtract(20, 1)
	if result != 19 {
		t.Errorf("Expected 19, got %g", result)
	}
}

func TestSubtraction2(t *testing.T) {
	result := subtract(-5, 4)
	if result != -9 {
		t.Errorf("Expected -9, got %g", result)
	}
}

func TestSubtraction3(t *testing.T) {
	result := subtract(-3, -4)
	if result != 1 {
		t.Errorf("Expected 1, got %g", result)
	}
}

// ============================ Multiplication =============================

func TestMultiplication1(t *testing.T) {
	result := multiply(9, 10)
	if result != 90 {
		t.Errorf("Expected 90, got %g", result)
	}
}

func TestMultiplication2(t *testing.T) {
	result := multiply(-5, 4)
	if result != -20 {
		t.Errorf("Expected -20, got %g", result)
	}
}

// ============================ Division =============================

func TestDivision1(t *testing.T) {
	result, err := divide(100, 10)
	if err != nil || result != 10 {
		t.Errorf("Expected 10, got %g (error: %v)", result, err)
	}
}

func TestDivision4(t *testing.T) {
	_, err := divide(10, 0)
	if err == nil {
		t.Errorf("Expected division by zero error")
	}
}

// ============================ Square Root =============================

func TestSquareRoot1(t *testing.T) {
	result, err := squareRoot(25)
	if err != nil || result != 5 {
		t.Errorf("Expected 5, got %g", result)
	}
}

func TestSquareRoot3(t *testing.T) {
	_, err := squareRoot(-4)
	if err == nil {
		t.Errorf("Expected error for negative input")
	}
}

// ============================ Factorial =============================

func TestFactorial1(t *testing.T) {
	result, err := fact(0)
	if err != nil || result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestFactorial3(t *testing.T) {
	result, err := fact(5)
	if err != nil || result != 120 {
		t.Errorf("Expected 120, got %d", result)
	}
}

// ============================ Natural Log =============================

func TestNaturalLog1(t *testing.T) {
	result, err := naturalLog(1)
	if err != nil || result != 0 {
		t.Errorf("Expected 0, got %g", result)
	}
}

// ============================ Power =============================

func TestPow1(t *testing.T) {
	result := pow(2, 3)
	if result != 8 {
		t.Errorf("Expected 8, got %g", result)
	}
}
