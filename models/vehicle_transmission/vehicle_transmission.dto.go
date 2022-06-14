package models

type VehicleTransmissionRequest struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Icon        string  `json:"icon"`
	Description string  `json:"description"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

type VehicleTransmissionResponse struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Icon        string  `json:"icon"`
	Description string  `json:"description"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

func (p VehicleTransmissionRequest) ParseRequest() VehicleTransmission {
	return VehicleTransmission{
		ID:   p.ID,
		Name: p.Name,
	}
}

func (p VehicleTransmissionResponse) ParseResponse() VehicleTransmission {
	return VehicleTransmission{
		ID:        p.ID,
		Name:      p.Name,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (cr VehicleTransmissionRequest) TableName() string {
	return "vehicle_transmission"
}

func (cr VehicleTransmissionResponse) TableName() string {
	return "vehicle_transmission"
}
