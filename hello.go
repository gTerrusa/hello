package main

import (
	"fmt"
)

// Hello accepts a string, an returns the string "Hello, " + name
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
