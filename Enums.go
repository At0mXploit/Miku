package main

import "fmt"

/*
ENUMS IN GO:

- Go does not have built-in enum type.
- Use `const` and `iota` to create a sequence of constants.
- `iota` automatically increments for each constant in a block.
- Useful for readable named constants instead of magic numbers.
*/

// Define Days of the Week using iota
const (
	Sunday = iota   // 0
	Monday          // 1
	Tuesday         // 2
	Wednesday       // 3
	Thursday        // 4
	Friday          // 5
	Saturday        // 6
)

// You can also define custom type for stronger type safety
type Day int

const (
	Sun Day = iota
	Mon
	Tue
	Wed
	Thu
	Fri
	Sat
)

func main() {
	fmt.Println("=== Basic Enum Using iota ===")
	fmt.Println("Sunday:", Sunday)
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("Saturday:", Saturday)

	fmt.Println("\n=== Enum with Custom Type ===")
	var today Day = Wed
	fmt.Println("Today is:", today)

	// You can use enums in switch statements
	fmt.Println("\n=== Using Enum in Switch ===")
	switch today {
	case Sun, Sat:
		fmt.Println("It's weekend!")
	case Mon, Tue, Wed, Thu, Fri:
		fmt.Println("It's a weekday!")
	default:
		fmt.Println("Unknown day")
	}

	// =====================
	// Custom enum example
	// =====================
	fmt.Println("\n=== Custom Status Enum Example ===")
	type Status int
	const (
		Pending Status = iota
		InProgress
		Completed
		Failed
	)

	currentStatus := InProgress
	fmt.Println("Current Status:", currentStatus)

	switch currentStatus {
	case Pending:
		fmt.Println("Pending task")
	case InProgress:
		fmt.Println("Task in progress")
	case Completed:
		fmt.Println("Task completed")
	case Failed:
		fmt.Println("Task failed")
	}
}

/*
Explanation:

1. `iota`:
   - Starts at 0 and increments by 1 in a constant block.
   - Makes creating sequences of constants easy.

2. Custom type (e.g., Day, Status):
   - Adds type safety.
   - Prevents accidentally mixing with other integers.

3. Enums are often used in:
   - Switch statements
   - Representing statuses, modes, days, etc.

Key idea:
- Go enums are **constants with optional custom type**.
- `iota` simplifies sequential constants.
*/

