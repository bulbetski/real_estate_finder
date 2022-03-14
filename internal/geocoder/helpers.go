package geocoder

import (
	"real_estate_finder/real_estate_finder/internal/dto"
	"real_estate_finder/real_estate_finder/internal/repository/types"

	"github.com/codingsince1985/geo-golang"
)

func fromGeoLoc(loc *geo.Location) *types.Location {
	if loc != nil {
		return &types.Location{
			Lng: loc.Lng,
			Lat: loc.Lat,
		}
	}

	return &types.Location{}
}

func fromGeoAddr(addr *geo.Address) *types.Address {
	if addr == nil {
		return &types.Address{}
	}
	return &types.Address{
		FormattedAddress: addr.FormattedAddress,
		Street:           addr.Street,
		HouseNumber:      addr.HouseNumber,
		Suburb:           addr.Suburb,
		Postcode:         addr.Postcode,
		State:            addr.State,
		StateCode:        addr.StateCode,
		StateDistrict:    addr.StateDistrict,
		County:           addr.County,
		Country:          addr.Country,
		CountryCode:      addr.CountryCode,
		City:             addr.City,
	}
}

func toDtoAddr(addr *types.Address) *dto.GetAddressByIDResponseBody {
	if addr == nil {
		return &dto.GetAddressByIDResponseBody{}
	}
	return &dto.GetAddressByIDResponseBody{
		FormattedAddress: addr.FormattedAddress,
		Street:           addr.Street,
		HouseNumber:      addr.HouseNumber,
		Suburb:           addr.Suburb,
		Postcode:         addr.Postcode,
		State:            addr.State,
		StateCode:        addr.StateCode,
		StateDistrict:    addr.StateDistrict,
		County:           addr.County,
		Country:          addr.Country,
		CountryCode:      addr.CountryCode,
		City:             addr.City,
	}
}
