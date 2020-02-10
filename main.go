package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/xiote/stringutil"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("What is your name? ")
	for scanner.Scan() {
		var name string = scanner.Text()

		var isempty bool = strutil.IsEmptyS(name)
		if isempty == true {
			fmt.Println("Type valid name")
			fmt.Println()
			fmt.Printf("What is your name? ")
			continue
		}

		var message = Hello(name)
		fmt.Println(message)
		fmt.Println()
		fmt.Printf("What is your name? ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
