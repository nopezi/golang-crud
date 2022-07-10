package postalcode

type Postalcode struct {
	ID         int64   `json:"id"`
	PostalCode string  `json:"postal_code"`
	Region     string  `json:"region"`
	District   string  `json:"district"`
	City       string  `json:"city"`
	Province   string  `json:"province"`
	Enabled    string  `json:"enabled"`
	CreatedAt  *string `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
}
