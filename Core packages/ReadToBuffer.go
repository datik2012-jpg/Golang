package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("Let us catch up over a cup of coffee")
	// your code goes here
	buf := make([]byte, 5)
	for {
		n, err := reader.Read(buf)
		if n > 0 {
			//fmt.Println(string(buf[:n]))
			fmt.Println(buf[:n], err)
			//fmt.Println(buf[:n])
		}
		if err == io.EOF {
			break
			//fmt.Println("[] EOF")
			fmt.Println([]byte{}, err)
		}

	}

}
