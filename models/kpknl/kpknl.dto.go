package models

type KpknlRequest struct {
	ID        int64   `json:"id"`
	Desc      string  `json:"desc"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type KpknlResponse struct {
	ID        int64   `json:"id"`
	Desc      string  `json:"desc"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

func (p KpknlRequest) ParseRequest() Kpknl {
	return Kpknl{
		ID:   p.ID,
		Desc: p.Desc,
	}
}

func (p KpknlResponse) ParseResponse() Kpknl {
	return Kpknl{
		ID:        p.ID,
		Desc:      p.Desc,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (kr KpknlRequest) TableName() string {
	return "ref_kpknl"
}

func (kr KpknlResponse) TableName() string {
	return "ref_kpknl"
}
