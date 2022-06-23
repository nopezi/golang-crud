package models

type Addresses struct {
	ID         int64
	AssetID    int64
	PostalCode string
	Address    string
	Longitude  string
	Langitude  string
	UpdatedAt  *string
	CreatedAt  *string
}
