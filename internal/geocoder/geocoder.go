package geocoder

//func (g *geocoder) FindAll() ([]*dto.FindAllResponseBody, error) {
//	addresses, err := g.repository.FindAll()
//	if err != nil {
//		return nil, fmt.Errorf("error getting addresses from repo: %v", err)
//	}
//
//	return addressesToDTO(addresses), nil
//}
//
//func (g *geocoder) AddAddress(addr string) (*dto.AddAddressResponseBody, error) {
//	resp := &dto.AddAddressResponseBody{}
//	loc, err := g.gc.Geocode(addr)
//	if err != nil {
//		return resp, fmt.Errorf("error occured while trying to geocode addr: %v", err)
//	}
//
//	id, err := g.repository.AddAddress(loc.Lat, loc.Lng, addr)
//	if err != nil {
//		return resp, fmt.Errorf("error occured while trying to save addr to db: %v", err)
//	}
//	resp.ID = id
//	resp.Success = true
//
//	return resp, nil
//}
