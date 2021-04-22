package main

import "testing"

func TestValues(t *testing.T) {
	var result = arabicNumber("I")
	if result != 1 {
		t.Errorf("Expecting 1 got %d", result)
	}
	result = arabicNumber("II")
	if result != 2 {
		t.Errorf("Expecting 2 got %d", result)
	}

	result = arabicNumber("III")
	if result != 3 {
		t.Errorf("Expecting 3 got %d", result)
	}
	result = arabicNumber("IV")
	if result != 4 {
		t.Errorf("Expecting 4 got %d", result)
	}
	result = arabicNumber("V")
	if result != 5 {
		t.Errorf("Expecting 5 got %d", result)
	}
	result = arabicNumber("VI")
	if result != 6 {
		t.Errorf("Expecting 6 got %d", result)
	}
	result = arabicNumber("XIV")
	if result != 14 {
		t.Errorf("Expecting 14 got %d", result)
	}

	result = arabicNumber("XCIX")
	if result != 99 {
		t.Errorf("Expecting 99 got %d", result)
	}
}
