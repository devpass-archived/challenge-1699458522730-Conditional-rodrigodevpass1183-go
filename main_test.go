package main

import (
	"testing"
)

func TestMax(t *testing.T) {
	num1 := 10
	num2 := 5

	result := Max(num1, num2)

	if result != num1 {
		t.Errorf("Expected %d to be greater than %d, but got %d", num1, num2, result)
	}

	num1 = 3
	num2 = 8

	result = Max(num1, num2)

	if result != num2 {
		t.Errorf("Expected %d to be greater than %d, but got %d", num2, num1, result)
	}
}