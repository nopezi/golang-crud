package models

type AssetApprovals struct {
	ID         int64
	AssetID    int64
	ApprovalID int64
	Status     string
	UpdatedAt  string
	CreatedAt  string
}
