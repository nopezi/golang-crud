package models

type ContactsRequest struct {
	AssetID     int64   `json:"asset_id"`
	DebiturName string  `json:"debitur_name"`
	PicName     string  `json:"pic_name"`
	PicPhone    string  `json:"pic_phone"`
	PicEmail    string  `json:"pic_email"`
	Cif         string  `json:"cif"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

type ContactsResponse struct {
	ID          int64   `json:"id"`
	AssetID     int64   `json:"asset_id"`
	DebiturName string  `json:"debitur_name"`
	PicName     string  `json:"pic_name"`
	PicPhone    string  `json:"pic_phone"`
	PicEmail    string  `json:"pic_email"`
	Cif         string  `json:"cif"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

func (p ContactsRequest) ParseRequest() Contacts {
	return Contacts{
		AssetID:     p.AssetID,
		DebiturName: p.DebiturName,
		PicName:     p.PicName,
		PicPhone:    p.PicPhone,
		PicEmail:    p.PicEmail,
		Cif:         p.Cif,
	}
}

func (p ContactsResponse) ParseResponse() Contacts {
	return Contacts{
		ID:          p.ID,
		AssetID:     p.AssetID,
		DebiturName: p.DebiturName,
		PicName:     p.PicName,
		PicPhone:    p.PicPhone,
		PicEmail:    p.PicEmail,
		Cif:         p.Cif,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (c ContactsRequest) TableName() string {
	return "contacts"
}

func (c ContactsResponse) TableName() string {
	return "contacts"
}
