package building_assets

type BuildingAssetsRequest struct {
	AssetID           int64  `json:"asset_id"`
	CertificateType   string `json:"certificate_type"`
	CertificateNumber string `json:"certificate_number"`
	BuildYear         int64  `json:"build_year"`
	SurfaceArea       int64  `json:"burface_area"`
	BuildingArea      int64  `json:"building_area"`
	Derection         string `json:"derection"`
	NumberOfFloors    int64  `json:"number_of_floors`
	NumberOfBedrooms  int64  `json:"number_of_bedrooms"`
	NumberOfBathrooms int64  `json:"number_of_bathrooms"`
	ElectricalPower   int64  `json:"electrical_power"`
	Carport           int64  `json:"carport"`
}

type BuildingAssetsResponse struct {
	ID                int64  `json:"id"`
	AssetID           int64  `json:"asset_id"`
	CertificateType   string `json:"certificate_type"`
	CertificateNumber string `json:"certificate_number"`
	BuildYear         int64  `json:"build_year"`
	SurfaceArea       int64  `json:"burface_area"`
	BuildingArea      int64  `json:"building_area"`
	Derection         string `json:"derection"`
	NumberOfFloors    int64  `json:"number_of_floors`
	NumberOfBedrooms  int64  `json:"number_of_bedrooms"`
	NumberOfBathrooms int64  `json:"number_of_bathrooms"`
	ElectricalPower   int64  `json:"electrical_power"`
	Carport           int64  `json:"carport"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

func (p BuildingAssetsRequest) ParseRequest() BuildingAssets {
	return BuildingAssets{
		AssetID:           p.AssetID,
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
		AssetID:           p.AssetID,
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
