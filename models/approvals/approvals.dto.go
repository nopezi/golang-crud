package models

type ApprovalsRequest struct {
	ID          int64  `json:"id,string"`
	DebiturName string `json:"debitur_name"`
	PicName     string `json:"pic_name"`
	PicPhone    string `json:"pic_phone"`
	PicEmail    string `json:"pic_email"`
	Cif         string `json:"cif"`
}

type ApprovalsResponse struct {
	ID          int64  `json:"id,string"`
	DebiturName string `json:"debitur_name"`
	PicName     string `json:"pic_name"`
	PicPhone    string `json:"pic_phone"`
	PicEmail    string `json:"pic_email"`
	Cif         string `json:"cif"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (p ApprovalsRequest) ParseRequest() Approvals {
	return Approvals{
		ID:          p.ID,
		DebiturName: p.DebiturName,
		PicName:     p.PicName,
		PicPhone:    p.PicPhone,
		PicEmail:    p.PicEmail,
		Cif:         p.Cif,
	}
}

func (p ApprovalsResponse) ParseResponse() Approvals {
	return Approvals{
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
