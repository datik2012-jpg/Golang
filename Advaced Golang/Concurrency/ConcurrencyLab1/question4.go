package main

/* Select the correct behavior for the following program: */
//My answer programs will fail - what is in the line 18 ? will cause to the issue - deadlock cause we close the channle before
import (
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		ch <- 1
		wg.Done()
	}()
	wg.Wait()
	close(ch)
	<-ch
}
