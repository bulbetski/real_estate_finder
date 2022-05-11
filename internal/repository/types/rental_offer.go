package types

import "time"

type RentalOffer struct {
	ID           int64        `db:"id"`
	Latitude     float64      `db:"latitude"`
	Longitude    float64      `db:"longitude"`
	FullAddress  string       `db:"full_address"`
	Link         string       `db:"link"`
	Price        int64        `db:"price"`
	PropertyType PropertyType `db:"property_type"`
	Created      time.Time    `db:"created"`
}

type PropertyType string

const (
	PropertyType0     = "STUDIO"
	PropertyType1     = "1"
	PropertyType2     = "2"
	PropertyType3     = "3"
	PropertyType4plus = "PLUS_4"
)
