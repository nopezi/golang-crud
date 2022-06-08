package models

type DistrictsRequest struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	ProvinceID int64  `json:"province_id"`
	CityID     int64  `json:"city_id"`
}

type DistrictsResponse struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	ProvinceID int64  `json:"province_id"`
	CityID     int64  `json:"city_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func (p DistrictsRequest) ParseRequest() Districts {
	return Districts{
		ID:         p.ID,
		Name:       p.Name,
		ProvinceID: p.ProvinceID,
		CityID:     p.CityID,
	}
}

func (p DistrictsResponse) ParseResponse() Districts {
	return Districts{
		ID:         p.ID,
		Name:       p.Name,
		ProvinceID: p.ProvinceID,
		CityID:     p.CityID,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}
