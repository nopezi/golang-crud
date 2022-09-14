package models

type RiskIndicatorRequest struct {
	ID                int64   `json:"id"`
	RiskIndicatorCode string  `json:"risk_indicator_code"`
	RiskIndicator     string  `json:"risk_indicator"`
	ActivityID        int64   `json:"activity_id"`
	ProductID         int64   `json:"product_id"`
	CreatedAt         *string `json:"created_at"`
	UpdatedAt         *string `json:"updated_at"`
}

type RiskIndicatorResponse struct {
	ID                int64   `json:"id"`
	RiskIndicatorCode string  `json:"risk_indicator_code"`
	RiskIndicator     string  `json:"risk_indicator"`
	ActivityID        int64   `json:"activity_id"`
	ProductID         int64   `json:"product_id"`
	CreatedAt         *string `json:"created_at"`
	UpdatedAt         *string `json:"updated_at"`
}

type RiskIndicatorResponses struct {
	ID                int64   `json:"id"`
	RiskIndicatorCode string  `json:"risk_indicator_code"`
	RiskIndicator     string  `json:"risk_indicator"`
	ActivityID        int64   `json:"activity_id"`
	Activity          string  `json:"activity"`
	ProductID         int64   `json:"product_id"`
	Product           string  `json:"product"`
	CreatedAt         *string `json:"created_at"`
	UpdatedAt         *string `json:"updated_at"`
}

func (p RiskIndicatorRequest) ParseRequest() RiskIndicator {
	return RiskIndicator{
		ID:                p.ID,
		RiskIndicatorCode: p.RiskIndicatorCode,
		RiskIndicator:     p.RiskIndicator,
		ActivityID:        p.ActivityID,
		ProductID:         p.ProductID,
	}
}

func (p RiskIndicatorResponse) ParseRequest() RiskIndicator {
	return RiskIndicator{
		ID:                p.ID,
		RiskIndicatorCode: p.RiskIndicatorCode,
		RiskIndicator:     p.RiskIndicator,
		ActivityID:        p.ActivityID,
		ProductID:         p.ProductID,
		CreatedAt:         p.CreatedAt,
		UpdatedAt:         p.UpdatedAt,
	}
}

func (pr RiskIndicatorRequest) TableName() string {
	return "risk_indicator"
}

func (pr RiskIndicatorResponse) TableName() string {
	return "risk_indicator"
}
