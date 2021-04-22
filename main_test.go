package main

import "testing"

func TestRomanToInt(t *testing.T) {

	input := "MCMXCIV"
	expected := 1994
	result := romanToInt(input)
	if result != expected {
		t.Errorf("Test failed : Expected %v |  Actual %v", expected, result)
	} else {
		t.Logf("Pass")
	}
}

func TestFibonacci(t *testing.T) {
	memo := map[int]int{}
	var result = fibonacci(6, memo)
	if result != 8 {
		t.Errorf("Expecting 8 got %d", result)
	}
}
