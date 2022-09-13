package models

type RiskIssueRequest struct {
	ID             int64   `json:"id"`
	RiskTypeID     int64   `json:"risk_type_id"`
	RiskIssueCode  string  `json:"risk_issue_code"`
	RiskIssue      string  `json:"risk_issue"`
	Deskripsi      string  `json:"deskripsi"`
	KategoriRisiko string  `json:"kategori_risiko"`
	Status         string  `json:"status"`
	CreatedAt      *string `json:"created_at"`
	UpdatedAt      *string `json:"updated_at"`
}

type RiskIssueResponse struct {
	ID             int64   `json:"id"`
	RiskTypeID     int64   `json:"risk_type_id"`
	RiskIssueCode  string  `json:"risk_issue_code"`
	RiskIssue      string  `json:"risk_issue"`
	Deskripsi      string  `json:"deskripsi"`
	KategoriRisiko string  `json:"kategori_risiko"`
	Status         string  `json:"status"`
	CreatedAt      *string `json:"created_at"`
	UpdatedAt      *string `json:"updated_at"`
}

func (p RiskIssueRequest) ParseRequest() RiskIssue {
	return RiskIssue{
		ID:             p.ID,
		RiskTypeID:     p.RiskTypeID,
		RiskIssueCode:  p.RiskIssueCode,
		RiskIssue:      p.RiskIssue,
		Deskripsi:      p.Deskripsi,
		KategoriRisiko: p.KategoriRisiko,
		Status:         p.Status,
	}
}

func (p RiskIssueResponse) ParseRequest() RiskIssue {
	return RiskIssue{
		ID:             p.ID,
		RiskTypeID:     p.RiskTypeID,
		RiskIssueCode:  p.RiskIssueCode,
		RiskIssue:      p.RiskIssue,
		Deskripsi:      p.Deskripsi,
		KategoriRisiko: p.KategoriRisiko,
		Status:         p.Status,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
	}
}

func (pr RiskIssueRequest) TableName() string {
	return "risk_issue"
}

func (pr RiskIssueResponse) TableName() string {
	return "risk_issue"
}
