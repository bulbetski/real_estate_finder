package webscraper

import (
	"real_estate_finder/real_estate_finder/internal/dto"
	"real_estate_finder/real_estate_finder/internal/repository/types"
)

func toDtoRentalOffers(offers []*types.RentalOffer) []*dto.RentalOffer {
	dtoOffer := make([]*dto.RentalOffer, 0, len(offers))

	for _, offer := range offers {
		repoOffer := &dto.RentalOffer{
			ID:          offer.ID,
			Latitude:    offer.Latitude,
			Longitude:   offer.Longitude,
			FullAddress: offer.FullAddress,
			Link:        offer.Link,
			Created:     offer.Created,
		}
		dtoOffer = append(dtoOffer, repoOffer)
	}

	return dtoOffer
}
func toRepoRentalOffers(offers []*dto.RentalOffer) []*types.RentalOffer {
	repoOffers := make([]*types.RentalOffer, 0, len(offers))

	for _, offer := range offers {
		repoOffer := &types.RentalOffer{
			Latitude:    offer.Latitude,
			Longitude:   offer.Longitude,
			FullAddress: offer.FullAddress,
			Link:        offer.Link,
		}
		repoOffers = append(repoOffers, repoOffer)
	}

	return repoOffers
}
