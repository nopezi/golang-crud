package models

type VehicleColorRequest struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Icon        string  `json:"icon"`
	Description string  `json:"description"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

type VehicleColorResponse struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Icon        string  `json:"icon"`
	Description string  `json:"description"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

func (p VehicleColorRequest) ParseRequest() VehicleColor {
	return VehicleColor{
		ID:   p.ID,
		Name: p.Name,
	}
}

func (p VehicleColorResponse) ParseResponse() VehicleColor {
	return VehicleColor{
		ID:        p.ID,
		Name:      p.Name,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (cr VehicleColorRequest) TableName() string {
	return "vehicle_color"
}

func (cr VehicleColorResponse) TableName() string {
	return "vehicle_color"
}
