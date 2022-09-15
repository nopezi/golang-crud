package models

type ProductRequest struct {
	ID          int64   `json:"id"`
	KodeProduct string  `json:"kode_product"`
	Product     string  `json:"product"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"update_at"`
}

type ProductResponse struct {
	ID          int64   `json:"id"`
	KodeProduct string  `json:"kode_product"`
	Product     string  `json:"product"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"update_at"`
}

func (p ProductRequest) ParseRequest() Product {
	return Product{
		ID:          p.ID,
		KodeProduct: p.KodeProduct,
		Product:     p.Product,
	}
}

func (p ProductResponse) ParseRequest() Product {
	return Product{
		ID:          p.ID,
		KodeProduct: p.KodeProduct,
		Product:     p.Product,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (pr ProductRequest) TableName() string {
	return "product"
}

func (pr ProductResponse) TableName() string {
	return "product"
}
