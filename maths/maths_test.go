package maths

import (
	"testing"
)

type Check struct {
	Input    int
	Expected int
}

var factorialChecks = []Check{
	{0, 1},
	{1, 1},
	{2, 2},
	{3, 6},
	{4, 24},
	{10, 3628800},
	{15, 1307674368000},
}

func TestFactorial(t *testing.T) {
	for _, check := range factorialChecks {
		actual := Factorial(check.Input)
		if actual != check.Expected {
			t.Errorf("Factorial(%d) = %d, want %d", check.Input, actual, check.Expected)
		}
	}
}

var fibonacciChecks = []Check{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{10, 55},
}

func TestFibonacci(t *testing.T) {
	for _, check := range fibonacciChecks {
		actual := Fibonacci(check.Input)
		if actual != check.Expected {
			t.Errorf("Fibonacci(%d) = %d, want %d", check.Input, actual, check.Expected)
		}
	}
}
