package models

type SubCategoriesRequest struct {
	ID         int64   `json:"id"`
	CategoryID int64   `json:"category_id"`
	Name       string  `json:"name"`
	Form       string  `json:"form"`
	Status     bool    `json:"status" `
	CreatedAt  *string `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
}

type SubCategoriesResponse struct {
	ID         int64   `json:"id"`
	CategoryID int64   `json:"category_id"`
	Name       string  `json:"name"`
	Form       string  `json:"form"`
	Status     bool    `json:"status"`
	CreatedAt  *string `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
}

func (p SubCategoriesRequest) ParseRequest() SubCategories {
	return SubCategories{
		ID:         p.ID,
		CategoryID: p.CategoryID,
		Name:       p.Name,
	}
}

func (p SubCategoriesResponse) ParseResponse() SubCategories {
	return SubCategories{
		ID:         p.ID,
		CategoryID: p.CategoryID,
		Name:       p.Name,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}

func (kr SubCategoriesRequest) TableName() string {
	return "sub_categories"
}

func (kr SubCategoriesResponse) TableName() string {
	return "sub_categories"
}
