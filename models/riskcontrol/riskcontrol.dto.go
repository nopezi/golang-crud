package models

type RiskControlRequest struct {
	ID          int64   `json:"id"`
	Kode        string  `json:"kode"`
	RiskControl string  `json:"risk_control"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

type RiskControlResponse struct {
	ID          int64   `json:"id"`
	Kode        string  `json:"kode"`
	RiskControl string  `json:"risk_control"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

func (p RiskControlRequest) ParseRequest() RiskControl {
	return RiskControl{
		ID:          p.ID,
		Kode:        p.Kode,
		RiskControl: p.RiskControl,
	}
}

func (p RiskControlResponse) ParseResonse() RiskControl {
	return RiskControl{
		ID:          p.ID,
		Kode:        p.Kode,
		RiskControl: p.RiskControl,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (rc RiskControlRequest) TableName() string {
	return "risk_control"
}

func (rc RiskControlResponse) TableName() string {
	return "risk_control"
}
