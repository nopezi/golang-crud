package models

type RiskIssueRequest struct {
	ID        int64   `json:"id"`
	RiskCode  string  `json:"risk_code"`
	Name      string  `json:"name"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type RiskIssueResponse struct {
	ID        int64   `json:"id"`
	RiskCode  string  `json:"risk_code"`
	Name      string  `json:"name"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

func (p RiskIssueRequest) ParseRequest() RiskIssue {
	return RiskIssue{
		ID:       p.ID,
		RiskCode: p.RiskCode,
		Name:     p.Name,
	}
}

func (p RiskIssueResponse) ParseRequest() RiskIssue {
	return RiskIssue{
		ID:        p.ID,
		RiskCode:  p.RiskCode,
		Name:      p.Name,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (pr RiskIssueRequest) TableName() string {
	return "risk_issue"
}

func (pr RiskIssueResponse) TableName() string {
	return "risk_issue"
}
