package postalcode

type PostalcodeRequest struct {
	ID         string  `json:"id"`
	PostalCode string  `json:"postal_code" binding:"required"`
	Region     string  `json:"region"`
	District   string  `json:"district"`
	City       string  `json:"city"`
	Province   string  `json:"province"`
	Enabled    string  `json:"enabled"`
	CreatedAt  *string `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
}

type PostalcodeResponse struct {
	ID         string  `json:"id"`
	PostalCode string  `json:"postal_code"`
	Region     string  `json:"region"`
	District   string  `json:"district"`
	City       string  `json:"city"`
	Province   string  `json:"province"`
	Enabled    string  `json:"enabled"`
	CreatedAt  *string `json:"created_at"`
	CpdatedAt  *string `json:"updated_at"`
}

func (pc PostalcodeRequest) TableName() string {
	return "ref_postal_code"
}

func (pr PostalcodeResponse) TableName() string {
	return "ref_postal_code"
}
