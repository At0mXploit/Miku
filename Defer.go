package main

import (
	"fmt"
	"os"
)

/*
DEFER IN GO:

- The deferred function executes AFTER the surrounding function finishes.
- Useful for cleanup tasks like:
    - closing files
    - unlocking mutexes
    - releasing resources
- Multiple defer calls execute in reverse order (stack behavior).
*/

func main() {

	fmt.Println("=== 1. Basic Defer Example ===")
	fmt.Println("Start of main")

	// This will be executed at the end of main
	defer fmt.Println("Deferred: End of main")

	fmt.Println("Doing some work...")
	fmt.Println("Work done")

	// Output order:
	// Start of main
	// Doing some work...
	// Work done
	// Deferred: End of main

	// =====================
	// 2. Defer with multiple statements
	// =====================
	fmt.Println("\n=== 2. Multiple Defers ===")

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	defer fmt.Println("Deferred 3")

	fmt.Println("Main function executing...")

	// Execution order of defer:
	// Deferred 3
	// Deferred 2
	// Deferred 1

	// =====================
	// 3. Defer for cleanup (closing file)
	// =====================
	fmt.Println("\n=== 3. Defer for cleanup ===")

	// Create a file
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	// Ensure the file is closed at the end of the function
	defer file.Close()

	// Write something to the file
	file.WriteString("Hello, defer!\n")
	fmt.Println("Wrote to file, defer will close it automatically")

	// =====================
	// 4. Defer with function calls and arguments
	// =====================
	fmt.Println("\n=== 4. Defer with function calls ===")

	x := 5
	defer fmt.Println("Deferred value of x:", x) // x is evaluated at the time of defer
	x = 10
	fmt.Println("x changed to:", x)

	// =====================
	// 5. Defer with recovery from panic
	// =====================
	fmt.Println("\n=== 5. Defer with panic recovery ===")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("About to panic...")
	panic("Something went wrong!") // deferred recovery function will handle this

	fmt.Println("This line will not execute due to panic")
}

