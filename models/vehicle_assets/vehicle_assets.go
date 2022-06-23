package models

type VehicleAssets struct {
	ID                int64
	AssetID           int64
	VehicleTypeID     int64
	CertificateTypeID int64
	CertificateNumber string
	Series            string
	BrandID           int64
	Type              string
	ProductionYear    string
	TransmissionID    int64
	MachineCapacityID int64
	ColorID           int64
	NumberOfSeat      int64
	NumberOfUsage     string
	MachineNumber     string
	BodyNumber        string
	LicenceDate       string
	UpdatedAt         *string
	CreatedAt         *string
	// BrandName           string
	// TransmissionName    string
	// MachineCapacityName string
	// ColorName           string
}
