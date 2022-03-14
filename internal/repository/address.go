package repository

import (
	"fmt"
	"real_estate_finder/real_estate_finder/internal/repository/types"
)

func (r *repository) GetAddress(id int64) (*types.Address, error) {
	// TODO: remove asterisk
	q := `SELECT * FROM "address" WHERE id = $1`

	var addr types.Address
	err := r.db.QueryRow(q, id).Scan(
		&id,
		&addr.FormattedAddress,
		&addr.Street,
		&addr.HouseNumber,
		&addr.Suburb,
		&addr.Postcode,
		&addr.State,
		&addr.StateCode,
		&addr.StateDistrict,
		&addr.County,
		&addr.Country,
		&addr.CountryCode,
		&addr.City,
	)
	if err != nil {
		return nil, fmt.Errorf("err getting addr from db: %v", err)
	}

	return &addr, nil
}

func (r *repository) AddAddress(addr *types.Address) error {
	cols := 12
	q := fmt.Sprintf(`INSERT INTO address
								(formatted_address,
								 street,
								 house_number,
								 suburb,
								 postcode,
								 state,
								 state_code,
								 state_district,
								 county,
								 country,
								 country_code,
								 city) VALUES %v RETURNING id`, getValuesString(cols, 1, 1))
	_, err := r.db.Exec(
		q,
		addr.FormattedAddress,
		addr.Street,
		addr.HouseNumber,
		addr.Suburb,
		addr.Postcode,
		addr.State,
		addr.StateCode,
		addr.StateDistrict,
		addr.County,
		addr.Country,
		addr.CountryCode,
		addr.City,
	)
	if err != nil {
		return fmt.Errorf("error adding address to db: %v", err)
	}
	return nil
}
