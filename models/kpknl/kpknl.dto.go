package models

type KpknlRequest struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Fax     string `json:"fax"`
}

type KpknlResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Fax       string `json:"fax"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (p KpknlRequest) ParseRequest() Kpknl {
	return Kpknl{
		ID:      p.ID,
		Name:    p.Name,
		Address: p.Address,
		Phone:   p.Phone,
		Fax:     p.Fax,
	}
}

func (p KpknlResponse) ParseResponse() Kpknl {
	return Kpknl{
		ID:        p.ID,
		Name:      p.Name,
		Address:   p.Address,
		Phone:     p.Phone,
		Fax:       p.Fax,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
