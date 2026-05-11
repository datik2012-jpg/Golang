package main

import (
	"fmt"
	"sync"
)

func main() {
	var c chan int
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		c <- 1
		wg.Done()
	}()

	go func() {
		val := <-c
		fmt.Println(val)
		wg.Done()
	}()
	wg.Wait()

}

//Why does the following program fail?
//Sending and recieving a nil channel is blocking // THIS IS A NIL CHANNEL line 9
//should be c := make(chan int)

//What happens if you read from an unbuffered, closed channel?
//it will return ZERO
