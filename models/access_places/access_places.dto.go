package models

type AccessPlacesRequest struct {
	ID          int64  `json:"id,string"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
}

type AccessPlacesResponse struct {
	ID          int64  `json:"id,string"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (p AccessPlacesRequest) ParseRequest() AccessPlaces {
	return AccessPlaces{
		ID:          p.ID,
		Name:        p.Name,
		Icon:        p.Icon,
		Description: p.Description,
	}
}

func (p AccessPlacesResponse) ParseResponse() AccessPlaces {
	return AccessPlaces{
		ID:          p.ID,
		Name:        p.Name,
		Icon:        p.Icon,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}
