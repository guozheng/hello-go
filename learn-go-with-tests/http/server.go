package http

import (
	"fmt"
	"net/http"
)

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}