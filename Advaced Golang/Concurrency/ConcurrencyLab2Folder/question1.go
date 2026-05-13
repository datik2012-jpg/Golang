/* Complete the code below to create two goroutines that will run concurrently using a WaitGroup. The first goroutine should print the numbers 1 to 5, and the second goroutine should print the letters a to e. The main function should wait for both goroutines to finish before exiting. */
/*Expected Output: need to guess */

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// Create first goroutine here
	// Create second goroutine here
	wg.Add(2)
	go printNumbers(&wg)
	go printLetters(&wg)

	// Wait for both goroutines to finish here

	wg.Wait() //we are waiting in the main - cause we do not want it will be closing
}

func printNumbers(wg *sync.WaitGroup) {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func printLetters(wg *sync.WaitGroup) {
	chatsslice := []string{"a", "b", "c", "d", "e"}

	for _, val := range chatsslice {
		fmt.Println(val)
	}
	wg.Done()
}

//Next question: Which of the following is NOT a correct statement about the Golang runtime scheduler?


/*
package main

import "fmt"

func main() {
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2
    close(ch)
    for n := range ch {
        fmt.Println(n)
    }
}
Will print 1 in one line and 2 in the next line
*/
