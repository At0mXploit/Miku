package main

import "fmt"

// Function with parameters and return
func add(a, b int) int {
	return a + b
}

// Function returning multiple values
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	sum := add(10, 20)
	fmt.Println("Sum:", sum)

	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

