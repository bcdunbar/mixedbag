package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create a buffered channel to receive the results
	resultChan := make(chan time.Duration, 2)

	// Create a wait group to wait for the goroutines to finish
	var wg sync.WaitGroup

	// Run the long-running tasks asynchronously
	wg.Add(1)
	go async(&wg, resultChan, "Async")

	// Run the long-running tasks synchronously
	wg.Add(1)
	go synchronously(&wg, resultChan, "Sync")


	// Print the results
	result := <-resultChan
	fmt.Printf("The async task took %v\n", result)

	resultTwo := <-resultChan
	fmt.Printf("The sync task took %v\n", resultTwo)
}

func async(wg *sync.WaitGroup, resultChan chan<- time.Duration, taskName string) {
	defer wg.Done()

	startTime := time.Now()
	fmt.Println("Starting async task...")

	// Create a buffered channel to receive the results as time
	resultInnerChan := make(chan int, 2)

	// Create a wait group to wait for the goroutines to finish
	var wgInner sync.WaitGroup

	// Start the first goroutine
	wgInner.Add(1)
	go longRunningAsyncTask(&wgInner, resultInnerChan, "Task 1", 2)

	// Start the second goroutine
	wgInner.Add(1)
	go longRunningAsyncTask(&wgInner, resultInnerChan, "Task 2", 3)

	// Do some other work while the goroutines are executing
	// fmt.Println("Doing some other work...")

	// Wait for the goroutines to finish
	wgInner.Wait()

	// Print the results
	// result := <-resultInnerChan
	// resultTwo := <-resultInnerChan

	resultChan <- time.Since(startTime)

}

func synchronously(wg *sync.WaitGroup, resultChan chan<- time.Duration, taskName string) {
	defer wg.Done()

	startTime := time.Now()
	fmt.Println("Starting sync task...")

    // Run the first long-running task synchronously
    longRunningTask(2, "Task 1")

    // Run the second long-running task synchronously
    longRunningTask(3, "Task 2")

	resultChan <- time.Since(startTime)
}

func longRunningAsyncTask(wg *sync.WaitGroup, resultChan chan<- int, taskName string, num int) {
    defer wg.Done()

    // Multiply the number by itself 10 times
    result := num
    for i := 0; i < 10; i++ {
        result *= num
        // fmt.Printf("%s: %d\n", taskName, result)
        time.Sleep(1 * time.Second)
    }

    // Completed
    // fmt.Printf("%s completed! Result: %d\n", taskName, result)

    // Send the result on the channel
    resultChan <- result
}

func longRunningTask(num int, taskName string) int {
    // Multiply the number by itself 10 times
    result := num
    for i := 0; i < 10; i++ {
        result *= num
        // fmt.Printf("Task: %d\n", result)
        time.Sleep(1 * time.Second)
    }

    // Completed
    // fmt.Printf("%s completed! Result: %d\n", taskName, result)

    return result
}