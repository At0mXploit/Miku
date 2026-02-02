package main

import "fmt"

// =====================
// Struct example
// =====================
type Person struct {
	Name string
	Age  int
}

// Function that doubles a number using pointer
func double(n *int) {
	*n = *n * 2
}

// Function that modifies a struct using pointer
func celebrateBirthday(p *Person) {
	p.Age += 1
}

func main() {
	// =====================
	// 1. Basic pointer
	// =====================
	fmt.Println("=== Basic pointer ===")
	x := 10          // normal variable
	p := &x          // pointer to x
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", p)
	fmt.Println("Dereferencing pointer:", *p)

	*p = 20          // modify value via pointer
	fmt.Println("x after modification via pointer:", x)

	// =====================
	// 2. Pointer in functions
	// =====================
	fmt.Println("\n=== Pointer in functions ===")
	num := 5
	fmt.Println("Before double:", num)
	double(&num)      // pass address of num
	fmt.Println("After double:", num)

	// =====================
	// 3. Pointer with structs
	// =====================
	fmt.Println("\n=== Pointer with structs ===")
	alice := Person{Name: "Alice", Age: 30}
	fmt.Println("Before birthday:", alice)

	celebrateBirthday(&alice) // pass pointer to struct
	fmt.Println("After birthday:", alice)

	// =====================
	// 4. Nil pointers
	// =====================
	fmt.Println("\n=== Nil pointers ===")
	var p2 *int   // nil pointer
	if p2 == nil {
		fmt.Println("p2 is nil")
	}
	// Uncommenting the next line will cause a panic
	// fmt.Println(*p2)

	// =====================
	// 5. Pointer to pointer
	// =====================
	fmt.Println("\n=== Pointer to pointer ===")
	a := 100
	ptr1 := &a          // pointer to a
	ptr2 := &ptr1       // pointer to pointer
	fmt.Println("Value of a:", a)
	fmt.Println("Value via ptr1:", *ptr1)
	fmt.Println("Value via ptr2:", **ptr2)

	// Modify a using pointer to pointer
	**ptr2 = 500
	fmt.Println("Value of a after modification via ptr2:", a)
}

