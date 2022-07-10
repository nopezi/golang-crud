package models

type CategoryRequest struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Status    bool    `json:"status"`
	Icon      string  `json:"icon"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type CategoryResponse struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Status    bool    `json:"status"`
	Icon      string  `json:"icon"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

func (p CategoryRequest) ParseRequest() Categories {
	return Categories{
		ID:   p.ID,
		Name: p.Name,
	}
}

func (p CategoryResponse) ParseResponse() Categories {
	return Categories{
		ID:        p.ID,
		Name:      p.Name,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (cr CategoryRequest) TableName() string {
	return "categories"
}

func (cr CategoryResponse) TableName() string {
	return "categories"
}
