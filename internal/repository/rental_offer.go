package repository

import (
	"context"
	"database/sql"
	"fmt"
	"real_estate_finder/real_estate_finder/internal/repository/types"
)

// TODO: сделать нормализацию 3NF
// на данный момент это не удалось, зря продрочился два дня

func (r *repository) InsertRentalOffers(rentalOffers []*types.RentalOffer) error {
	ctx := context.Background()
	validOffers := validateAndTrimOffers(rentalOffers)

	err := r.withTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		err := addRentalOffers(validOffers, tx)
		if err != nil {
			return fmt.Errorf("failed to insert rental offers: %v", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("repository.InsertRentalOffers: %v", err)
	}
	return nil
}

func validateAndTrimOffers(offers []*types.RentalOffer) []*types.RentalOffer {
	i := 0
	for _, offer := range offers {
		if isValid(offer) {
			offers[i] = offer
			i++
		}
	}

	return offers[:i]
}

func isValid(offer *types.RentalOffer) bool {
	hasCoords := offer.Latitude != -1 && offer.Longitude != -1
	hasAddr := offer.FullAddress != ""
	hasLink := offer.Link != ""
	hasPrice := offer.Price != 0

	return hasCoords && hasAddr && hasLink && hasPrice
}

// Разделяет предложения об аренде по уникальности координат
//func splitOffers(offers []*types.RentalOffer) (uniqueOffers []*types.RentalOffer, duplicateOffers []*types.RentalOffer) {
//	type coords struct {
//		lat float64
//		lng float64
//	}
//	uniqueCoordsMap := make(map[*coords]struct{})
//
//	uniqueOffers = make([]*types.RentalOffer, 0)
//	duplicateOffers = make([]*types.RentalOffer, 0)
//	for _, offer := range offers {
//		offerCoords := &coords{
//			lat: offer.Latitude,
//			lng: offer.Longitude,
//		}
//		if _, ok := uniqueCoordsMap[offerCoords]; !ok {
//			uniqueOffers = append(uniqueOffers, offer)
//			uniqueCoordsMap[offerCoords] = struct{}{}
//			continue
//		}
//		duplicateOffers = append(duplicateOffers, offer)
//	}
//
//	return
//}

func addRentalOffers(offers []*types.RentalOffer, tx *sql.Tx) error {
	cols := 6

	// todo: on conflict do nothing
	q := fmt.Sprintf(`
		INSERT INTO rental_offer (latitude, longitude, full_address, link, price, property_type) 
			VALUES %s`, genStmtValuesString(cols, len(offers), 1))

	vals := make([]interface{}, 0, len(offers)*cols)
	for _, offer := range offers {
		vals = append(
			vals,
			offer.Latitude,
			offer.Longitude,
			offer.FullAddress,
			offer.Link,
			offer.Price,
			offer.PropertyType,
		)
	}

	_, err := tx.Exec(q, vals...)

	return err
}

func (r *repository) GetRentalOffers() ([]*types.RentalOffer, error) {
	q := `SELECT id, latitude, longitude, full_address, link, price, property_type, created  from rental_offer`

	rentalOfers := make([]*types.RentalOffer, 0)
	rows, err := r.db.Query(q)
	defer rows.Close()

	for rows.Next() {
		offer := &types.RentalOffer{}
		err = rows.Scan(
			&offer.ID,
			&offer.Latitude,
			&offer.Longitude,
			&offer.FullAddress,
			&offer.Link,
			&offer.Price,
			&offer.PropertyType,
			&offer.Created,
		)
		if err != nil {
			return nil, err
		}
		rentalOfers = append(rentalOfers, offer)
	}

	return rentalOfers, nil
}
