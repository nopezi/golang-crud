package models

type AssetFacilitiesRequest struct {
	ID         int64   `json:"id"`
	AssetID    int64   `json:"asset_id"`
	FacilityID int64   `json:"access_place_id"`
	CreatedAt  *string `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
}

type AssetFacilitiesResponse struct {
	ID         int64   `json:"id"`
	AssetID    int64   `json:"asset_id"`
	FacilityID int64   `json:"access_place_id"`
	CreatedAt  *string `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
}

func (p AssetFacilitiesRequest) ParseRequest() AssetFacilities {
	return AssetFacilities{
		ID:         p.ID,
		AssetID:    p.AssetID,
		FacilityID: p.FacilityID,
	}
}

func (p AssetFacilitiesResponse) ParseResponse() AssetFacilities {
	return AssetFacilities{
		ID:         p.ID,
		AssetID:    p.AssetID,
		FacilityID: p.FacilityID,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}

func (af AssetFacilitiesRequest) TableName() string {
	return "asset_facilities"
}

func (af AssetFacilitiesResponse) TableName() string {
	return "asset_facilities"
}
