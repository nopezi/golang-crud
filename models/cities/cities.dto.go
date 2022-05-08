package models

type CitiesRequest struct {
	ID         int64  `json:"id,string"`
	Name       string `json:"name"`
	ProvinceID int64  `json:"province_id",string`
}

type CitiesResponse struct {
	ID         int64  `json:"id,string"`
	Name       string `json:"name"`
	ProvinceID int64  `json:"province_id",string`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func (p CitiesRequest) ParseRequest() Cities {
	return Cities{
		ID:         p.ID,
		Name:       p.Name,
		ProvinceID: p.ProvinceID,
	}
}

func (p CitiesResponse) ParseResponse() Cities {
	return Cities{
		ID:         p.ID,
		Name:       p.Name,
		ProvinceID: p.ProvinceID,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}
