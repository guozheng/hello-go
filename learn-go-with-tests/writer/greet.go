package writer

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string)  {
	fmt.Fprintf(writer, "Hi, %s", name)
}

func HttpRequestHandler(writer http.ResponseWriter, r *http.Request)  {
	Greet(writer, "Claire")
}

func main()  {
	http.ListenAndServe(":8080", http.HandlerFunc(HttpRequestHandler))
}
