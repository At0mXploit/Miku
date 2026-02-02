package main

import (
	"fmt"
	"mathutil" // import our custom package
)

func main() {
	// Using functions from the mathutil package
	sum := mathutil.Add(10, 5)
	product := mathutil.Multiply(4, 3)

	fmt.Println("Sum:", sum)
	fmt.Println("Product:", product)

	// Trying to access subtract will give error:
	// diff := mathutil.subtract(10, 5) // ❌ cannot use subtract (unexported)

	// Using slices and maps along with packages
	numbers := []int{1, 2, 3, 4, 5}
	total := 0
	for _, n := range numbers {
		total = mathutil.Add(total, n)
	}
	fmt.Println("Sum of slice:", total)

	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	for name, age := range ages {
		fmt.Println(name, "is", age, "years old")
	}
}

