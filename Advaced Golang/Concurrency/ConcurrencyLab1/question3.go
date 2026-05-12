package main

/* What should be the order of statements to prevent the error in the below code? */
/* Expected Output: Inside Goroutine */

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// Line 1  wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		fmt.Println("Inside Goroutine")
		// Line 2 wg.Done()
	}()
	// Line 3 wg.Wait()
}
