package main

import "fmt"

// =====================
// 1. Define an interface
// =====================

// Shape is an interface with a single method Area.
// Any type that has a method `Area() float64` automatically satisfies this interface.
type Shape interface {
	Area() float64
}

// =====================
// 2. Structs that implement the interface
// =====================

// Rectangle has Width and Height fields
type Rectangle struct {
	Width, Height float64
}

// Rectangle implements the Area method, satisfying the Shape interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle has a Radius field
type Circle struct {
	Radius float64
}

// Circle implements the Area method, satisfying the Shape interface
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// =====================
// 3. Function that accepts any Shape
// =====================

// printArea accepts a Shape interface
// This function can take any type that implements Shape
func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

// =====================
// 4. Main function
// =====================
func main() {
	// Create a Rectangle instance
	r := Rectangle{Width: 5, Height: 4}

	// Create a Circle instance
	c := Circle{Radius: 3}

	// Both r and c implement Shape, so we can pass them to printArea
	fmt.Println("Rectangle:")
	printArea(r) // Rectangle's Area method will be called

	fmt.Println("Circle:")
	printArea(c) // Circle's Area method will be called

	// =====================
	// 5. Interface values explanation
	// =====================

	var s Shape
	// At this moment, s has:
	// - type: nil
	// - value: nil

	// Assign Rectangle to s
	s = r
	fmt.Println("s now holds Rectangle:", s.Area())

	// Assign Circle to s
	s = c
	fmt.Println("s now holds Circle:", s.Area())

	// =====================
	// 6. Interfaces and slices
	// =====================

	shapes := []Shape{r, c} // slice of interfaces
	for i, shape := range shapes {
		fmt.Printf("Shape %d area: %v\n", i+1, shape.Area())
	}

	// =====================
	// 7. Nil interface vs non-nil
	// =====================

	var s2 Shape // nil interface
	fmt.Println("Nil interface:", s2) // prints <nil>
}

