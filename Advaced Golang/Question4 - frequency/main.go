package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string) map[string]int {
	// TODO: implement this function
	dict := make(map[string]int)

	words := strings.Fields(text)

	for _, w := range words {
		dict[w]++
	}
    
	//just for fun just loop over all words  -  in string
	for i, w := range words {
		//fmt.Println(w)
		fmt.Println(i, w)
	}
    //just for fun - loop over dictionarys
	for key, value := range dict{
		fmt.Println(key, value)
	}

	return dict

}

func main() {
	text := "The quick brown fox jumps over the lazy dog"
	fmt.Println(wordFrequency(text))

}