package dto

type AddAddressRequestBody struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type AddAddressResponseBody struct {
	Success bool `json:"success,omitempty"`
}
