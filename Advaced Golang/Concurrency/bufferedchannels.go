package main

//when channel buffer get full - it will be blocked!!!

import (
	"fmt"
	"sync"
)

func main() {
	//fmt.Println("Dani")
	ch := make(chan int, 3) //create buffered channel buffer with size 3 - which holds integer values
	var wg sync.WaitGroup   //define a wait group with the size 2 . It will wait for only 2 gorutines to complete.It will block the main function do not close UNTIL these 2 gorutines completed to work
	wg.Add(2)
	//need to add Wait in the end of the main method - which will wait for these two gorutines to be completed and not finish the main before they done

	//create the go rutine which will write to the channel
	//in that go rutine when it complite write to the channel need call the READ go rutine from the channel
	go sell(ch, &wg) //we send to the method two parameters CHANNEL and WAIT GROUP address
	//ch <- 100        //write integer: 100 to the chanel

	wg.Wait() //wait for the gorutines completed and do not close the main finction
}

func sell(ch chan int, s *sync.WaitGroup) { //method which send the VALUES to the CHANNEL
	//we know the size of the channel is 3 so we can send 3 values to the channel
	ch <- 100 //write integer: 100 to the chanel
	ch <- 200
	ch <- 300
	go buy(ch, s) //by executed in another gorutine
	fmt.Println("we send 3 data to the channel")
	s.Done() //we do -1 for the wait group
}

func buy(ch chan int, s *sync.WaitGroup) { //method  which READ the VALUES from the CHANNEL
	fmt.Println("Waiting for a data from the channel")
	fmt.Println("Recieved data: ", <-ch)
	s.Done() //we do -1 for a wait group
}

//can be correct version which read 3 values + close the channel. Even when channel is close - it will resieve try until it has at least one value inside.When channels is epty it will recieve - false: val, ok := <-ch
/*
package main

import (
    "fmt"
    "sync"
)

func main() {
    ch := make(chan int, 3)
    var wg sync.WaitGroup

    wg.Add(2)

    go sell(ch, &wg)
    go buy(ch, &wg)

    wg.Wait()
}

func sell(ch chan int, s *sync.WaitGroup) {
    ch <- 100
    ch <- 200
    ch <- 300
    close(ch) // important!
    fmt.Println("Sent 3 values")
    s.Done()
}

func buy(ch chan int, s *sync.WaitGroup) {
    for v := range ch {
        fmt.Println("Received:", v)
    }
    s.Done()
}
*/
