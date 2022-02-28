package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Initial commit")
	})
	err := http.ListenAndServe(":82", nil)
	if err != nil {
		log.Fatal(err)
	}
}
