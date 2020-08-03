package main

import "testing"

func Test_plus_func_must_eq_expected_must_pass(t *testing.T) {
	expected := 5
	actual := plus(2, 3)
	if actual != expected {
		t.Errorf("expected  %d but actual  %d", expected, actual)
	}
}


func Test_plus_func_must_not_eq_expected_must_pass(t *testing.T) {
	expected := 5
	actual := plus(2, 2)
	if actual == expected {
		t.Errorf("expected not equal to %d but actual %d", expected, actual)
	}
}

func plus(a, b int) int { return a + b }
