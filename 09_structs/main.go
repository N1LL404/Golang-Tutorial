// ============================================================
// LESSON 9: Structs (Custom Data Types)
// ============================================================
// Structs bundle related data together - like classes but simpler!
// TO RUN: go run main.go
// ============================================================

package main

import (
	"fmt"
	"math"
)

// Person represents a human being
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
}

// Rectangle shape
type Rectangle struct {
	Width, Height float64
}

// Circle shape
type Circle struct {
	Radius float64
}

// Methods for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// Methods for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Methods for Person
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p *Person) HaveBirthday() {
	p.Age++
}

// Nested structs
type Address struct {
	City, Country string
}

type Employee struct {
	Person // Embedded (promoted fields)
	Role   string
	Salary float64
	Address
}

func main() {
	fmt.Println("=== Creating Structs ===")

	person1 := Person{"John", "Doe", 30, "john@example.com"}
	fmt.Println("Person 1:", person1)

	person2 := Person{
		FirstName: "Alice",
		LastName:  "Smith",
		Age:       25,
		Email:     "alice@example.com",
	}
	fmt.Println("Person 2:", person2)

	fmt.Println("\n=== Struct Methods ===")
	fmt.Println("Full name:", person2.FullName())

	person2.HaveBirthday()
	fmt.Printf("Age after birthday: %d\n", person2.Age)

	fmt.Println("\n=== Geometry ===")
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Rectangle: %+v, Area: %.2f\n", rect, rect.Area())

	rect.Scale(2)
	fmt.Printf("After scaling: %+v, Area: %.2f\n", rect, rect.Area())

	circle := Circle{Radius: 5}
	fmt.Printf("Circle Area: %.2f\n", circle.Area())

	fmt.Println("\n=== Nested/Embedded Structs ===")
	emp := Employee{
		Person:  Person{FirstName: "Jane", LastName: "Doe", Age: 28},
		Role:    "Engineer",
		Salary:  85000,
		Address: Address{City: "NYC", Country: "USA"},
	}

	// Embedded fields are promoted
	fmt.Println("Name:", emp.FullName())
	fmt.Println("City:", emp.City)
	fmt.Printf("Employee: %+v\n", emp)

	fmt.Println("\n=== Slice of Structs ===")
	people := []Person{
		{FirstName: "A", LastName: "One", Age: 20},
		{FirstName: "B", LastName: "Two", Age: 25},
	}
	for _, p := range people {
		fmt.Printf("- %s (%d)\n", p.FullName(), p.Age)
	}
}
