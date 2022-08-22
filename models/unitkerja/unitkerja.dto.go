package models

type UnitKerjaRequest struct {
	ID        int64   `json:"id"`
	KodeUker  int64   `json:"kode_uker"`
	NamaUker  string  `json:"nama_uker"`
	Status    int64   `json:"status"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type UnitKerjaResponse struct {
	ID        int64   `json:"id"`
	KodeUker  int64   `json:"kode_uker"`
	NamaUker  string  `json:"nama_uker"`
	Status    int64   `json:"status"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

func (p UnitKerjaRequest) ParseRequest() UnitKerja {
	return UnitKerja{
		ID:       p.ID,
		KodeUker: p.KodeUker,
		NamaUker: p.NamaUker,
		Status:   p.Status,
	}
}

func (p UnitKerjaResponse) ParseRequest() UnitKerja {
	return UnitKerja{
		ID:        p.ID,
		KodeUker:  p.KodeUker,
		NamaUker:  p.NamaUker,
		Status:    p.Status,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (pr UnitKerjaRequest) TableName() string {
	return "unit_kerja"
}

func (pr UnitKerjaResponse) TableName() string {
	return "unit_kerja"
}
