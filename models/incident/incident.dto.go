package models

type IncidentRequest struct {
	ID                int64   `json:"id"`
	KodeKejadian      string  `json:"kode_kejadian"`
	PenyebabKejadian1 string  `json:"penyebab_kejadian1"`
	CreatedAt         *string `json:"created_at"`
	UpdatedAt         *string `json:"updated_at"`
}

type IncidentResponse struct {
	ID                int64   `json:"id"`
	KodeKejadian      string  `json:"kode_kejadian"`
	PenyebabKejadian1 string  `json:"penyebab_kejadian1"`
	CreatedAt         *string `json:"created_at"`
	UpdatedAt         *string `json:"updated_at"`
}

func (p IncidentRequest) ParseRequest() Incident {
	return Incident{
		ID:                p.ID,
		KodeKejadian:      p.KodeKejadian,
		PenyebabKejadian1: p.PenyebabKejadian1,
	}
}

func (p IncidentResponse) ParseResponse() Incident {
	return Incident{
		ID:                p.ID,
		KodeKejadian:      p.KodeKejadian,
		PenyebabKejadian1: p.PenyebabKejadian1,
		CreatedAt:         p.CreatedAt,
		UpdatedAt:         p.UpdatedAt,
	}
}

func (pr IncidentRequest) TableName() string {
	return "incident_cause"
}

func (pr IncidentResponse) TableName() string {
	return "incident_cause"
}
