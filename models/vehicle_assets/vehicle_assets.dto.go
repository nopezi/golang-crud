package models

type VehicleAssetsRequest struct {
	AssetID           int64  `json:"asset_id"`
	VehicleType       string `json:"vehicle_type"`
	CertificateTypeID int64  `json:"certificate_type_id"`
	CertificateNumber string `json:"certificate_number"`
	Series            string `json:"series"`
	BrandID           int64  `json:"brand_id"`
	Type              string `json:"type"`
	ProductionYear    string `json:"productionYear"`
	TransmissionID    int64  `json:"transmission_id"`
	MachineCapacityID int64  `json:"machineCapacity_id"`
	ColorID           int64  `json:"color_id"`
	NumberOfSeat      int64  `json:"number_of_seat"`
	NumberOfUsage     string `json:"number_of_usage"`
	MachineNumber     string `json:"machine_number"`
	BodyNumber        string `json:"body_number"`
	LicenceDate       string `json:"licence_date"`
}

type VehicleAssetsResponse struct {
	ID                int64  `json:"id"`
	AssetID           int64  `json:"asset_id"`
	VehicleType       string `json:"vehicle_type"`
	CertificateTypeID int64  `json:"certificate_type_id"`
	CertificateNumber string `json:"certificate_number"`
	Series            string `json:"series"`
	BrandID           int64  `json:"brand_id"`
	Type              string `json:"type"`
	ProductionYear    string `json:"production_year"`
	TransmissionID    int64  `json:"transmission_id"`
	MachineCapacityID int64  `json:"machine_capacity_id"`
	ColorID           int64  `json:"color_id"`
	NumberOfSeat      int64  `json:"number_of_seat"`
	NumberOfUsage     string `json:"number_of_usage"`
	MachineNumber     string `json:"machine_number"`
	BodyNumber        string `json:"body_number"`
	LicenceDate       string `json:"licence_date"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

func (p VehicleAssetsRequest) ParseRequest() VehicleAssets {
	return VehicleAssets{
		AssetID:           p.AssetID,
		VehicleType:       p.VehicleType,
		CertificateTypeID: p.CertificateTypeID,
		CertificateNumber: p.CertificateNumber,
		Series:            p.Series,
		BrandID:           p.BrandID,
		Type:              p.Type,
		ProductionYear:    p.ProductionYear,
		TransmissionID:    p.TransmissionID,
		MachineCapacityID: p.MachineCapacityID,
		ColorID:           p.ColorID,
		NumberOfSeat:      p.NumberOfSeat,
		NumberOfUsage:     p.NumberOfUsage,
		MachineNumber:     p.MachineNumber,
		BodyNumber:        p.BodyNumber,
		LicenceDate:       p.LicenceDate,
	}
}

func (p VehicleAssetsResponse) ParseResponse() VehicleAssets {
	return VehicleAssets{
		ID:                p.ID,
		AssetID:           p.AssetID,
		VehicleType:       p.VehicleType,
		CertificateTypeID: p.CertificateTypeID,
		CertificateNumber: p.CertificateNumber,
		Series:            p.Series,
		BrandID:           p.BrandID,
		Type:              p.Type,
		ProductionYear:    p.ProductionYear,
		TransmissionID:    p.TransmissionID,
		MachineCapacityID: p.MachineCapacityID,
		ColorID:           p.ColorID,
		NumberOfSeat:      p.NumberOfSeat,
		NumberOfUsage:     p.NumberOfUsage,
		MachineNumber:     p.MachineNumber,
		BodyNumber:        p.BodyNumber,
		LicenceDate:       p.LicenceDate,
		CreatedAt:         p.CreatedAt,
		UpdatedAt:         p.UpdatedAt,
	}
}
