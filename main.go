package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	fmt.Println("Server is running on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
