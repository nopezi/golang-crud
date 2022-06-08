package models

type ProvincesRequest struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ProvincesResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (p ProvincesRequest) ParseRequest() Provinces {
	return Provinces{
		ID:   p.ID,
		Name: p.Name,
	}
}

func (p ProvincesResponse) ParseResponse() Provinces {
	return Provinces{
		ID:        p.ID,
		Name:      p.Name,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
