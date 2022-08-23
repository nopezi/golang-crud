package models

type UnitKerjaRequest struct {
	ID         int64   `json:"id"`
	KodeUker   int64   `json:"kode_uker"`
	NamaUker   string  `json:"nama_uker"`
	KodeCabang int64   `json:"kode_cabang"`
	NamaCabang string  `json:"nama_cabang"`
	KanwilID   int64   `json:"kanwil_id"`
	KodeKanwil string  `json:"kode_kanwil"`
	Kanwil     string  `json:"kanwil"`
	Status     int64   `json:"status"`
	CreatedAt  *string `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
}

type UnitKerjaResponse struct {
	ID         int64   `json:"id"`
	KodeUker   int64   `json:"kode_uker"`
	NamaUker   string  `json:"nama_uker"`
	KodeCabang int64   `json:"kode_cabang"`
	NamaCabang string  `json:"nama_cabang"`
	KanwilID   int64   `json:"kanwil_id"`
	KodeKanwil string  `json:"kode_kanwil"`
	Kanwil     string  `json:"kanwil"`
	Status     int64   `json:"status"`
	CreatedAt  *string `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
}

func (p UnitKerjaRequest) ParseRequest() UnitKerja {
	return UnitKerja{
		ID:         p.ID,
		KodeUker:   p.KodeUker,
		NamaUker:   p.NamaUker,
		KodeCabang: p.KanwilID,
		NamaCabang: p.NamaCabang,
		KanwilID:   p.KanwilID,
		KodeKanwil: p.KodeKanwil,
		Kanwil:     p.Kanwil,
		Status:     p.Status,
	}
}

func (p UnitKerjaResponse) ParseRequest() UnitKerja {
	return UnitKerja{
		ID:         p.ID,
		KodeUker:   p.KodeUker,
		NamaUker:   p.NamaUker,
		KodeCabang: p.KanwilID,
		NamaCabang: p.NamaCabang,
		KanwilID:   p.KanwilID,
		KodeKanwil: p.KodeKanwil,
		Kanwil:     p.Kanwil,
		Status:     p.Status,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}

func (pr UnitKerjaRequest) TableName() string {
	return "unit_kerja"
}

func (pr UnitKerjaResponse) TableName() string {
	return "unit_kerja"
}
