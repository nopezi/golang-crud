package models

type SubIncident struct {
	ID                       int64
	KodeKejadian             string
	KodeSubKejadian          string
	KriteriaPenyebabKejadian string
	CreatedAt                *string
	UpdatedAt                *string
}
