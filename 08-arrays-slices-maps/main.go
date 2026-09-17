package main

import "fmt"

func main() {
	// =========================
	// Arrays
	// =========================

	fmt.Println("=== Arrays ===")

	numbers := [3]int{10, 20, 30}

	fmt.Println("array:", numbers)
	fmt.Println("length:", len(numbers))

	numbers[1] = 99

	fmt.Println("modified array:", numbers)

	// This would panic because the array has only 3 elements:
	// numbers[3] = 100

	// =========================
	// Slices
	// =========================

	fmt.Println("\n=== Slices ===")

	users := []string{"Daniel", "Alex", "John"}

	fmt.Println("slice:", users)
	fmt.Println("length:", len(users))
	fmt.Println("capacity:", cap(users))

	users[1] = "Mike"

	fmt.Println("modified slice:", users)

	users = append(users, "Sarah")

	fmt.Println("after append:", users)
	fmt.Println("length:", len(users))
	fmt.Println("capacity:", cap(users))

	// =========================
	// Slice sharing
	// =========================

	fmt.Println("\n=== Slice Sharing ===")

	original := []int{10, 20, 30}

	copyOfSlice := original

	copyOfSlice[0] = 999

	fmt.Println("original:", original)
	fmt.Println("copy:", copyOfSlice)

	// Both slices can point to the same backing array.

	// =========================
	// Slice copy
	// =========================

	fmt.Println("\n=== Copy ===")

	source := []int{1, 2, 3}
	destination := make([]int, len(source))

	copied := copy(destination, source)

	destination[0] = 100

	fmt.Println("source:", source)
	fmt.Println("destination:", destination)
	fmt.Println("elements copied:", copied)

	// =========================
	// Nil slice
	// =========================

	fmt.Println("\n=== Nil Slice ===")

	var empty []int

	fmt.Println("slice:", empty)
	fmt.Println("is nil:", empty == nil)
	fmt.Println("length:", len(empty))
	fmt.Println("capacity:", cap(empty))

	empty = append(empty, 10)

	fmt.Println("after append:", empty)

	// =========================
	// Maps
	// =========================

	fmt.Println("\n=== Maps ===")

	usersMap := map[string]string{
		"1": "Daniel",
		"2": "Alex",
		"3": "John",
	}

	fmt.Println("map:", usersMap)

	usersMap["2"] = "Mike"

	fmt.Println("after update:", usersMap)

	delete(usersMap, "3")

	fmt.Println("after delete:", usersMap)

	// =========================
	// Map lookup
	// =========================

	fmt.Println("\n=== Map Lookup ===")

	name := usersMap["1"]
	fmt.Println("name:", name)

	missingName := usersMap["999"]
	fmt.Println("missing name:", missingName)

	name, ok := usersMap["999"]

	fmt.Println("name:", name)
	fmt.Println("exists:", ok)

	// =========================
	// Nil map
	// =========================

	fmt.Println("\n=== Nil Map ===")

	var nilMap map[string]string

	fmt.Println("map:", nilMap)
	fmt.Println("is nil:", nilMap == nil)

	// Reading from a nil map is safe.
	fmt.Println("missing value:", nilMap["1"])

	// Writing to a nil map would panic:
	// nilMap["1"] = "Daniel"

	// Initialize it before writing.
	nilMap = make(map[string]string)

	nilMap["1"] = "Daniel"

	fmt.Println("initialized map:", nilMap)
}