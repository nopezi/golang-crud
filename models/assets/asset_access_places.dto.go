package models

type AssetAccessPlacesRequest struct {
	ID            int64 `json:"id"`
	AssetID       int64 `json:"asset_id"`
	AccessPlaceID int64 `json:"access_place_id"`
	Status        bool  `json:"status"`
}

type AssetAccessPlacesResponse struct {
	ID            int64   `json:"id"`
	AssetID       int64   `json:"asset_id"`
	AccessPlaceID int64   `json:"access_place_id"`
	Status        bool    `json:"status"`
	CreatedAt     *string `json:"created_at"`
	UpdatedAt     *string `json:"updated_at"`
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

func (aap AssetAccessPlacesRequest) TableName() string {
	return "asset_access_places"
}

func (aap AssetAccessPlacesResponse) TableName() string {
	return "asset_access_places"
}
