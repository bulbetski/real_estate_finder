package geocoder

import (
	"fmt"
	"real_estate_finder/real_estate_finder/internal/dto"
	repoTypes "real_estate_finder/real_estate_finder/internal/repository/types"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/yandex"
)

type repositoryInterface interface {
	AddAddress(addr *repoTypes.Address) error
	GetAddress(id int64) (*repoTypes.Address, error)
}

type geocoder struct {
	gc         geo.Geocoder
	repository repositoryInterface
}

func New(token string, repository repositoryInterface) *geocoder {
	g := yandex.Geocoder(token)
	return &geocoder{
		gc:         g,
		repository: repository,
	}
}

func (g *geocoder) AddAddress(req *dto.AddAddressRequestBody) (*dto.AddAddressResponseBody, error) {
	// todo: 1) pretty error formatting
	// 		 2) figure out where to use pointers
	if req == nil {
		return nil, fmt.Errorf("nil req")
	}
	resp := &dto.AddAddressResponseBody{}

	addr, err := g.gc.ReverseGeocode(req.Latitude, req.Longitude)
	if err != nil {
		return nil, fmt.Errorf("error while trying to reverse geocode: %v", err)
	}

	err = g.repository.AddAddress(fromGeoAddr(addr))
	if err != nil {
		return nil, err
	}
	resp.Success = true

	return resp, nil
}

func (g *geocoder) GetAddressByID(req *dto.GetAddressByIDRequestBody) (*dto.GetAddressByIDResponseBody, error) {
	id := req.ID

	addr, err := g.repository.GetAddress(id)
	if err != nil {
		return nil, fmt.Errorf("error getting addr from repo: %v", err)
	}

	return toDtoAddr(addr), nil
}
