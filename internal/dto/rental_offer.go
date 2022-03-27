package dto

import "time"

type RentalOffer struct {
	ID          int64     `json:"id"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	FullAddress string    `json:"full_address"`
	Link        string    `json:"link"`
	Created     time.Time `json:"created"`
}
