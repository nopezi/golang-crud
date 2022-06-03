package building_assets

type BuildingAssetsRequest struct {
	CertificateType   string `json:"certificateType"`
	CertificateNumber string `json:"certificateNumber"`
	BuildYear         string `json:"buildYear"`
	SurfaceArea       string `json:"burfaceArea"`
	BuildingArea      string `json:"buildingArea"`
	Derection         string `json:"derection"`
	NumberOfFloors    string `json:"numberOfFloors`
	NumberOfBedrooms  string `json:"numberOfBedrooms"`
	NumberOfBathrooms string `json:"numberOfBathrooms"`
	ElectricalPower   string `json:"electricalPower"`
	Carport           string `json:"carport"`
}

type BuildingAssetsResponse struct {
	ID                int64  `json:"id,string"`
	CertificateType   string `json:"certificateType"`
	CertificateNumber string `json:"certificateNumber"`
	BuildYear         string `json:"buildYear"`
	SurfaceArea       string `json:"burfaceArea"`
	BuildingArea      string `json:"buildingArea"`
	Derection         string `json:"derection"`
	NumberOfFloors    string `json:"numberOfFloors`
	NumberOfBedrooms  string `json:"numberOfBedrooms"`
	NumberOfBathrooms string `json:"numberOfBathrooms"`
	ElectricalPower   string `json:"electricalPower"`
	Carport           string `json:"carport"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

func (p BuildingAssetsRequest) ParseRequest() BuildingAssets {
	return BuildingAssets{
		CertificateType:   p.CertificateType,
		CertificateNumber: p.CertificateNumber,
		BuildYear:         p.BuildYear,
		SurfaceArea:       p.SurfaceArea,
		BuildingArea:      p.BuildingArea,
		Derection:         p.Derection,
		NumberOfFloors:    p.NumberOfFloors,
		NumberOfBedrooms:  p.NumberOfBedrooms,
		NumberOfBathrooms: p.NumberOfBathrooms,
		ElectricalPower:   p.ElectricalPower,
		Carport:           p.Carport,
	}
}

func (p BuildingAssetsResponse) ParseResponse() BuildingAssets {
	return BuildingAssets{
		ID:                p.ID,
		CertificateType:   p.CertificateType,
		CertificateNumber: p.CertificateNumber,
		BuildYear:         p.BuildYear,
		SurfaceArea:       p.SurfaceArea,
		BuildingArea:      p.BuildingArea,
		Derection:         p.Derection,
		NumberOfFloors:    p.NumberOfFloors,
		NumberOfBedrooms:  p.NumberOfBedrooms,
		NumberOfBathrooms: p.NumberOfBathrooms,
		ElectricalPower:   p.ElectricalPower,
		Carport:           p.Carport,
		CreatedAt:         p.CreatedAt,
		UpdatedAt:         p.UpdatedAt,
	}
}
