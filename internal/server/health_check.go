package server

import (
	"fmt"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("pong sent for %s\n", r.URL)
	w.Write([]byte("pong"))
}
