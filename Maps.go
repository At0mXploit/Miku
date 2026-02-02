package main

import "fmt"

/*
MENTAL MAP BEFORE YOU READ:
- Array  → fixed size, copied
- Slice  → flexible, backed by array
- len    → usable elements
- cap    → total memory available
- Map    → key-value store (unordered)
*/

func main() {

	// =====================
	// ARRAYS
	// =====================
	fmt.Println("=== ARRAYS ===")

	arr := [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	changeArray(arr)
	fmt.Println("After function call:", arr) // unchanged (copied)

	// =====================
	// SLICES BASICS
	// =====================
	fmt.Println("\n=== SLICES ===")

	s := []int{10, 20, 30}
	fmt.Println("Slice:", s)
	fmt.Println("len:", len(s), "cap:", cap(s))

	// append increases LENGTH
	s = append(s, 40)
	fmt.Println("After append:", s)
	fmt.Println("len:", len(s), "cap:", cap(s))

	// =====================
	// LENGTH vs CAPACITY
	// =====================
	fmt.Println("\n=== LEN vs CAP ===")

	a := make([]int, 2, 5) // len=2, cap=5
	a[0], a[1] = 1, 2

	fmt.Println("Slice a:", a)
	fmt.Println("len:", len(a), "cap:", cap(a))

	// You cannot access a[2] yet ❌
	a = append(a, 3) // now length increases
	fmt.Println("After append:", a)
	fmt.Println("len:", len(a), "cap:", cap(a))

	// =====================
	// SHARED BACKING ARRAY
	// =====================
	fmt.Println("\n=== SHARED BACKING ARRAY ===")

	base := []int{100, 200, 300, 400}
	sub := base[1:3] // shares memory

	sub[0] = 999
	fmt.Println("Base slice:", base)
	fmt.Println("Sub slice:", sub)

	// =====================
	// APPEND REALLOCATION
	// =====================
	fmt.Println("\n=== APPEND REALLOCATION ===")

	x := make([]int, 2, 2)
	x[0], x[1] = 1, 2

	y := x[:1] // shares array

	x = append(x, 3) // new array allocated

	y[0] = 99
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	// =====================
	// COPYING SLICES
	// =====================
	fmt.Println("\n=== COPY ===")

	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)

	dst[0] = 99
	fmt.Println("src:", src)
	fmt.Println("dst:", dst)

	// =====================
	// SLICE IN FUNCTIONS
	// =====================
	fmt.Println("\n=== SLICE IN FUNCTION ===")

	nums := []int{5, 6, 7}
	modifySlice(nums)
	fmt.Println("After modifySlice:", nums)

	nums = addToSlice(nums)
	fmt.Println("After addToSlice:", nums)

	// =====================
	// RANGE GOTCHA
	// =====================
	fmt.Println("\n=== RANGE GOTCHA ===")

	r := []int{1, 2, 3}

	for _, v := range r {
		v = 10 // does NOT modify slice
	}
	fmt.Println("Wrong range:", r)

	for i := range r {
		r[i] = 10
	}
	fmt.Println("Correct range:", r)

	// =====================
	// MAPS
	// =====================
	fmt.Println("\n=== MAPS ===")

	ages := make(map[string]int)
	ages["Alice"] = 25
	ages["Bob"] = 30

	fmt.Println("Map:", ages)
	fmt.Println("Alice age:", ages["Alice"])
	fmt.Println("Unknown key:", ages["Eve"]) // zero value

	// comma-ok idiom
	age, ok := ages["Eve"]
	if ok {
		fmt.Println("Eve exists:", age)
	} else {
		fmt.Println("Eve not found")
	}

	// delete
	delete(ages, "Bob")
	fmt.Println("After delete:", ages)

	// =====================
	// MAP OF SLICES
	// =====================
	fmt.Println("\n=== MAP OF SLICES ===")

	groups := make(map[string][]string)

	groups["admins"] = append(groups["admins"], "Alice")
	groups["admins"] = append(groups["admins"], "Bob")
	groups["users"] = append(groups["users"], "Eve")

	fmt.Println("Groups:", groups)
}

// ---------- FUNCTIONS ----------

func changeArray(a [3]int) {
	a[0] = 99 // does NOT affect original
}

func modifySlice(s []int) {
	s[0] = 99 // DOES affect original
}

func addToSlice(s []int) []int {
	return append(s, 100)
}

