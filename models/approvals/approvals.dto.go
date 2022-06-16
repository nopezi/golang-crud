package models

type ApprovalsRequest struct {
	AssetID        int64   `json:"asset_id"`
	CheckerID      string  `json:"checker_id"`
	CheckerDesc    string  `json:"checker_desc"`
	CheckerComment *string `json:"checker_domment"`
	CheckerDate    *string `json:"checker_date"`
	SignerID       string  `json:"signer_id"`
	SignerDesc     string  `json:"signer_desc"`
	SignerComment  *string `json:"signer_comment"`
	SignerDate     *string `json:"signer_date"`
	CreatedAt      *string `json:"created_at"`
	UpdatedAt      *string `json:"updated_at"`
}

type ApprovalsResponse struct {
	ID             int64   `json:"id"`
	AssetID        int64   `json:"asset_id"`
	CheckerID      string  `json:"checker_id"`
	CheckerDesc    string  `json:"checker_desc"`
	CheckerComment *string `json:"checkerd_domment"`
	CheckerDate    *string `json:"checker_date"`
	SignerID       string  `json:"signer_id"`
	SignerDesc     string  `json:"signer_desc"`
	SignerComment  *string `json:"signer_comment"`
	SignerDate     *string `json:"signer_date"`
	CreatedAt      *string `json:"created_at"`
	UpdatedAt      *string `json:"updated_at"`
}

func (p ApprovalsRequest) ParseRequest() Approvals {
	return Approvals{
		AssetID:        p.AssetID,
		CheckerID:      p.CheckerID,
		CheckerDesc:    p.CheckerDesc,
		CheckerComment: p.CheckerComment,
		CheckerDate:    p.CheckerDate,
		SignerID:       p.SignerID,
		SignerDesc:     p.SignerDesc,
		SignerComment:  p.SignerComment,
		SignerDate:     p.SignerDate,
	}
}

func (p ApprovalsResponse) ParseResponse() Approvals {
	return Approvals{
		ID:             p.ID,
		AssetID:        p.AssetID,
		CheckerID:      p.CheckerID,
		CheckerDesc:    p.CheckerDesc,
		CheckerComment: p.CheckerComment,
		CheckerDate:    p.CheckerDate,
		SignerID:       p.SignerID,
		SignerDesc:     p.SignerDesc,
		SignerComment:  p.SignerComment,
		SignerDate:     p.SignerDate,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
	}
}

func (a ApprovalsRequest) TableName() string {
	return "approvals"
}

func (a ApprovalsResponse) TableName() string {
	return "approvals"
}
