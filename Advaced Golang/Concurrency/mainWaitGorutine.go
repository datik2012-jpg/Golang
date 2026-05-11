package main

//i will add wait for gorutine completed for more controll

import (
	"fmt"
	"sync"
	"time"
)

func calculate(i int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println(i * i)
	wg.Done() //notification to do -1, when go rutine finished the work
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup // this is wait group which we need for the task - to controll the finishing the gorutine
	//we will create wait group of 100 (cause we have 100 go rutines ) and it will count-1 when go rutine finished to work
	wg.Add(100)
	for i := 1; i <= 100; i++ {
		fmt.Println("damo")
		go calculate(i, &wg)
	}
	elapsed := time.Since(start)
	wg.Wait()                   // will proceed - only when all 100 go rutines completed
	time.Sleep(2 * time.Second) //wihout this line main will not wait for goritines methods and will exist , so will not see the numbers printed
	fmt.Println("Started from ", elapsed)
}
