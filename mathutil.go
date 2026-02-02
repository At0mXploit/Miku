package mathutil

// Add adds two integers and returns the result
func Add(a, b int) int {
	return a + b
}

// Multiply multiplies two integers
func Multiply(a, b int) int {
	return a * b
}

// unexported function (cannot be accessed outside the package)
func subtract(a, b int) int {
	return a - b
}

