package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello mbak salam dari baim!")
	})

	http.ListenAndServe(":8001", nil)
}