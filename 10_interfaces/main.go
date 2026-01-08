// ============================================================
// LESSON 10: Interfaces
// ============================================================
// Interfaces define behavior (what methods a type must have)
// A type implements an interface by implementing its methods
// TO RUN: go run main.go
// ============================================================

package main

import (
	"fmt"
	"math"
)

// Shape interface - any type with Area() method implements this
type Shape interface {
	Area() float64
}

// Describer - any type with Describe() method
type Describer interface {
	Describe() string
}

// Rectangle implements Shape
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Describe() string {
	return fmt.Sprintf("Rectangle %.1fx%.1f", r.Width, r.Height)
}

// Circle implements Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Describe() string {
	return fmt.Sprintf("Circle radius %.1f", c.Radius)
}

// Function that works with ANY Shape
func printArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

// Combined interface
type ShapeDescriber interface {
	Shape
	Describer
}

func describeShape(sd ShapeDescriber) {
	fmt.Printf("%s has area %.2f\n", sd.Describe(), sd.Area())
}

// Empty interface - accepts ANY type
func printAnything(v interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", v, v)
}

func main() {
	fmt.Println("=== Interfaces ===")

	rect := Rectangle{10, 5}
	circle := Circle{7}

	// Both implement Shape
	printArea(rect)
	printArea(circle)

	fmt.Println("\n=== Combined Interface ===")
	describeShape(rect)
	describeShape(circle)

	fmt.Println("\n=== Slice of Interfaces ===")
	shapes := []Shape{rect, circle, Rectangle{3, 4}}
	for _, s := range shapes {
		fmt.Printf("Shape area: %.2f\n", s.Area())
	}

	fmt.Println("\n=== Empty Interface (any type) ===")
	printAnything(42)
	printAnything("hello")
	printAnything(true)
	printAnything(rect)

	fmt.Println("\n=== Type Assertion ===")
	var s Shape = rect
	r, ok := s.(Rectangle)
	if ok {
		fmt.Printf("It's a Rectangle with width %.1f\n", r.Width)
	}

	fmt.Println("\n=== Type Switch ===")
	checkType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Println("Integer:", v)
		case string:
			fmt.Println("String:", v)
		case Shape:
			fmt.Println("Shape with area:", v.Area())
		default:
			fmt.Println("Unknown type")
		}
	}

	checkType(100)
	checkType("hello")
	checkType(circle)
}
