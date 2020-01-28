package main

import (
	"fmt"
)

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s, nice to meet you!", name)
}

func Validate(name string) bool {
	return (len(name) > 0)
}
