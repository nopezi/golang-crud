package models

type RiskTypeRequest struct {
	ID           int64   `json:"id"`
	RiskTypeCode string  `json:"risk_type_code"`
	RiskType     string  `json:"risk_type"`
	CreatedAt    *string `json:"created_at"`
	UpdatedAt    *string `json:"updated_at"`
}

type RiskTypeResponse struct {
	ID           int64   `json:"id"`
	RiskTypeCode string  `json:"risk_type_code"`
	RiskType     string  `json:"risk_type"`
	CreatedAt    *string `json:"created_at"`
	UpdatedAt    *string `json:"updated_at"`
}

func (p RiskTypeRequest) ParseRequest() RiskType {
	return RiskType{
		ID:           p.ID,
		RiskTypeCode: p.RiskTypeCode,
		RiskType:     p.RiskType,
	}
}

func (p RiskTypeResponse) ParseResponse() RiskType {
	return RiskType{
		ID:           p.ID,
		RiskTypeCode: p.RiskTypeCode,
		RiskType:     p.RiskType,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}

func (pr RiskTypeRequest) TableName() string {
	return "risk_type"
}

func (pr RiskTypeResponse) TableName() string {
	return "risk_type"
}
