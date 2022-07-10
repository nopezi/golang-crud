package models

type AssetAccessPlaces struct {
	ID            int64
	AssetID       int64
	AccessPlaceID int64
	Status        bool
	UpdatedAt     *string
	CreatedAt     *string
}
