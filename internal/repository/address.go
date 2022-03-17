package repository

import (
	"real_estate_finder/real_estate_finder/internal/repository/types"
)

func (r *repository) FindAll() ([]*types.Address, error) {
	q := `SELECT id, latitude, longitude, full_address from address`

	result := make([]*types.Address, 0)
	rows, err := r.db.Query(q)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		addr := &types.Address{}
		err = rows.Scan(
			&addr.ID,
			&addr.Latitude,
			&addr.Longitude,
			&addr.FullAddress,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, addr)
	}

	return result, nil
}

func (r *repository) AddAddress(lat float64, lng float64, addr string) (id int64, err error) {
	q := `INSERT INTO address (latitude, longitude, full_address) VALUES ($1, $2, $3) RETURNING id`
	err = r.db.QueryRow(q, lat, lng, addr).Scan(&id)

	return
}
