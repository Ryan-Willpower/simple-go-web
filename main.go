package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", homeHandler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(":80", err)
	}
}
