package models

type RiskIndicator struct {
	ID                int64
	RiskIndicatorCode string
	RiskIndicator     string
	ActivityID        int64
	ProductID         int64
	CreatedAt         *string
	UpdatedAt         *string
}
