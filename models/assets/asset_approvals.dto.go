package models

type AssetApprovalsRequest struct {
	ID         int64  `json:"id"`
	AssetID    int64  `json:"asset_id"`
	ApprovalID int64  `json:"approval_id"`
	Status     string `json:"status"`
}

type AssetApprovalsResponse struct {
	ID         int64  `json:"id"`
	AssetID    int64  `json:"asset_id"`
	ApprovalID int64  `json:"approval_id"`
	Status     string `json:"status"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func (p AssetApprovalsRequest) ParseRequest() AssetApprovals {
	return AssetApprovals{
		ID:         p.ID,
		AssetID:    p.AssetID,
		ApprovalID: p.ApprovalID,
		Status:     p.Status,
	}
}

func (p AssetApprovalsResponse) ParseResponse() AssetApprovals {
	return AssetApprovals{
		ID:         p.ID,
		AssetID:    p.AssetID,
		ApprovalID: p.ApprovalID,
		Status:     p.Status,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}
