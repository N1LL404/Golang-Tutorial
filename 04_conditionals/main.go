// ============================================================
// LESSON 4: Conditionals (if, else, switch)
// ============================================================
//
// KEY CONCEPTS:
// - if/else lets you make decisions in your code
// - Comparison operators: == != < > <= >=
// - Logical operators: && (and), || (or), ! (not)
// - switch is great for multiple conditions
//
// NOTE: Go does NOT require parentheses around conditions!
//       This is different from many other languages.
//
// TO RUN: go run main.go
// ============================================================

package main

import "fmt"

func main() {
    // ==========================================
    // SIMPLE IF STATEMENT
    // ==========================================
    
    age := 20
    
    fmt.Println("=== Simple If ===")
    
    if age >= 18 {
        fmt.Println("You are an adult!")
    }
    
    // ==========================================
    // IF-ELSE
    // ==========================================
    
    fmt.Println("\n=== If-Else ===")
    
    temperature := 35
    
    if temperature > 30 {
        fmt.Println("It's hot outside! 🔥")
    } else {
        fmt.Println("It's not too hot.")
    }
    
    // ==========================================
    // IF-ELSE IF-ELSE (Multiple Conditions)
    // ==========================================
    
    fmt.Println("\n=== If-Else If-Else ===")
    
    score := 85
    
    if score >= 90 {
        fmt.Println("Grade: A - Excellent!")
    } else if score >= 80 {
        fmt.Println("Grade: B - Good job!")
    } else if score >= 70 {
        fmt.Println("Grade: C - Not bad")
    } else if score >= 60 {
        fmt.Println("Grade: D - Need improvement")
    } else {
        fmt.Println("Grade: F - Keep trying!")
    }
    
    // ==========================================
    // COMPARISON OPERATORS
    // ==========================================
    
    fmt.Println("\n=== Comparison Operators ===")
    
    a, b := 10, 20
    
    fmt.Printf("%d == %d: %v\n", a, b, a == b)  // Equal
    fmt.Printf("%d != %d: %v\n", a, b, a != b)  // Not equal
    fmt.Printf("%d < %d: %v\n", a, b, a < b)    // Less than
    fmt.Printf("%d > %d: %v\n", a, b, a > b)    // Greater than
    fmt.Printf("%d <= %d: %v\n", a, b, a <= b)  // Less than or equal
    fmt.Printf("%d >= %d: %v\n", a, b, a >= b)  // Greater than or equal
    
    // ==========================================
    // LOGICAL OPERATORS
    // ==========================================
    
    fmt.Println("\n=== Logical Operators ===")
    
    hasLicense := true
    isAdult := true
    isTired := false
    
    // AND (&&) - both must be true
    if hasLicense && isAdult {
        fmt.Println("✓ You can drive (has license AND is adult)")
    }
    
    // OR (||) - at least one must be true
    if isTired || !hasLicense {
        fmt.Println("Maybe don't drive right now")
    } else {
        fmt.Println("✓ Good to drive!")
    }
    
    // NOT (!) - reverses the boolean
    fmt.Printf("isTired: %v, !isTired: %v\n", isTired, !isTired)
    
    // ==========================================
    // IF WITH SHORT STATEMENT
    // (Declare variable in the if statement)
    // ==========================================
    
    fmt.Println("\n=== If with Short Statement ===")
    
    // The variable 'length' only exists inside this if block
    if length := len("Hello"); length > 3 {
        fmt.Printf("The word has %d characters (more than 3)\n", length)
    }
    
    // ==========================================
    // SWITCH STATEMENT
    // (Cleaner than many if-else statements)
    // ==========================================
    
    fmt.Println("\n=== Switch Statement ===")
    
    day := "Wednesday"
    
    switch day {
    case "Monday":
        fmt.Println("Start of the work week 😴")
    case "Tuesday", "Wednesday", "Thursday":
        fmt.Println("Middle of the week 💪")
    case "Friday":
        fmt.Println("TGIF! 🎉")
    case "Saturday", "Sunday":
        fmt.Println("Weekend! 🎊")
    default:
        fmt.Println("Invalid day")
    }
    
    // ==========================================
    // SWITCH WITHOUT EXPRESSION
    // (Like a cleaner if-else chain)
    // ==========================================
    
    fmt.Println("\n=== Switch Without Expression ===")
    
    hour := 14
    
    switch {
    case hour < 12:
        fmt.Println("Good morning! ☀️")
    case hour < 17:
        fmt.Println("Good afternoon! 🌤️")
    case hour < 21:
        fmt.Println("Good evening! 🌆")
    default:
        fmt.Println("Good night! 🌙")
    }
}
