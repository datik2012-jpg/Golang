package main

import (
	"fmt"
	"time"
)

/* Dani   Non-deterministic!!  this example shows go-rutine doesn't have parent or child .go rutines executed independent
In Process
In Start
or sometimes
Dani
In Start
In Process

so Non-deterministic in programming means that a program or algorithm can produce different outputs or exhibit different behaviors, even when given the same input.
*/

func main() {
	fmt.Println("Dani")
	go start()
	time.Sleep(1 * time.Second)
}

func start() {
	go process()
	fmt.Println("In Start")
	time.Sleep(1 * time.Second)
}

func process() {
	fmt.Println("In Process")
}
