package dto

type GetAddressByIDRequestBody struct {
	ID int64 `json:"id"`
}

type GetAddressByIDResponseBody struct {
	FormattedAddress string `json:"formatted_address"`
	Street           string `json:"street"`
	HouseNumber      string `json:"house_number"`
	Suburb           string `json:"suburb"`
	Postcode         string `json:"postcode"`
	State            string `json:"state"`
	StateCode        string `json:"state_code"`
	StateDistrict    string `json:"state_district"`
	County           string `json:"county"`
	Country          string `json:"country"`
	CountryCode      string `json:"country_code"`
	City             string `json:"city"`
}