package main

import (
	"fmt"
	"time"
)

func main() {
	//create 2 channels
	ch1 := make(chan string)
	ch2 := make(chan string)

	//we call two gorutine methods and passing to them channel
	go goOne(ch1)
	go goTwo(ch2)

	select { //even if both gorutines executed in the same time only one of them will be cathched with select state.Good when we are waiting response from some Server and doesn't matter from which one
	case msg1 := <-ch1:
		fmt.Println(msg1) //we even can add break; which will break the execution when case catched
	case msg2 := <-ch2:
		fmt.Println(msg2)
	default: //makes select not blocking us
		fmt.Println("no messages at all")
	}

	time.Sleep(5 * time.Second)
}

func goOne(ch1 chan string) {

	ch1 <- "Channel - 1"
	close(ch1)
}

func goTwo(ch2 chan string) {

	ch2 <- "Channel - 2"
	close(ch2)
}
