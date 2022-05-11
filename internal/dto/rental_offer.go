package dto

import "time"

type PropertyType struct {
	PropertyTypes []string `json:"property_types"`
}

type RentalOffer struct {
	ID           int64     `json:"id"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	FullAddress  string    `json:"full_address"`
	Link         string    `json:"link"`
	Price        int64     `json:"price"`
	PropertyType string    `json:"property_type"`
	Created      time.Time `json:"created"`
}
