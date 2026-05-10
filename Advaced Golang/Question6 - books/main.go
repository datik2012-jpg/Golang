package main

import "fmt"


type Book struct {
	Title  string
	Author string
	Pages  int
}
//point to one book in the slice
func updatePages(book *Book, pages int) {
	// your code goes here
	//fmt.Print(&book) //print address of the book
	//fmt.Println(*book) //prin exactly a book
    book.Pages = pages
	//or (*book).Pages = 200   // explicit, rarely needed
	
}

//point to the slice of the books
func printBooks(my_books *[]Book) {
	fmt.Println("Play to print from address")
	for _, val := range *my_books {
		fmt.Println(val)
	}
}

func main() {

	/*
		Create 3 Book Structs with the following data:

		Book 1:
		Title: "The Great Gatsby"
		Author: "F. Scott Fitzgerald"
		Pages: 180

		Book 2
		Title: "To Kill a Mockingbird"
		Author: "Harper Lee"
		Pages: 281

		Book 3
		Title: "Pride and Prejudice"
		Author: "Jane Austen"
		Pages: 279
	*/

	// your code for creating struct objects goes here
	books := []Book {
		{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Pages: 180},
		{Title: "To Kill a Mockingbird", Author: "Harper Lee", Pages: 281},
		{Title: "Pride and Prejudice", Author:"Jane Austen", Pages: 279},
	}

	/*
		Update the information for Books as following:

		Book 1: Updates Page Count to 210
		Book 2: Updates Page Count to 250
		Book 3: Updates Page Count to 295

	*/

	// your code for function calls to updatePages goes here
	fmt.Println(books[0]) //before just print book
	updatePages(&books[0], 210)
	fmt.Println(books[0]) //after update the paeg for the book

	//update Book 2:
	fmt.Println(books[1]) //before update the page for the book just print the book
	updatePages(&books[1], 250)
	fmt.Println(books[1]) //print after the update

	//update Book 3:
	fmt.Println(books[2]) //before update the page for the book just print the book
	updatePages(&books[2], 295)
	fmt.Println(books[2]) //print after the update



	/*
		Print all the struct objects
		fmt.Println(book)
	*/

	// your code for printing objects goes here

	//loop to print all of them
	fmt.Println("Print all the books in the loop")
	for _, book := range books {
		fmt.Println(book)
	}

	//send adress of books to print somewhere
	printBooks(&books)
}
