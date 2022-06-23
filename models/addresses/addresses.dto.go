package models

type AddressesRequest struct {
	AssetID    int64   `json:"asset_id"`
	PostalCode string  `json:"postalcode_id"`
	Address    string  `json:"address"`
	Longitude  string  `json:"longitude"`
	Langitude  string  `json:"langitude"`
	CreatedAt  *string `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
}

type AddressesResponse struct {
	ID         int64   `json:"id"`
	AssetID    int64   `json:"asset_id"`
	Address    string  `json:"address"`
	Longitude  string  `json:"longitude"`
	Langitude  string  `json:"langitude"`
	Longlat    string  `json:"longlat"`
	PostalCode string  `json:"postal_code"`
	Region     string  `json:"region"`
	District   string  `json:"district"`
	City       string  `json:"city"`
	Province   string  `json:"province"`
	CreatedAt  *string `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
}

func (a AddressesRequest) TableName() string {
	return "addresses"
}

func (a AddressesResponse) TableName() string {
	return "addresses"
}
