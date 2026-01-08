// ============================================================
// LESSON 11: Error Handling
// ============================================================
// Go handles errors explicitly with return values (no exceptions!)
// The error type is built-in and widely used
// TO RUN: go run main.go
// ============================================================

package main

import (
	"errors"
	"fmt"
)

// Function that returns an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// Custom error with more info
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func validateAge(age int) error {
	if age < 0 {
		return ValidationError{"age", "cannot be negative"}
	}
	if age > 150 {
		return ValidationError{"age", "seems unrealistic"}
	}
	return nil
}

// Function that wraps errors
func processUser(name string, age int) error {
	if name == "" {
		return fmt.Errorf("invalid name: %w", errors.New("name cannot be empty"))
	}
	if err := validateAge(age); err != nil {
		return fmt.Errorf("invalid user data: %w", err)
	}
	return nil
}

func main() {
	fmt.Println("=== Basic Error Handling ===")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\n=== Custom Errors ===")

	if err := validateAge(-5); err != nil {
		fmt.Println("Validation error:", err)
	}

	if err := validateAge(200); err != nil {
		fmt.Println("Validation error:", err)
	}

	if err := validateAge(25); err == nil {
		fmt.Println("Age 25 is valid!")
	}

	fmt.Println("\n=== Wrapped Errors ===")

	if err := processUser("", 25); err != nil {
		fmt.Println("Error:", err)
	}

	if err := processUser("John", -1); err != nil {
		fmt.Println("Error:", err)

		// Unwrap to check original error
		var valErr ValidationError
		if errors.As(err, &valErr) {
			fmt.Printf("  Field: %s, Message: %s\n", valErr.Field, valErr.Message)
		}
	}

	fmt.Println("\n=== Error Patterns ===")

	// Pattern 1: Return early on error
	process := func() error {
		_, err := divide(10, 0)
		if err != nil {
			return err
		}
		fmt.Println("This won't print")
		return nil
	}

	if err := process(); err != nil {
		fmt.Println("Process failed:", err)
	}

	// Pattern 2: Ignore error (use _ but be careful!)
	result, _ = divide(10, 5) // We're sure this won't error
	fmt.Println("Ignoring error (safe case):", result)

	// Pattern 3: Check specific error types
	var ErrNotFound = errors.New("not found")
	checkErr := func(err error) {
		if errors.Is(err, ErrNotFound) {
			fmt.Println("Item was not found")
		}
	}
	checkErr(ErrNotFound)
}
