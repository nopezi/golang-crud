package models

type VehicleCapacityRequest struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Icon        string  `json:"icon"`
	Description string  `json:"description"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

type VehicleCapacityResponse struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Icon        string  `json:"icon"`
	Description string  `json:"description"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

func (p VehicleCapacityRequest) ParseRequest() VehicleCapacity {
	return VehicleCapacity{
		ID:   p.ID,
		Name: p.Name,
	}
}

func (p VehicleCapacityResponse) ParseResponse() VehicleCapacity {
	return VehicleCapacity{
		ID:        p.ID,
		Name:      p.Name,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (cr VehicleCapacityRequest) TableName() string {
	return "vehicle_capacity"
}

func (cr VehicleCapacityResponse) TableName() string {
	return "vehicle_capacity"
}
