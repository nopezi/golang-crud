package models

type AssetFacilitiesRequest struct {
	ID           int64 `json:"id,string"`
	AssetID      int64 `json:"asset_id,string"`
	FacilitiesID int64 `json:"access_place_id,string"`
}

type AssetFacilitiesResponse struct {
	ID           int64  `json:"id,string"`
	AssetID      int64  `json:"asset_id,string"`
	FacilitiesID int64  `json:"access_place_id,string"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func (p AssetFacilitiesRequest) ParseRequest() AssetFacilities {
	return AssetFacilities{
		ID:           p.ID,
		AssetID:      p.AssetID,
		FacilitiesID: p.FacilitiesID,
	}
}

func (p AssetFacilitiesResponse) ParseResponse() AssetFacilities {
	return AssetFacilities{
		ID:           p.ID,
		AssetID:      p.AssetID,
		FacilitiesID: p.FacilitiesID,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}
