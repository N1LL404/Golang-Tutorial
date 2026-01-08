// ============================================================
// LESSON 6: Functions
// ============================================================
//
// KEY CONCEPTS:
// - Functions are reusable blocks of code
// - func keyword declares a function
// - Functions can take parameters and return values
// - Go can return MULTIPLE values (unique feature!)
// - Named return values can make code cleaner
//
// SYNTAX:
// func functionName(param1 type, param2 type) returnType {
//     // code
//     return value
// }
//
// TO RUN: go run main.go
// ============================================================

package main

import (
	"fmt"
	"strings"
)

// ==========================================
// SIMPLE FUNCTION (no parameters, no return)
// ==========================================

func sayHello() {
	fmt.Println("Hello from a function!")
}

// ==========================================
// FUNCTION WITH PARAMETERS
// ==========================================

func greet(name string) {
	fmt.Printf("Hello, %s! Welcome to Go!\n", name)
}

// Same type parameters can be shortened
func greetTwo(firstName, lastName string) {
	fmt.Printf("Hello, %s %s!\n", firstName, lastName)
}

// ==========================================
// FUNCTION WITH RETURN VALUE
// ==========================================

func add(a, b int) int {
	return a + b
}

func square(x int) int {
	return x * x
}

// ==========================================
// FUNCTION WITH MULTIPLE RETURN VALUES
// (This is a unique and powerful Go feature!)
// ==========================================

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil // nil means "no error"
}

func getMinMax(numbers []int) (int, int) {
	if len(numbers) == 0 {
		return 0, 0
	}

	min, max := numbers[0], numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max
}

// ==========================================
// NAMED RETURN VALUES
// (Makes code self-documenting)
// ==========================================

func calculate(a, b int) (sum int, difference int, product int) {
	sum = a + b
	difference = a - b
	product = a * b
	return // "naked return" - returns the named values
}

// ==========================================
// VARIADIC FUNCTIONS
// (Accept any number of arguments)
// ==========================================

func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// ==========================================
// FUNCTIONS AS VALUES
// (Functions can be stored in variables!)
// ==========================================

func multiply(a, b int) int {
	return a * b
}

// ==========================================
// HIGHER-ORDER FUNCTION
// (Function that takes a function as parameter)
// ==========================================

func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// ==========================================
// PRACTICAL EXAMPLE FUNCTIONS
// ==========================================

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1) // Recursion!
}

// ==========================================
// MAIN FUNCTION
// ==========================================

func main() {
	// Simple function call
	fmt.Println("=== Simple Function ===")
	sayHello()

	// Function with parameters
	fmt.Println("\n=== Function with Parameters ===")
	greet("Alice")
	greetTwo("John", "Doe")

	// Function with return value
	fmt.Println("\n=== Function with Return Value ===")
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)
	fmt.Printf("7 squared = %d\n", square(7))

	// Multiple return values
	fmt.Println("\n=== Multiple Return Values ===")
	quotient, err := divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", quotient)
	}

	// Handling error case
	quotient, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Min and Max
	nums := []int{5, 2, 9, 1, 7, 3}
	min, max := getMinMax(nums)
	fmt.Printf("Numbers: %v\n", nums)
	fmt.Printf("Min: %d, Max: %d\n", min, max)

	// Named return values
	fmt.Println("\n=== Named Return Values ===")
	sum, diff, prod := calculate(10, 4)
	fmt.Printf("10 and 4: sum=%d, difference=%d, product=%d\n", sum, diff, prod)

	// Variadic functions
	fmt.Println("\n=== Variadic Functions ===")
	fmt.Printf("Sum of 1,2,3 = %d\n", sumAll(1, 2, 3))
	fmt.Printf("Sum of 1,2,3,4,5 = %d\n", sumAll(1, 2, 3, 4, 5))
	fmt.Printf("Sum of nothing = %d\n", sumAll())

	// Functions as values
	fmt.Println("\n=== Functions as Values ===")
	mathFunc := multiply // Store function in variable
	fmt.Printf("Using function variable: 4 × 5 = %d\n", mathFunc(4, 5))

	// Anonymous function (lambda)
	double := func(x int) int {
		return x * 2
	}
	fmt.Printf("Anonymous function: double(7) = %d\n", double(7))

	// Higher-order function
	fmt.Println("\n=== Higher-Order Function ===")
	fmt.Printf("Apply add: %d\n", applyOperation(10, 5, add))
	fmt.Printf("Apply multiply: %d\n", applyOperation(10, 5, multiply))

	// Practical examples
	fmt.Println("\n=== Practical Examples ===")

	words := []string{"radar", "hello", "level", "world", "civic"}
	for _, word := range words {
		if isPalindrome(word) {
			fmt.Printf("'%s' is a palindrome ✓\n", word)
		} else {
			fmt.Printf("'%s' is NOT a palindrome ✗\n", word)
		}
	}

	fmt.Printf("\nFactorial of 5 = %d\n", factorial(5))
	fmt.Printf("Factorial of 10 = %d\n", factorial(10))
}
