package main

import "fmt"

//nondetermenistic behaviour - this is the correct answer

func consume(ch chan int) {
	for {
		select {
		case num := <-ch:
			fmt.Printf("%d ", num)
			break
		}
	}
}

func main() {
	ch := make(chan int)
	go consume(ch)

	for i := 0; i < 5; i++ {
		ch <- i
	}

}
