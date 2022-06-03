package models

type VehicleAssetsRequest struct {
	AssetID           int64  `json:"asset_id,string"`
	VehicleType       string `json:"vehicle_type"`
	CertificateType   string `json:"certificate_type"`
	CertificateNumber string `json:"certificate_number"`
	Series            string `json:"series"`
	Brand             string `json:"brand"`
	Type              string `json:"type"`
	ProductionYear    string `json:"productionYear"`
	Transmission      string `json:"transmission"`
	MachineCapacity   string `json:"machineCapacity"`
	Color             string `json:"color"`
	NumberOfSeat      string `json:"numberOfSeat"`
	NumberOfUsage     string `json:"numberOfUsage"`
	MachineNumber     string `json:"machineNumber"`
	BodyNumber        string `json:"bodyNumber"`
	LicenceDate       string `json:"licenceDate"`
}

type VehicleAssetsResponse struct {
	ID                int64  `json:"id,string"`
	AssetID           int64  `json:"asset_id,string"`
	VehicleType       string `json:"vehicle_type"`
	CertificateType   string `json:"certificate_type"`
	CertificateNumber string `json:"certificate_number"`
	Series            string `json:"series"`
	Brand             string `json:"brand"`
	Type              string `json:"type"`
	ProductionYear    string `json:"productionYear"`
	Transmission      string `json:"transmission"`
	MachineCapacity   string `json:"machineCapacity"`
	Color             string `json:"color"`
	NumberOfSeat      string `json:"numberOfSeat"`
	NumberOfUsage     string `json:"numberOfUsage"`
	MachineNumber     string `json:"machineNumber"`
	BodyNumber        string `json:"bodyNumber"`
	LicenceDate       string `json:"licenceDate"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

func (p VehicleAssetsRequest) ParseRequest() VehicleAssets {
	return VehicleAssets{
		AssetID:           p.AssetID,
		VehicleType:       p.VehicleType,
		CertificateType:   p.CertificateType,
		CertificateNumber: p.CertificateNumber,
		Series:            p.Series,
		Brand:             p.Brand,
		Type:              p.Type,
		ProductionYear:    p.ProductionYear,
		Transmission:      p.Transmission,
		MachineCapacity:   p.MachineCapacity,
		Color:             p.Color,
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
		CertificateType:   p.CertificateType,
		CertificateNumber: p.CertificateNumber,
		Series:            p.Series,
		Brand:             p.Brand,
		Type:              p.Type,
		ProductionYear:    p.ProductionYear,
		Transmission:      p.Transmission,
		MachineCapacity:   p.MachineCapacity,
		Color:             p.Color,
		NumberOfSeat:      p.NumberOfSeat,
		NumberOfUsage:     p.NumberOfUsage,
		MachineNumber:     p.MachineNumber,
		BodyNumber:        p.BodyNumber,
		LicenceDate:       p.LicenceDate,
		CreatedAt:         p.CreatedAt,
		UpdatedAt:         p.UpdatedAt,
	}
}
