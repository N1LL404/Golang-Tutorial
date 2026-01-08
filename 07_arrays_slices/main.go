// ============================================================
// LESSON 7: Arrays and Slices
// ============================================================
//
// KEY CONCEPTS:
// - Arrays have FIXED size (rarely used in practice)
// - Slices are DYNAMIC and flexible (used everywhere!)
// - Slices are like "windows" into arrays
// - Use make() to create slices with specific capacity
//
// ARRAYS vs SLICES:
// - Array: [5]int{1,2,3,4,5} - fixed size of 5
// - Slice: []int{1,2,3,4,5} - dynamic, can grow/shrink
//
// TO RUN: go run main.go
// ============================================================

package main

import "fmt"

func main() {
	// ==========================================
	// ARRAYS (Fixed Size)
	// ==========================================

	fmt.Println("=== Arrays (Fixed Size) ===")

	// Declare an array of 5 integers
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	fmt.Println("Array:", numbers)
	fmt.Println("Length:", len(numbers))

	// Initialize array with values
	fruits := [3]string{"Apple", "Banana", "Cherry"}
	fmt.Println("Fruits:", fruits)

	// Let Go count the elements with [...]
	colors := [...]string{"Red", "Green", "Blue", "Yellow"}
	fmt.Println("Colors:", colors)
	fmt.Println("Colors length:", len(colors))

	// ==========================================
	// SLICES (Dynamic - This is what you'll use!)
	// ==========================================

	fmt.Println("\n=== Slices (Dynamic) ===")

	// Create a slice (no size specified)
	scores := []int{95, 87, 92, 88, 91}
	fmt.Println("Scores:", scores)
	fmt.Println("Length:", len(scores))
	fmt.Println("Capacity:", cap(scores))

	// ==========================================
	// APPEND - Add Elements to Slice
	// ==========================================

	fmt.Println("\n=== Append Elements ===")

	animals := []string{"Dog", "Cat"}
	fmt.Println("Before append:", animals)

	animals = append(animals, "Bird")
	fmt.Println("After append:", animals)

	// Append multiple elements
	animals = append(animals, "Fish", "Rabbit", "Hamster")
	fmt.Println("After multiple append:", animals)

	// ==========================================
	// SLICING - Getting Parts of a Slice
	// slice[start:end] - includes start, excludes end
	// ==========================================

	fmt.Println("\n=== Slicing ===")

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Original:", nums)
	fmt.Println("nums[2:5]:", nums[2:5]) // Elements 2,3,4
	fmt.Println("nums[:4]:", nums[:4])   // First 4 elements
	fmt.Println("nums[6:]:", nums[6:])   // From index 6 to end
	fmt.Println("nums[:]:", nums[:])     // All elements (copy)

	// ==========================================
	// MAKE - Create Slice with Length/Capacity
	// ==========================================

	fmt.Println("\n=== Make Function ===")

	// make(type, length, capacity)
	slice1 := make([]int, 5)     // length 5, capacity 5
	slice2 := make([]int, 3, 10) // length 3, capacity 10

	fmt.Println("slice1:", slice1, "len:", len(slice1), "cap:", cap(slice1))
	fmt.Println("slice2:", slice2, "len:", len(slice2), "cap:", cap(slice2))

	// ==========================================
	// COPY - Copy Elements Between Slices
	// ==========================================

	fmt.Println("\n=== Copy Function ===")

	source := []int{1, 2, 3, 4, 5}
	destination := make([]int, len(source))

	copied := copy(destination, source)
	fmt.Println("Source:", source)
	fmt.Println("Destination:", destination)
	fmt.Printf("Copied %d elements\n", copied)

	// ==========================================
	// REMOVE ELEMENT FROM SLICE
	// ==========================================

	fmt.Println("\n=== Remove Element ===")

	letters := []string{"a", "b", "c", "d", "e"}
	fmt.Println("Before removal:", letters)

	// Remove element at index 2 ("c")
	indexToRemove := 2
	letters = append(letters[:indexToRemove], letters[indexToRemove+1:]...)
	fmt.Println("After removing index 2:", letters)

	// ==========================================
	// ITERATING OVER SLICES
	// ==========================================

	fmt.Println("\n=== Iterating Over Slices ===")

	cities := []string{"Tokyo", "Paris", "New York", "London"}

	// Using range
	fmt.Println("With index and value:")
	for i, city := range cities {
		fmt.Printf("  %d: %s\n", i, city)
	}

	// Just values
	fmt.Println("Just values:")
	for _, city := range cities {
		fmt.Printf("  - %s\n", city)
	}

	// Classic for loop
	fmt.Println("Classic for loop:")
	for i := 0; i < len(cities); i++ {
		fmt.Printf("  [%d] %s\n", i, cities[i])
	}

	// ==========================================
	// 2D SLICES (Slice of Slices)
	// ==========================================

	fmt.Println("\n=== 2D Slices ===")

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Matrix:")
	for i, row := range matrix {
		fmt.Printf("  Row %d: %v\n", i, row)
	}
	fmt.Printf("Element at [1][2]: %d\n", matrix[1][2]) // 6

	// ==========================================
	// PRACTICAL EXAMPLE: Filter Slice
	// ==========================================

	fmt.Println("\n=== Practical: Filter Even Numbers ===")

	allNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNumbers := []int{}

	for _, n := range allNumbers {
		if n%2 == 0 {
			evenNumbers = append(evenNumbers, n)
		}
	}

	fmt.Println("All numbers:", allNumbers)
	fmt.Println("Even numbers:", evenNumbers)

	// ==========================================
	// NIL SLICE
	// ==========================================

	fmt.Println("\n=== Nil Slice ===")

	var nilSlice []int
	fmt.Println("Nil slice:", nilSlice)
	fmt.Println("Is nil?", nilSlice == nil)
	fmt.Println("Length:", len(nilSlice))

	// You can still append to nil slice!
	nilSlice = append(nilSlice, 1, 2, 3)
	fmt.Println("After append:", nilSlice)
}
