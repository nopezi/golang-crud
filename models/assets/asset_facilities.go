package models

type AssetFacilities struct {
	ID         int64
	AssetID    int64
	FacilityID int64
	Status     bool
	UpdatedAt  *string
	CreatedAt  *string
}
