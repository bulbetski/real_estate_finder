package main

import (
	"fmt"
	"log"
	"net/http"
	"real_estate_finder/real_estate_finder/internal/server"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Initial commit")
	})

	srv := server.New()
	if err := srv.Start(":82"); err != nil {
		log.Fatalln(err.Error())
	}
}
