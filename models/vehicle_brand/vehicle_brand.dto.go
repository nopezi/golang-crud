package models

type VehicleBrandRequest struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Status    bool    `json:"status"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type VehicleBrandResponse struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Status    bool    `json:"status"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

func (p VehicleBrandRequest) ParseRequest() VehicleBrand {
	return VehicleBrand{
		ID:   p.ID,
		Name: p.Name,
	}
}

func (p VehicleBrandResponse) ParseResponse() VehicleBrand {
	return VehicleBrand{
		ID:        p.ID,
		Name:      p.Name,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (cr VehicleBrandRequest) TableName() string {
	return "vehicle_brand"
}

func (cr VehicleBrandResponse) TableName() string {
	return "vehicle_brand"
}
