package main

import (
	"fmt"
	"time"
)

func calculate(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println(i * i)
}

func main() {
	start := time.Now()
	for i := 1; i <= 100; i++ {
		fmt.Println("damo")
		go calculate(i)
	}
	elapsed := time.Since(start)
	time.Sleep(2 * time.Second) //wihout this line main will not wait for goritines methods and will exist , so will not see the numbers printed
	fmt.Println("Started from ", elapsed)
}
