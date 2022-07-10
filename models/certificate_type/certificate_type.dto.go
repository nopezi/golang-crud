package models

type CertificateTypeRequest struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	// Status      bool    `json:"status"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type CertificateTypeRequests []map[string]interface{}
type CertificateTypeResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	// Status      bool    `json:"status"`
	Description string  `json:"description"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

func (p CertificateTypeRequest) ParseRequest() CertificateType {
	return CertificateType{
		Name:        p.Name,
		Icon:        p.Icon,
		Description: p.Description,
	}
}

func (p CertificateTypeResponse) ParseResponse() CertificateType {
	return CertificateType{
		ID:          p.ID,
		Name:        p.Name,
		Icon:        p.Icon,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (kr CertificateTypeRequest) TableName() string {
	return "certificate_type"
}

func (kr CertificateTypeResponse) TableName() string {
	return "certificate_type"
}
