package access_places

type AccessPlacesRequest struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Icon        string  `json:"icon"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

type AccessPlacesRequests []map[string]interface{}

type AccessPlacesResponse struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Icon        string  `json:"icon"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

// func (p AccessPlacesRequest) ParseRequest() AccessPlaces {
// 	return AccessPlaces{
// 		ID:          p.ID,
// 		Name:        p.Name,
// 		Icon:        p.Icon,
// 		Description: p.Description,
// 	}
// }

// func (p AccessPlacesResponse) ParseResponse() AccessPlaces {
// 	return AccessPlaces{
// 		ID:          p.ID,
// 		Name:        p.Name,
// 		Icon:        p.Icon,
// 		Description: p.Description,
// 		CreatedAt:   p.CreatedAt,
// 		UpdatedAt:   p.UpdatedAt,
// 	}
// }

func (kr AccessPlacesRequest) TableName() string {
	return "access_places"
}

func (kr AccessPlacesResponse) TableName() string {
	return "access_places"
}
