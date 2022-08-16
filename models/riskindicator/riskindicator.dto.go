package models

type RiskIndicatorRequest struct {
	ID            int64   `json:"id"`
	IndicatorCode string  `json:"indicator_code"`
	Name          string  `json:"name"`
	CreatedAt     *string `json:"created_at"`
	UpdatedAt     *string `json:"updated_at"`
}

type RiskIndicatorResponse struct {
	ID            int64   `json:"id"`
	IndicatorCode string  `json:"indicator_code"`
	Name          string  `json:"name"`
	CreatedAt     *string `json:"created_at"`
	UpdatedAt     *string `json:"updated_at"`
}

func (p RiskIndicatorRequest) ParseRequest() RiskIndicator {
	return RiskIndicator{
		ID:            p.ID,
		IndicatorCode: p.IndicatorCode,
		Name:          p.Name,
	}
}

func (p RiskIndicatorResponse) ParseRequest() RiskIndicator {
	return RiskIndicator{
		ID:            p.ID,
		IndicatorCode: p.IndicatorCode,
		Name:          p.Name,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

func (pr RiskIndicatorRequest) TableName() string {
	return "risk_indicator"
}

func (pr RiskIndicatorResponse) TableName() string {
	return "risk_indicator"
}
