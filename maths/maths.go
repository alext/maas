package maths

func Factorial(n int) int {
	if n < 0 {
		panic("Can't calculate factorial for negative number")
	}

	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func Fibonacci(n int) int {
	if n < 0 {
		panic("Can't calculate fibonacci for negative number")
	}

	if n == 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	return Fibonacci(n-2) + Fibonacci(n-1)
}
