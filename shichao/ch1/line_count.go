package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	l := ""
	for input.Scan() {
		l = input.Text()
		if l == "Z" {
			break
		}
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%s count=%d\n", line, n)
		}
	}
}