package math

import "testing"

func TestAverage(t *testing.T) {
	// var v float64
	v := Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}

	v = Average([]float64{})
	if v != 0 {
		t.Error("Expected 0 for empty list, got ", v)
	}
}

func TestMin(t *testing.T) {
	// var v float64
	v := Min([]float64{1, 2, 6, 0.6})
	if v != 0.6 {
		t.Error("Expected 0.6, got ", v)
	}

	v = Min([]float64{})
    if v != 0 { // Assuming your Min function returns 0 for an empty list
        t.Error("Expected 0 for empty list, got ", v)
    }
}

func TestMax(t *testing.T) {
	// var v float64
	v := Max([]float64{1, 2, 6, 0.6})
	if v != 6 {
		t.Error("Expected 6, got ", v)
	}

	v = Max([]float64{})
    if v != 0 { // Assuming your Min function returns 0 for an empty list
        t.Error("Expected 0 for empty list, got ", v)
    }
}

func TestFib(t *testing.T) {
	// var v float64
	v := Fib(10)
	if v != 55 {
		t.Error("Expected 55, got ", v)
	}
}
