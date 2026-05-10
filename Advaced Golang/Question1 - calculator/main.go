package main

import "fmt"

func calculate(a int, b int) []float64 {
	sum := float64(a + b)
	diff := float64(a - b)
	prod := float64(a * b)
	quot := float64(a) / float64(b)
    
	results := []float64{}
	results = append(results, sum, diff, prod, quot)
	//fmt.Print(results)
    return results
}

func main() {
	fmt.Println(calculate(20, 10))
	fmt.Println(calculate(700, 70))
}