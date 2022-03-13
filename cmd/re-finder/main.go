package main

import (
	"log"
	"os"
	"real_estate_finder/real_estate_finder/internal/server"
)

func main() {
	token := os.Getenv("YANDEX_API_TOKEN")
	srv := server.New(token)

	srv.LoadHTML("templates/*")

	if err := srv.Start(":82"); err != nil {
		log.Fatalln(err.Error())
	}
	//geocoder := yandex.Geocoder(token)
	//loc, err := geocoder.Geocode("Москва, СВАО, р-н Марьина роща, м. Марьина роща, улица 4-я Марьиной рощи, 12К1")
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//addr, err := geocoder.ReverseGeocode(loc.Lat, loc.Lng)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//fmt.Println(addr)
}
