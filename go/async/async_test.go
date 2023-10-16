package main

import (
    "testing"
    "time"
	"sync"
)

func TestAsync(t *testing.T) {
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

    // Wait for the results
    result := <-resultChan
    resultTwo := <-resultChan

    // Check that the async task completed faster than the sync task
    if result >= resultTwo {
        t.Errorf("Async task did not complete faster than sync task. Async: %v, Sync: %v", result, resultTwo)
    }
}