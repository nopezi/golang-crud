package building_assets

type BuildingAssets struct {
	ID                int64
	AssetID           int64
	CertificateTypeID string
	CertificateNumber string
	BuildYear         int64
	SurfaceArea       int64
	BuildingArea      int64
	Direction         string
	NumberOfFloors    int64
	NumberOfBedrooms  int64
	NumberOfBathrooms int64
	ElectricalPower   int64
	Carport           int64
	UpdatedAt         *string
	CreatedAt         *string
}
