package models

type ContactsRequest struct {
	DebiturName string `json:"debitur_name"`
	PicName     string `json:"pic_name"`
	PicPhone    string `json:"pic_phone"`
	PicEmail    string `json:"pic_email"`
	Cif         string `json:"cif"`
}

type ContactsResponse struct {
	ID          int64  `json:"id,string"`
	DebiturName string `json:"debitur_name"`
	PicName     string `json:"pic_name"`
	PicPhone    string `json:"pic_phone"`
	PicEmail    string `json:"pic_email"`
	Cif         string `json:"cif"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (p ContactsRequest) ParseRequest() Contacts {
	return Contacts{
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
		DebiturName: p.DebiturName,
		PicName:     p.PicName,
		PicPhone:    p.PicPhone,
		PicEmail:    p.PicEmail,
		Cif:         p.Cif,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}
