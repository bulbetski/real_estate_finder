package types

import "time"

type RentalOffer struct {
	ID          int64     `db:"id"`
	Latitude    float64   `db:"latitude"`
	Longitude   float64   `db:"longitude"`
	FullAddress string    `db:"full_address"`
	Link        string    `db:"link"`
	Created     time.Time `db:"created"`
}
