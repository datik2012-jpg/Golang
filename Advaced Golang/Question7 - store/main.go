package main

import "fmt"

// Declare the Expense struct here
type Expense struct {
	name string
	amount int
	date string
}



// Implement the Total as a function to calculate the total amount spent
// your code goes here
func Total (exp []Expense) float64 {
	var total float64 = 0 
	for _, e := range exp {
		//fmt.Println(e.amount)
		total += float64(e.amount)
	}
	
	return total
}

func (e Expense) getName() string {
	return e.name
}

// Implement the getName method on the Expense struct here
// your code goes here

func main() {
	expenses := []Expense{
		Expense{"Grocery", 50.0, "2022-01-01"},
		Expense{"Gas", 30.0, "2022-01-02"},
		Expense{"Restaurant", 40.0, "2022-01-03"},
	}

	fmt.Println(Total(expenses))
	fmt.Println(expenses[0].getName())
}




//////////////implemented Total as e method and wrapper  - second solution (a third can be by type []Expense ) /////////


package main

import "fmt"

// Declare the Expense struct here
type Expense struct {
	name string
	amount int
	date string
}



// Implement the Total method to calculate the total amount spent
// your code goes here

//this is function - works but we have been asked for a METHOD of the structure
/*func Total (exp []Expense) float64 {
	var total float64 = 0 
	for _, e := range exp {
		//fmt.Println(e.amount)
		total += float64(e.amount)
	}
	
	return total
} */

//method for the structure !!!
func (e Expense) Total(exp []Expense) float64 {
    var total float64
    for _, item := range exp {
        total += float64(item.amount)
    }
    return total
}

//wrapper needed cause we need to use a method 
func Total(exp []Expense) float64 {
    return Expense{}.Total(exp)
}



func (e Expense) getName() string {
	return e.name
}

// Implement the getName method on the Expense struct here
// your code goes here

func main() {
	expenses := []Expense{
		Expense{"Grocery", 50.0, "2022-01-01"},
		Expense{"Gas", 30.0, "2022-01-02"},
		Expense{"Restaurant", 40.0, "2022-01-03"},
	}

	fmt.Println(Total(expenses))
	fmt.Println(expenses[0].getName())
}