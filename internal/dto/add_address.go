package dto

type AddAddressRequestBody struct {
	FullAddress string `json:"full_address"`
}

type AddAddressResponseBody struct {
	ID      int64 `json:"id"`
	Success bool  `json:"success,omitempty"`
}
