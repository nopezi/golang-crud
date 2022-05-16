package models

type Approvals struct {
	ID             int64
	CheckerID      string
	CheckerDesc    string
	CheckerComment string
	CheckerDate    string
	SignerID       string
	SignerDesc     string
	SignerComment  string
	SignerDate     string
	UpdatedAt      string
	CreatedAt      string
}
