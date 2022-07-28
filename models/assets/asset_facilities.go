package models

type AssetFacilities struct {
	ID         int64
	AssetID    int64
	FacilityID int64
	Status     string
	UpdatedAt  *string
	CreatedAt  *string
}
