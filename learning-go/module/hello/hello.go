package main

import (
	"fmt"
	"module/greetings"
)

func main() {
	message := greetings.Hello("Alex")
	fmt.Println(message)
	values := []int{1, 2, 3, 4, 5}
	fmt.Println(greetings.Sum(values...))
}