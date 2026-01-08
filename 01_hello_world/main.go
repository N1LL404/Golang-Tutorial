// ============================================================
// LESSON 1: Hello World - Your First Go Program
// ============================================================
// 
// KEY CONCEPTS:
// - Every Go program starts with a package declaration
// - The "main" package is special - it creates an executable program
// - The main() function is the entry point (where your program starts)
// - We import packages to use additional functionality
//
// TO RUN THIS PROGRAM:
// Open terminal in this folder and type: go run main.go
// ============================================================

package main // This declares that this file belongs to the "main" package

import "fmt" // "fmt" is the formatting package - used for printing output

// main() is the entry point of every Go executable program
func main() {
    // Println prints a line of text to the console
    fmt.Println("Hello, World!")
    fmt.Println("Welcome to Go programming! 🚀")
    
    // You can print multiple things on one line
    fmt.Println("Go is", "simple", "and", "powerful!")
}
