package main

import "fmt"

func discountedPrice(product string, price float64) float64 {
	
	switch product {
		case "apples":
			discount := price * 0.1
			after_disc := price - discount
			//fmt.Println(after_disc)
			return after_disc
		case "oranges":
			return price
		
		case "bananas":
			discount := price * 0.2
			after_disc := price - discount
			//fmt.Println(after_disc)
			return after_disc		
		default:
			return price
	}	
}

func main() {
	fmt.Println(discountedPrice("apples", 100))
	fmt.Println(discountedPrice("orange", 100))
	fmt.Println(discountedPrice("bananas", 100))
	fmt.Println(discountedPrice("bananas", 100))
	fmt.Println(discountedPrice("oranges", 100))
}