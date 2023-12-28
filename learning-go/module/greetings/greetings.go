package greetings

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello, %v. Welcome", name)
}

func Sum(values ...int) int {
	total := 0
	for _, i := range values {
		total += i
	}
	return total
}