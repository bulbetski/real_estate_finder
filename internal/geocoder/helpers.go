package geocoder

import (
	"real_estate_finder/real_estate_finder/internal/dto"
	"real_estate_finder/real_estate_finder/internal/repository/types"
)

func addressesToDTO(addresses []*types.Address) []*dto.FindAllResponseBody {
	if addresses == nil {
		return nil
	}
	result := make([]*dto.FindAllResponseBody, len(addresses))
	for i, addr := range addresses {
		result[i] = addressToDTO(addr)
	}
	return result
}

func addressToDTO(addr *types.Address) *dto.FindAllResponseBody {
	return &dto.FindAllResponseBody{
		ID:          addr.ID,
		Latitude:    addr.Latitude,
		Longitude:   addr.Longitude,
		FullAddress: addr.FullAddress,
	}
}
