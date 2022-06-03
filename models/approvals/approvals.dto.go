package models

type ApprovalsRequest struct {
	CheckerID      string `json:"checkerID"`
	CheckerDesc    string `json:"checkerDesc"`
	CheckerComment string `json:"checkerComment"`
	CheckerDate    string `json:"checkerDate"`
	SignerID       string `json:"signerID"`
	SignerDesc     string `json:"signerDesc"`
	SignerComment  string `json:"signerComment"`
	SignerDate     string `json:"signerDate"`
}

type ApprovalsResponse struct {
	ID int64 `json:"id,string"`

	CheckerID      string `json:"checkerID"`
	CheckerDesc    string `json:"checkerDesc"`
	CheckerComment string `json:"checkerComment"`
	CheckerDate    string `json:"checkerDate"`
	SignerID       string `json:"signerID"`
	SignerDesc     string `json:"signerDesc"`
	SignerComment  string `json:"signerComment"`
	SignerDate     string `json:"signerDate"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
}

func (p ApprovalsRequest) ParseRequest() Approvals {
	return Approvals{
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
