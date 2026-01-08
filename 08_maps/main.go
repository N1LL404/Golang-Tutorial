// ============================================================
// LESSON 8: Maps (Key-Value Pairs)
// ============================================================
//
// KEY CONCEPTS:
// - Maps store key-value pairs (like dictionaries in Python)
// - Keys must be comparable (strings, ints, etc.)
// - Use make() to create an empty map
// - Maps are not ordered!
//
// SYNTAX:
// map[keyType]valueType
// Example: map[string]int
//
// TO RUN: go run main.go
// ============================================================

package main

import "fmt"

func main() {
	// ==========================================
	// CREATING MAPS
	// ==========================================

	fmt.Println("=== Creating Maps ===")

	// Method 1: Using make()
	ages := make(map[string]int)
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 35
	fmt.Println("Ages:", ages)

	// Method 2: Map literal (initialize with values)
	scores := map[string]int{
		"Math":    95,
		"English": 87,
		"Science": 92,
	}
	fmt.Println("Scores:", scores)

	// ==========================================
	// ACCESSING VALUES
	// ==========================================

	fmt.Println("\n=== Accessing Values ===")

	fmt.Println("Alice's age:", ages["Alice"])
	fmt.Println("Math score:", scores["Math"])

	// Accessing non-existent key returns zero value
	fmt.Println("Unknown person's age:", ages["Unknown"]) // Returns 0

	// ==========================================
	// CHECKING IF KEY EXISTS
	// ==========================================

	fmt.Println("\n=== Checking Key Existence ===")

	// Two-value form: value, ok := map[key]
	age, exists := ages["Alice"]
	if exists {
		fmt.Printf("Alice exists with age %d\n", age)
	}

	age, exists = ages["Unknown"]
	if !exists {
		fmt.Println("Unknown person does not exist")
	}

	// Common pattern: check and use in one step
	if score, ok := scores["History"]; ok {
		fmt.Println("History score:", score)
	} else {
		fmt.Println("History score not found")
	}

	// ==========================================
	// UPDATING VALUES
	// ==========================================

	fmt.Println("\n=== Updating Values ===")

	fmt.Println("Before update - Alice's age:", ages["Alice"])
	ages["Alice"] = 26
	fmt.Println("After update - Alice's age:", ages["Alice"])

	// ==========================================
	// DELETING ENTRIES
	// ==========================================

	fmt.Println("\n=== Deleting Entries ===")

	fmt.Println("Before delete:", ages)
	delete(ages, "Charlie")
	fmt.Println("After delete:", ages)

	// Deleting non-existent key is safe (no error)
	delete(ages, "NonExistent")

	// ==========================================
	// ITERATING OVER MAPS
	// ==========================================

	fmt.Println("\n=== Iterating Over Maps ===")

	countries := map[string]string{
		"US": "United States",
		"UK": "United Kingdom",
		"JP": "Japan",
		"BD": "Bangladesh",
		"IN": "India",
	}

	// Iterate with key and value
	for code, name := range countries {
		fmt.Printf("%s: %s\n", code, name)
	}

	// Just keys
	fmt.Println("\nJust keys:")
	for code := range countries {
		fmt.Printf("- %s\n", code)
	}

	// ==========================================
	// MAP LENGTH
	// ==========================================

	fmt.Println("\n=== Map Length ===")

	fmt.Println("Number of countries:", len(countries))

	// ==========================================
	// NESTED MAPS
	// ==========================================

	fmt.Println("\n=== Nested Maps ===")

	// Map of maps
	students := map[string]map[string]int{
		"Alice": {
			"Math":    95,
			"English": 88,
			"Science": 92,
		},
		"Bob": {
			"Math":    82,
			"English": 79,
			"Science": 85,
		},
	}

	fmt.Println("Alice's Math:", students["Alice"]["Math"])
	fmt.Println("Bob's English:", students["Bob"]["English"])

	// Print all student scores
	for name, subjects := range students {
		fmt.Printf("\n%s's scores:\n", name)
		for subject, score := range subjects {
			fmt.Printf("  %s: %d\n", subject, score)
		}
	}

	// ==========================================
	// MAPS WITH SLICE VALUES
	// ==========================================

	fmt.Println("\n=== Maps with Slice Values ===")

	// Map where values are slices
	hobbies := map[string][]string{
		"Alice": {"Reading", "Gaming", "Cooking"},
		"Bob":   {"Sports", "Music"},
	}

	for person, hobbyList := range hobbies {
		fmt.Printf("%s's hobbies: %v\n", person, hobbyList)
	}

	// Add a hobby
	hobbies["Alice"] = append(hobbies["Alice"], "Painting")
	fmt.Println("Alice's updated hobbies:", hobbies["Alice"])

	// ==========================================
	// PRACTICAL EXAMPLE: Word Counter
	// ==========================================

	fmt.Println("\n=== Practical: Word Counter ===")

	text := "the quick brown fox jumps over the lazy dog the fox is quick"
	words := []string{}

	// Simple word splitting (in real code, use strings.Fields)
	word := ""
	for _, char := range text {
		if char == ' ' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}
	if word != "" {
		words = append(words, word)
	}

	// Count word frequency
	wordCount := make(map[string]int)
	for _, w := range words {
		wordCount[w]++
	}

	fmt.Println("Word frequencies:")
	for w, count := range wordCount {
		fmt.Printf("  '%s': %d\n", w, count)
	}

	// ==========================================
	// NIL MAP
	// ==========================================

	fmt.Println("\n=== Nil Map Warning ===")

	var nilMap map[string]int
	fmt.Println("Nil map:", nilMap)
	fmt.Println("Is nil?", nilMap == nil)

	// Reading from nil map is OK (returns zero value)
	fmt.Println("Reading from nil map:", nilMap["key"])

	// Writing to nil map causes PANIC!
	// nilMap["key"] = 1  // This would crash!

	// Always initialize your maps before writing
	nilMap = make(map[string]int)
	nilMap["key"] = 1 // Now it's safe
	fmt.Println("After initialization:", nilMap)
}
