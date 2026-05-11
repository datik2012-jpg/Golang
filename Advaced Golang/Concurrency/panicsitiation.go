package main

import "fmt"

func main() {
	//example of close the channel and panic
	ch := make(chan int, 10) //create a channel size 10 of int
	ch <- 10
	ch <- 20
	val, ok := <-ch //we reading from the channel .ok - to check if channel is not empty and closed
	fmt.Println(val, ok)
	close(ch) //we close the channel
	//try to read again - will work cause channel close but still have 20 inside
	val, ok = <-ch
	fmt.Println(val, ok) //will print 20 and OK

	//try to read again from close channel
	val, ok = <-ch
	fmt.Println(val, ok) //will return 0 and FALSE

	//PANIC SITUATIONS examples:
	//panic situtation when we trying to send a value to the channel which is CLOSE message: panic: send on closed channel
	//ch <- 30
	//another panic situtation is when CLOSE already closes channel
	//close(ch) //panic: close of closed channel
}
