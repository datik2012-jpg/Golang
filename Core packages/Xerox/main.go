package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// your code goes here
	var count int = 0

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Good")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Print("hello")
		}
		count += n
	}
	fmt.Println(count)
}
