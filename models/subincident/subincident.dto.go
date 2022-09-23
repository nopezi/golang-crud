package models

import "riskmanagement/lib"

type SubIncidentRequest struct {
	ID                       int64   `json:"id"`
	KodeKejadian             string  `json:"kode_kejadian"`
	KodeSubKejadian          string  `json:"kode_sub_kejadian"`
	KriteriaPenyebabKejadian string  `json:"kriteria_penyebab_kejadian"`
	CreatedAt                *string `json:"created_at"`
	UpdatedAt                *string `json:"updated_at"`
}

type SubIncidentResponse struct {
	ID                       int64   `json:"id"`
	KodeKejadian             string  `json:"kode_kejadian"`
	KodeSubKejadian          string  `json:"kode_sub_kejadian"`
	KriteriaPenyebabKejadian string  `json:"kriteria_penyebab_kejadian"`
	CreatedAt                *string `json:"created_at"`
	UpdatedAt                *string `json:"updated_at"`
}

type SubIncidentListFilter struct {
	ID                       lib.NullInt64  `json:"id"`
	KodeKejadian             lib.NullString `json:"kode_kejadian"`
	PenyebabKejadian         lib.NullString `json:"penyebab_kejadian"`
	KodeSubKejadian          lib.NullString `json:"kode_sub_kejadian"`
	KriteriaPenyebabKejadian lib.NullString `json:"kriteria_penyebab_kejadian"`
	CreatedAt                lib.NullString `json:"created_at"`
	UpdatedAt                lib.NullString `json:"updated_at"`
}

type SubIncidentFilterRequest struct {
	KodeKejadian string `json:"kode_kejadian"`
}

type SubIncidentResponses struct {
	ID                       int64   `json:"id"`
	KodeKejadian             string  `json:"kode_kejadian"`
	PenyebabKejadian         string  `json:"penyebab_kejadian"`
	KodeSubKejadian          string  `json:"kode_sub_kejadian"`
	KriteriaPenyebabKejadian string  `json:"kriteria_penyebab_kejadian"`
	CreatedAt                *string `json:"created_at"`
	UpdatedAt                *string `json:"updated_at"`
}

func (p SubIncidentRequest) ParseRequest() SubIncident {
	return SubIncident{
		ID:                       p.ID,
		KodeKejadian:             p.KodeKejadian,
		KodeSubKejadian:          p.KodeSubKejadian,
		KriteriaPenyebabKejadian: p.KriteriaPenyebabKejadian,
	}
}

func (p SubIncidentResponse) ParseResponse() SubIncident {
	return SubIncident{
		ID:                       p.ID,
		KodeKejadian:             p.KodeKejadian,
		KodeSubKejadian:          p.KodeSubKejadian,
		KriteriaPenyebabKejadian: p.KriteriaPenyebabKejadian,
		CreatedAt:                p.CreatedAt,
		UpdatedAt:                p.UpdatedAt,
	}
}

func (pr SubIncidentRequest) TableName() string {
	return "sub_incident_cause"
}

func (pr SubIncidentResponse) TableName() string {
	return "sub_incident_cause"
}
