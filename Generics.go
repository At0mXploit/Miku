package main

import "fmt"

/*
GENERICS IN GO:

- Generics allow you to write functions or structs that can work with **any type**.
- Type parameters are specified in square brackets [].
- "any" means the type can be anything (int, string, float, etc.)
- Useful to avoid writing the same function multiple times for different types.
*/

// ----------------------
// 1. Generic function
// ----------------------

// PrintSlice prints all elements of a slice of any type
func PrintSlice[T any](s []T) {
	for i, v := range s {
		fmt.Printf("Index %d: %v\n", i, v)
	}
}

// ----------------------
// 2. Generic swap function
// ----------------------

// Swap takes two values of any type and returns them in swapped order
func Swap[T any](a, b T) (T, T) {
	return b, a
}

// ----------------------
// 3. Generic struct
// ----------------------

// Pair holds two values of the same type
type Pair[T any] struct {
	First  T
	Second T
}

// Print prints the values of the Pair
func (p Pair[T]) Print() {
	fmt.Println("First:", p.First, "Second:", p.Second)
}

// ----------------------
// 4. Main function demonstrating all
// ----------------------
func main() {
	fmt.Println("=== 1. Generic Slice Example ===")
	intSlice := []int{1, 2, 3, 4}
	stringSlice := []string{"apple", "banana", "cherry"}

	// Can use the same PrintSlice function for any type
	PrintSlice(intSlice)
	PrintSlice(stringSlice)

	fmt.Println("\n=== 2. Generic Swap Example ===")
	a, b := 10, 20
	a, b = Swap(a, b) // swap two integers
	fmt.Println("Swapped integers:", a, b)

	x, y := "hello", "world"
	x, y = Swap(x, y) // swap two strings
	fmt.Println("Swapped strings:", x, y)

	fmt.Println("\n=== 3. Generic Struct Example ===")
	intPair := Pair[int]{First: 5, Second: 10}
	stringPair := Pair[string]{First: "foo", Second: "bar"}

	intPair.Print()
	stringPair.Print()
}

/*
Explanation:

1. PrintSlice[T any]:
   - T is a type parameter.
   - any means T can be any type (int, string, float, etc.)
   - This lets us use the same function for slices of different types.

2. Swap[T any]:
   - Generic function to swap values of any type.
   - Avoids writing multiple swap functions for int, string, etc.

3. Pair[T any]:
   - Generic struct holding two values of the same type.
   - Method Print prints the values.
   - Works for int, string, or any type we choose.

Key idea:
- Generics reduce code duplication.
- Type safety is maintained (cannot mix incompatible types).
- Works for functions, structs, and methods.
*/

