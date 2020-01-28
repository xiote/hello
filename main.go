package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("What is your name? ")
	for scanner.Scan() {
		var name string = scanner.Text()
		var message = Hello(name)
		fmt.Println(message)
		fmt.Println()
		fmt.Printf("What is your name? ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
