package main

import (
	"fmt"
)

// Hello greets the world
func Hello(name string) string {
	if name == "" {
		name = "nobody"
	}
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
