package models

type AssetAccessPlaces struct {
	ID            int64
	AssetID       int64
	AccessPlaceID int64
	Status        string
	UpdatedAt     *string
	CreatedAt     *string
}
