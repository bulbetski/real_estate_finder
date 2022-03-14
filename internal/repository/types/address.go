package types

type Address struct {
	FormattedAddress string `db:"formatted_address"`
	Street           string `db:"street"`
	HouseNumber      string `db:"house_number"`
	Suburb           string `db:"suburb"`
	Postcode         string `db:"postcode"`
	State            string `db:"state"`
	StateCode        string `db:"state_code"`
	StateDistrict    string `db:"state_district"`
	County           string `db:"county"`
	Country          string `db:"country"`
	CountryCode      string `db:"country_code"`
	City             string `db:"city"`
}
