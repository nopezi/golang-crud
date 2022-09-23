package models

type Incident struct {
	ID               int64
	KodeKejadian     string
	PenyebabKejadian string
	CreatedAt        *string
	UpdatedAt        *string
}
