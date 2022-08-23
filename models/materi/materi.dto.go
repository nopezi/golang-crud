package materi

type MateriRequest struct {
	ID        int64   `json:"id"`
	Filename  string  `json:"filename"`
	Path      string  `json:"path"`
	Extension string  `json:"extension"`
	Size      int64   `json:"size"`
	UpdatedAt *string `json:"updated_at"`
	CreatedAt *string `json:"created_at"`
}

type MateriResponses struct {
	ID        int64   `json:"id"`
	Filename  string  `json:"filename"`
	Path      string  `json:"path"`
	Extension string  `json:"extension"`
	Size      int64   `json:"size"`
	UpdatedAt *string `json:"updated_at"`
	CreatedAt *string `json:"created_at"`
}

func (p MateriRequest) ParseRequest() Materi {
	return Materi{
		ID:        p.ID,
		Filename:  p.Filename,
		Path:      p.Path,
		Extension: p.Extension,
		Size:      p.Size,
	}
}

func (p MateriResponses) ParseRequest() Materi {
	return Materi{
		ID:        p.ID,
		Filename:  p.Filename,
		Path:      p.Path,
		Extension: p.Extension,
		Size:      p.Size,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (m MateriRequest) TableName() string {
	return "rekomendasi_materi"
}

func (m MateriResponses) TableName() string {
	return "rekomendasi_materi"
}
