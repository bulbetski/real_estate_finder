package geocoder

import (
	"fmt"
	"real_estate_finder/real_estate_finder/internal/dto"
	"real_estate_finder/real_estate_finder/internal/repository/types"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/yandex"
)

type repositoryInterface interface {
	FindAll() ([]*types.Address, error)
	AddAddress(lat float64, lng float64, addr string) (int64, error)
	//GetAddress(id int64) (*repoTypes.Address, error)
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

func (g *geocoder) FindAll() ([]*dto.FindAllResponseBody, error) {
	addresses, err := g.repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("error getting addresses from repo: %v", err)
	}

	return addressesToDTO(addresses), nil
}

func (g *geocoder) AddAddress(addr string) (*dto.AddAddressResponseBody, error) {
	resp := &dto.AddAddressResponseBody{}
	loc, err := g.gc.Geocode(addr)
	if err != nil {
		return resp, fmt.Errorf("error occured while trying to geocode addr: %v", err)
	}

	id, err := g.repository.AddAddress(loc.Lat, loc.Lng, addr)
	if err != nil {
		return resp, fmt.Errorf("error occured while trying to save addr to db: %v", err)
	}
	resp.ID = id
	resp.Success = true

	return resp, nil
}

//func (g *geocoder) AddAddress(lat float64, lng float64) (*dto.AddAddressResponseBody, error) {
//	// todo: 1) pretty error formatting
//	resp := &dto.AddAddressResponseBody{}
//	addr, err := g.gc.ReverseGeocode(lat, lng)
//	if err != nil {
//		return nil, fmt.Errorf("error while trying to reverse geocode: %v", err)
//	}
//
//	err = g.repository.AddAddress(fromGeoAddr(addr))
//	if err != nil {
//		return nil, err
//	}
//	resp.Success = true
//
//	return resp, nil
//}
//
//func (g *geocoder) GetAddressByID(id int64) (*dto.GetAddressByIDResponseBody, error) {
//	addr, err := g.repository.GetAddress(id)
//	if err != nil {
//		return nil, fmt.Errorf("error getting addr from repo: %v", err)
//	}
//
//	return toDtoAddr(addr), nil
//}
