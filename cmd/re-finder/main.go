package main

import (
	"database/sql"
	"log"
	"os"
	"real_estate_finder/real_estate_finder/internal/geocoder"
	"real_estate_finder/real_estate_finder/internal/repository"
	"real_estate_finder/real_estate_finder/internal/server"

	_ "github.com/lib/pq"
)

var databaseURL = "host=localhost port=5432 user=zproselkov dbname=re_finder sslmode=disable"

func main() {

	// Connects to db
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
	repo := repository.New(db)

	token := os.Getenv("YANDEX_API_TOKEN")
	gc := geocoder.New(token, repo)

	srv := server.New(repo, gc)
	srv.LoadHTML("templates/*")

	if err := srv.Start(":82"); err != nil {
		log.Fatalln(err.Error())
	}
}
