/*
In this question, we are working on a project that involves processing a large amount of data in parallel using goroutines. We want to use channels to communicate between the goroutines and ensure that all goroutines have finished processing before the program exits.


Complete the code, so that the program runs successfully without any errors.

Expected Output:

sent job 1
received job 1
received job 2
sent job 2
received result 2
received result 4
received job 3
received result 6
sent job 3

*/

package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(3)
	jobs := make(chan int)
	results := make(chan int)

	//This runs as soon as the scheduler starts it, which is usually almost immediately, but asynchronously.
	go func() { //func will wait the input fron the jobs channel and reads from jobs
		for j := range jobs { //the loop finished only when channel is closed so need to close the channels in the line: 52 ?
			fmt.Println("received job", j)
			result := j * 2
			results <- result
		}
		// your code goes here
		close(results) //we can sloes this thannel - no need anymore
		wg.Done()      //we do -1 which is flagging we are done with the loop channel reading

	}()

	go func() { //func which sends to the channell .channel is unbuffered : if the channel is unbuffered, the sender must wait until a receiver reads the value.
		for j := 1; j <= 3; j++ {
			jobs <- j
			fmt.Println("sent job", j)
		}
		// your code goes here
		close(jobs)
		wg.Done() //we do -1 cause we done with writing to the channel

	}()

	go func() {
		for r := range results {
			fmt.Println("received result", r)
		}
		// your code goes here
		wg.Done()
		//close(results) //closing the channel in this place is not good cause 1,2,3 we will get to this step which is losed
	}()

	wg.Wait()
}
