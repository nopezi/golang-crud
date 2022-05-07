package models

type CategoryRequest struct {
	ID   int64  `json:"id,string"`
	Name string `json:"name"`
}

type CategoryResponse struct {
	ID        int64  `json:"id,string"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
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
