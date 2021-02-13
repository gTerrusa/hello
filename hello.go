package main

import (
	"fmt"
)

// Constants should improve performance of your application
// as it saves you creating the "Hello, " string instance
// every time Hello is called.
const englishHelloPrefix = "Hello, "

// Hello accepts a string, an returns the string "Hello, " + name
// defaults to return "Hello, World" if empty string given
func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World"))
}
