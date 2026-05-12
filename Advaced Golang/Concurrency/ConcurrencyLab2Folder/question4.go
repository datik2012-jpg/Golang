//Complete the code here, for the program to work successfully and give the following response with no error:

/*
sent job 1
sent job 2
sent job 3
sent all jobs
received job 1
received job 2
received job 3
received all jobs

*/

package main

import "fmt"

func main() {
	jobs := make(chan int, 5) //create buffered channel size 5
	done := make(chan bool)   //create not buffered channel

	go func() { //Block forever waiting on <-jobs

		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}

	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs) //needed to add this line cause The else branch only runs when: line 26 has more == false → which happens only when the channel is closed.

	// your code goes here
	fmt.Println("sent all jobs")

	<-done //Wait here until someone sends a value into the done channel.
}
