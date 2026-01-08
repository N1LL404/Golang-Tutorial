// ============================================================
// LESSON 2: Variables and Data Types
// ============================================================
//
// KEY CONCEPTS:
// - Variables store data in your program
// - Go is statically typed (each variable has a specific type)
// - Use := for short variable declaration (Go infers the type)
// - Use var for explicit declaration
//
// BASIC DATA TYPES:
// - string: text like "Hello"
// - int: whole numbers like 42, -10
// - float64: decimal numbers like 3.14, -0.5
// - bool: true or false
//
// TO RUN: go run main.go
// ============================================================

package main

import "fmt"

func main() {
    // ==========================================
    // METHOD 1: Short Declaration with := 
    // (Most common, Go figures out the type)
    // ==========================================
    
    name := "Alice"           // string (text)
    age := 25                 // int (integer/whole number)
    height := 5.8             // float64 (decimal number)
    isStudent := true         // bool (boolean: true/false)
    
    fmt.Println("=== Short Declaration (:=) ===")
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Height:", height)
    fmt.Println("Is Student:", isStudent)
    
    // ==========================================
    // METHOD 2: Explicit Declaration with var
    // (When you want to be specific about the type)
    // ==========================================
    
    var city string = "New York"
    var population int = 8000000
    var temperature float64 = 72.5
    var isCapital bool = false
    
    fmt.Println("\n=== Explicit Declaration (var) ===")
    fmt.Println("City:", city)
    fmt.Println("Population:", population)
    fmt.Println("Temperature:", temperature)
    fmt.Println("Is Capital:", isCapital)
    
    // ==========================================
    // METHOD 3: Declare first, assign later
    // (Variables get "zero values" by default)
    // ==========================================
    
    var score int          // Default: 0
    var greeting string    // Default: "" (empty string)
    var isActive bool      // Default: false
    var price float64      // Default: 0.0
    
    fmt.Println("\n=== Zero Values (defaults) ===")
    fmt.Println("Score (int default):", score)
    fmt.Println("Greeting (string default):", greeting)
    fmt.Println("IsActive (bool default):", isActive)
    fmt.Println("Price (float64 default):", price)
    
    // Now assign values
    score = 100
    greeting = "Hello!"
    isActive = true
    price = 29.99
    
    fmt.Println("\n=== After Assignment ===")
    fmt.Println("Score:", score)
    fmt.Println("Greeting:", greeting)
    fmt.Println("IsActive:", isActive)
    fmt.Println("Price:", price)
    
    // ==========================================
    // MULTIPLE VARIABLE DECLARATION
    // ==========================================
    
    var (
        firstName = "John"
        lastName  = "Doe"
        email     = "john@example.com"
    )
    
    fmt.Println("\n=== Multiple Variables ===")
    fmt.Println("First Name:", firstName)
    fmt.Println("Last Name:", lastName)
    fmt.Println("Email:", email)
    
    // Multiple on one line
    x, y, z := 10, 20, 30
    fmt.Println("\n=== Multiple on One Line ===")
    fmt.Println("x, y, z:", x, y, z)
}
