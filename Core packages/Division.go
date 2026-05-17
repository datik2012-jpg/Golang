package main

import (
	"fmt"
)

func divideNumber(m int, n int) (int, error) {
	// your code goes here
	if n == 0 {
		//log.Fatal("Cannot divide by zero", n) in case we want to print and exit the program Prints: 2026/05/17 09:40:35 Cannot divide by zero0 exit status 1
		return 0, fmt.Errorf("processing failed: %w", "Cannot divide by zero")
	}

	if n < 0 {
		return 0, fmt.Errorf("processing failed: %w", "Division is not supported for negative numbers")
	}

	return m / n, nil
}

func printResult(a int, b int) {
	res, err := divideNumber(a, b)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(res)
	}
}
func main() {
	a, b := 20, 10
	printResult(a, b)
	a, b = 20, -1
	printResult(a, b)
	a, b = 20, 0
	printResult(a, b)
}

/*

Expectation output is:

2
Error:  Division is not supported for negative numbers
Error:  Cannot divide by zero

*/
