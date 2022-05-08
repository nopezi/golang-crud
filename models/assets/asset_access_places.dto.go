package models

type AssetAccessPlacesRequest struct {
	ID            int64 `json:"id,string"`
	AssetID       int64 `json:"asset_id,string"`
	AccessPlaceID int64 `json:"access_place_id,string"`
}

type AssetAccessPlacesResponse struct {
	ID            int64  `json:"id,string"`
	AssetID       int64  `json:"asset_id,string"`
	AccessPlaceID int64  `json:"access_place_id,string"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

func (p AssetAccessPlacesRequest) ParseRequest() AssetAccessPlaces {
	return AssetAccessPlaces{
		ID:            p.ID,
		AssetID:       p.AssetID,
		AccessPlaceID: p.AccessPlaceID,
	}
}

func (p AssetAccessPlacesResponse) ParseResponse() AssetAccessPlaces {
	return AssetAccessPlaces{
		ID:            p.ID,
		AssetID:       p.AssetID,
		AccessPlaceID: p.AccessPlaceID,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}
