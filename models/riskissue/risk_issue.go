package models

type RiskIssue struct {
	ID             int64
	RiskTypeID     int64
	RiskIssueCode  string
	RiskIssue      string
	Deskripsi      string
	KategoriRisiko string
	Status         string
	CreatedAt      *string
	UpdatedAt      *string
}
