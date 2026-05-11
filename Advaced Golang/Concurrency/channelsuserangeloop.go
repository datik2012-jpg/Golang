package main

import "fmt"

func main() {

	my_range := []int{1, 2, 3, 4, 5}

	for key, val := range my_range {
		fmt.Println(key, val)

	}

	//how iterate values from the channel
	//if we run this code and channel will not be closed before - we get an error deadlock
	//so close the channel before you loop : close(ch) - put this close where needed int he code
	//it is ok to close the channel which still have values insde - they can be read
	//for val := range ch {
	//fmt.Println("Recieved from channel:", val)
	//}

}
