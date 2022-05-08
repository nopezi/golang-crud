package models

type AddressesRequest struct {
	ID         int64  `json:"id,string"`
	Name       string `json:"name"`
	ProvinceID int64  `json:"province_id",string`
	CityID     int64  `json:"city_id",string`
	DistrictID int64  `json:"district_id",string`
	Address    string `json:"address"`
	Longitude  string `json:"longitude"`
	Langitude  string `json:"langitude"`
}

type AddressesResponse struct {
	ID         int64  `json:"id,string"`
	Name       string `json:"name"`
	ProvinceID int64  `json:"province_id",string`
	CityID     int64  `json:"city_id",string`
	DistrictID int64  `json:"district_id",string`
	Address    string `json:"address"`
	Longitude  string `json:"longitude"`
	Langitude  string `json:"langitude"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func (p AddressesRequest) ParseRequest() Addresses {
	return Addresses{
		ID:         p.ID,
		Name:       p.Name,
		ProvinceID: p.ProvinceID,
		CityID:     p.CityID,
		DistrictID: p.DistrictID,
		Address:    p.Address,
		Longitude:  p.Longitude,
		Langitude:  p.Langitude,
	}
}

func (p AddressesResponse) ParseResponse() Addresses {
	return Addresses{
		ID:         p.ID,
		Name:       p.Name,
		ProvinceID: p.ProvinceID,
		CityID:     p.CityID,
		DistrictID: p.DistrictID,
		Address:    p.Address,
		Longitude:  p.Longitude,
		Langitude:  p.Langitude,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}
