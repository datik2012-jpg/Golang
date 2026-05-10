package main

import "fmt"

// Declare variable activeUserCount
// your code goes here
var activeUserCount int = 0


func entry() {
	// Hint: you can use the "++" operator to increment a variable by 1
	// your code goes here
	activeUserCount += 1
	//fmt.Printf("%v", activeUserCount)
}

func exit() {
	// Hint: you can use the "--" operator to decrement a variable by 1
	// your code goes here
	activeUserCount -= 1
}

func main() {
	entry()
	entry()
	exit()
	exit()
	entry()
	entry()
	fmt.Println(activeUserCount)
}