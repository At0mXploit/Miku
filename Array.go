package main

import "fmt"

// ---------- ARRAY EXAMPLE ----------
func arrayExample() {
	fmt.Println("=== ARRAYS ===")

	// Declare and initialize an array
	arr := [3]int{1, 2, 3}
	fmt.Println("Original array:", arr)

	// Arrays are copied when passed to functions
	modifyArray(arr)
	fmt.Println("After modifyArray:", arr)

	// Modify array using pointer
	modifyArrayPtr(&arr)
	fmt.Println("After modifyArrayPtr:", arr)
}

func modifyArray(a [3]int) {
	a[0] = 99 // does NOT affect original
}

func modifyArrayPtr(a *[3]int) {
	a[0] = 99 // affects original
}

// ---------- SLICE BASICS ----------
func sliceBasics() {
	fmt.Println("\n=== SLICE BASICS ===")

	s := []int{1, 2, 3}
	fmt.Println("Slice:", s)
	fmt.Println("len:", len(s), "cap:", cap(s))

	// Append
	s = append(s, 4)
	fmt.Println("After append:", s)
}

// ---------- SLICING & SHARED MEMORY ----------
func sliceSharing() {
	fmt.Println("\n=== SLICE SHARING ===")

	a := []int{10, 20, 30, 40}
	b := a[1:3] // shares backing array

	b[0] = 99
	fmt.Println("Original slice a:", a)
	fmt.Println("Sub-slice b:", b)
}

// ---------- APPEND REALLOCATION ----------
func appendReallocation() {
	fmt.Println("\n=== APPEND REALLOCATION ===")

	a := make([]int, 2, 2)
	a[0], a[1] = 1, 2

	b := a[:1]

	a = append(a, 3) // triggers reallocation

	b[0] = 99
	fmt.Println("Slice a:", a)
	fmt.Println("Slice b:", b)
}

// ---------- COPY ----------
func sliceCopy() {
	fmt.Println("\n=== COPY ===")

	src := []int{1, 2, 3}
	dst := make([]int, len(src))

	copy(dst, src)
	dst[0] = 99

	fmt.Println("Source:", src)
	fmt.Println("Destination:", dst)
}

// ---------- SLICE IN FUNCTION ----------
func sliceInFunction() {
	fmt.Println("\n=== SLICE IN FUNCTION ===")

	s := []int{1, 2, 3}
	modifySlice(s)
	fmt.Println("After modifySlice:", s)

	s = addToSlice(s)
	fmt.Println("After addToSlice:", s)
}

func modifySlice(s []int) {
	s[0] = 99
}

func addToSlice(s []int) []int {
	return append(s, 4)
}

// ---------- RANGE GOTCHA ----------
func rangeGotcha() {
	fmt.Println("\n=== RANGE GOTCHA ===")

	s := []int{1, 2, 3}

	for _, v := range s {
		v = 10 // does NOT modify slice
	}
	fmt.Println("After wrong range:", s)

	for i := range s {
		s[i] = 10 // correct
	}
	fmt.Println("After correct range:", s)
}

// ---------- MULTI-DIMENSIONAL SLICE ----------
func multiDimSlice() {
	fmt.Println("\n=== MULTI-DIMENSIONAL SLICE ===")

	grid := make([][]int, 2)

	for i := range grid {
		grid[i] = make([]int, 3)
	}

	grid[0][1] = 7
	fmt.Println("Grid:", grid)
}

// ---------- MAIN ----------
func main() {
	arrayExample()
	sliceBasics()
	sliceSharing()
	appendReallocation()
	sliceCopy()
	sliceInFunction()
	rangeGotcha()
	multiDimSlice()
}

