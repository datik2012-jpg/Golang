package main

import (
	"fmt"
	"strings"
)

func WordCount(s string, word string) int {
	// your code goes here
	var counter int = 0
	for _, w := range strings.Fields(s) {
		fmt.Println(w)
		//convert to lower all words to disable case sensitive
		//w := strings.ToLower(w)
		if strings.Contains(w, word) {
			fmt.Println("word appears")
			counter += 1
		}

	}

	return counter
}

func main() {
	count := WordCount("hello, Hello how have you been in helloworld", "hello")
	fmt.Println(count)
}
