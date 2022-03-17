package dto

type FindAllResponseBody struct {
	ID          int64   `json:"id"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	FullAddress string  `json:"full_address"`
}
