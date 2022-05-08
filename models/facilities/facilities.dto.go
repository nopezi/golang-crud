package models

type FacilitiesRequest struct {
	ID          int64  `json:"id,string"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
}

type FacilitiesResponse struct {
	ID          int64  `json:"id,string"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (p FacilitiesRequest) ParseRequest() Facilities {
	return Facilities{
		ID:          p.ID,
		Name:        p.Name,
		Icon:        p.Icon,
		Description: p.Description,
	}
}

func (p FacilitiesResponse) ParseResponse() Facilities {
	return Facilities{
		ID:          p.ID,
		Name:        p.Name,
		Icon:        p.Icon,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}
