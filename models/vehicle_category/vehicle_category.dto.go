package models

type VehicleCategoryRequest struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Status    bool    `json:"status"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type VehicleCategoryResponse struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Status    bool    `json:"status"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

func (p VehicleCategoryRequest) ParseRequest() VehicleCategory {
	return VehicleCategory{
		ID:   p.ID,
		Name: p.Name,
	}
}

func (p VehicleCategoryResponse) ParseResponse() VehicleCategory {
	return VehicleCategory{
		ID:        p.ID,
		Name:      p.Name,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (cr VehicleCategoryRequest) TableName() string {
	return "vehicle_category"
}

func (cr VehicleCategoryResponse) TableName() string {
	return "vehicle_category"
}
