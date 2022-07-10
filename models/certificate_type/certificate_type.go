package models

type CertificateType struct {
	ID          int64
	Name        string
	Icon        string
	Description string
	Status      bool
	UpdatedAt   *string
	CreatedAt   *string
}
