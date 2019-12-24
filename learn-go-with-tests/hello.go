package main

import (
	"fmt"
	"github.com/guozheng/hello-go/learn-go-with-tests/writer"
	"net/http"
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
	http.ListenAndServe(":8080", http.HandlerFunc(writer.HttpRequestHandler))
}
