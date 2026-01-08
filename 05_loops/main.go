// ============================================================
// LESSON 5: Loops (for is the only loop in Go!)
// ============================================================
//
// KEY CONCEPTS:
// - Go only has ONE loop keyword: for
// - But it can work like: for, while, and infinite loops
// - break exits the loop
// - continue skips to next iteration
// - range iterates over collections
//
// TO RUN: go run main.go
// ============================================================

package main

import "fmt"

func main() {
	// ==========================================
	// CLASSIC FOR LOOP
	// for initialization; condition; post { }
	// ==========================================

	fmt.Println("=== Classic For Loop ===")

	for i := 1; i <= 5; i++ {
		fmt.Printf("Count: %d\n", i)
	}

	// ==========================================
	// FOR AS A WHILE LOOP
	// (Just the condition, no init or post)
	// ==========================================

	fmt.Println("\n=== For as While Loop ===")

	count := 0
	for count < 3 {
		fmt.Printf("While count: %d\n", count)
		count++
	}

	// ==========================================
	// INFINITE LOOP (with break)
	// ==========================================

	fmt.Println("\n=== Infinite Loop with Break ===")

	num := 0
	for {
		num++
		fmt.Printf("Infinite loop iteration: %d\n", num)
		if num >= 3 {
			fmt.Println("Breaking out!")
			break // Exit the loop
		}
	}

	// ==========================================
	// CONTINUE - Skip to Next Iteration
	// ==========================================

	fmt.Println("\n=== Continue (Skip Even Numbers) ===")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 { // If even number
			continue // Skip the rest, go to next iteration
		}
		fmt.Printf("Odd number: %d\n", i)
	}

	// ==========================================
	// LOOPING THROUGH STRINGS
	// ==========================================

	fmt.Println("\n=== Looping Through String ===")

	greeting := "Hello"
	for index, char := range greeting {
		fmt.Printf("Index %d: %c\n", index, char)
	}

	// ==========================================
	// LOOPING THROUGH ARRAYS/SLICES
	// ==========================================

	fmt.Println("\n=== Looping Through Slice ===")

	fruits := []string{"Apple", "Banana", "Cherry", "Date"}

	// Using range (recommended)
	for index, fruit := range fruits {
		fmt.Printf("%d: %s\n", index, fruit)
	}

	// If you only need the value, use _ for index
	fmt.Println("\nJust values:")
	for _, fruit := range fruits {
		fmt.Println("- " + fruit)
	}

	// If you only need the index
	fmt.Println("\nJust indices:")
	for i := range fruits {
		fmt.Printf("Index: %d\n", i)
	}

	// ==========================================
	// NESTED LOOPS
	// ==========================================

	fmt.Println("\n=== Nested Loops (Multiplication Table) ===")

	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Printf("%d×%d=%2d  ", i, j, i*j)
		}
		fmt.Println() // New line after each row
	}

	// ==========================================
	// LOOP WITH LABEL (for breaking nested loops)
	// ==========================================

	fmt.Println("\n=== Loop with Label ===")

outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
			if i == 2 && j == 2 {
				fmt.Println("Breaking outer loop!")
				break outer // Breaks out of the outer loop
			}
		}
	}

	// ==========================================
	// PRACTICAL EXAMPLE: Sum of Numbers
	// ==========================================

	fmt.Println("\n=== Practical: Sum of 1 to 100 ===")

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i // Same as: sum = sum + i
	}
	fmt.Printf("Sum of 1 to 100 = %d\n", sum)
}
