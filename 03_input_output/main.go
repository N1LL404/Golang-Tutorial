// ============================================================
// LESSON 3: User Input and Formatted Output
// ============================================================
//
// KEY CONCEPTS:
// - fmt.Scanf() reads input from the user
// - fmt.Printf() prints with formatting placeholders
// - & operator gets the memory address (needed for Scanf)
//
// COMMON FORMAT VERBS:
// %s = string
// %d = integer (decimal)
// %f = float (decimal number)
// %.2f = float with 2 decimal places
// %v = any value (auto-detect type)
// %T = type of the variable
// \n = new line
//
// TO RUN: go run main.go
// ============================================================

package main

import "fmt"

func main() {
    // ==========================================
    // FORMATTED PRINTING with Printf
    // ==========================================
    
    name := "Gopher"
    age := 10
    score := 95.75
    
    fmt.Println("=== Printf Formatting ===")
    
    // %s for strings
    fmt.Printf("Hello, %s!\n", name)
    
    // %d for integers
    fmt.Printf("%s is %d years old.\n", name, age)
    
    // %f for floats (shows many decimals)
    fmt.Printf("Score: %f\n", score)
    
    // %.2f for floats with 2 decimal places
    fmt.Printf("Score (2 decimals): %.2f\n", score)
    
    // %v auto-detects the type
    fmt.Printf("Name: %v, Age: %v, Score: %v\n", name, age, score)
    
    // %T shows the type
    fmt.Printf("Type of name: %T\n", name)
    fmt.Printf("Type of age: %T\n", age)
    fmt.Printf("Type of score: %T\n", score)
    
    // ==========================================
    // USER INPUT with Scanf
    // ==========================================
    
    fmt.Println("\n=== User Input Demo ===")
    
    var userName string
    var userAge int
    
    // Reading a string
    fmt.Print("Enter your name: ")
    fmt.Scanf("%s", &userName)  // & gets the memory address
    
    // Reading an integer
    fmt.Print("Enter your age: ")
    fmt.Scanf("%d", &userAge)
    
    // Display the input
    fmt.Printf("\nHello, %s! You are %d years old.\n", userName, userAge)
    
    // ==========================================
    // CALCULATING AND DISPLAYING
    // ==========================================
    
    fmt.Println("\n=== Simple Calculator ===")
    
    var num1, num2 float64
    
    fmt.Print("Enter first number: ")
    fmt.Scanf("%f", &num1)
    
    fmt.Print("Enter second number: ")
    fmt.Scanf("%f", &num2)
    
    sum := num1 + num2
    difference := num1 - num2
    product := num1 * num2
    
    fmt.Printf("\nResults:\n")
    fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, sum)
    fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, difference)
    fmt.Printf("%.2f × %.2f = %.2f\n", num1, num2, product)
    
    if num2 != 0 {
        quotient := num1 / num2
        fmt.Printf("%.2f ÷ %.2f = %.2f\n", num1, num2, quotient)
    } else {
        fmt.Println("Cannot divide by zero!")
    }
}
