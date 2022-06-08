package models

type AddressesRequest struct {
	PostalCodeID int64  `json:"postalcode_id"`
	AssetID      int64  `json:"asset_id"`
	Address      string `json:"address"`
	Longitude    string `json:"longitude"`
	Langitude    string `json:"langitude"`
}

type AddressesResponse struct {
	ID           int64 `json:"id"`
	AssetID      int64
	PostalCodeID int64  `json:"postalcode_id"`
	Address      string `json:"address"`
	Longitude    string `json:"longitude"`
	Langitude    string `json:"langitude"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func (p AddressesRequest) ParseRequest() Addresses {
	return Addresses{
		PostalCodeID: p.PostalCodeID,
		Address:      p.Address,
		Longitude:    p.Longitude,
		Langitude:    p.Langitude,
	}
}

func (p AddressesResponse) ParseResponse() Addresses {
	return Addresses{
		ID:           p.ID,
		AssetID:      p.AssetID,
		PostalCodeID: p.PostalCodeID,
		Address:      p.Address,
		Longitude:    p.Longitude,
		Langitude:    p.Langitude,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}
