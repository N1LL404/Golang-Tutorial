// ============================================================
// LESSON 12: Goroutines and Channels (Concurrency)
// ============================================================
// Go's superpower! Run code concurrently with goroutines
// Communicate between goroutines with channels
// TO RUN: go run main.go
// ============================================================

package main

import (
	"fmt"
	"time"
)

func sayHello(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Hello, %s! (%d)\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func sum(nums []int, ch chan int) {
	total := 0
	for _, n := range nums {
		total += n
	}
	ch <- total // Send result to channel
}

func main() {
	fmt.Println("=== Goroutines ===")

	// Start goroutine with 'go' keyword
	go sayHello("Alice")
	go sayHello("Bob")

	// Main function continues immediately
	fmt.Println("Goroutines started!")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("\n=== Channels ===")

	// Create a channel
	ch := make(chan int)

	nums := []int{1, 2, 3, 4, 5, 6}
	mid := len(nums) / 2

	// Sum halves concurrently
	go sum(nums[:mid], ch)
	go sum(nums[mid:], ch)

	// Receive results (blocks until data arrives)
	x, y := <-ch, <-ch
	fmt.Printf("Sum of halves: %d + %d = %d\n", x, y, x+y)

	fmt.Println("\n=== Buffered Channel ===")

	buffered := make(chan string, 2)
	buffered <- "first"
	buffered <- "second"
	fmt.Println(<-buffered)
	fmt.Println(<-buffered)

	fmt.Println("\n=== Select Statement ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "from channel 2"
	}()

	// Select waits on multiple channels
	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("Received:", msg)
		case msg := <-ch2:
			fmt.Println("Received:", msg)
		}
	}

	fmt.Println("\n=== Worker Pattern ===")

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start 2 workers
	for w := 1; w <= 2; w++ {
		go func(id int) {
			for job := range jobs {
				fmt.Printf("Worker %d processing job %d\n", id, job)
				time.Sleep(50 * time.Millisecond)
				results <- job * 2
			}
		}(w)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for r := 1; r <= 5; r++ {
		fmt.Println("Result:", <-results)
	}
}
