package models

type VerifikasiFilesRequest struct {
	ID           int64 `json:"id"`
	VerifikasiID int64 `json:"verifikasi_id"`
	FilesID      int64 `json:"files_id"`
	// UpdatedAt    *string `json:"updated_at"`
	// CreatedAt    *string `json:"created_at"`
}

type VerifikasiFilesResponse struct {
	ID           int64 `json:"id"`
	VerifikasiID int64 `json:"verifikasi_id"`
	FilesID      int64 `json:"files_id"`
	// UpdatedAt    *string `json:"updated_at"`
	// CreatedAt    *string `json:"created_at"`
}

func (p VerifikasiFilesRequest) ParseRequest() VerifikasiFiles {
	return VerifikasiFiles{
		ID:           p.ID,
		VerifikasiID: p.VerifikasiID,
		FilesID:      p.FilesID,
	}
}

func (p VerifikasiFilesResponse) ParseResponse() VerifikasiFiles {
	return VerifikasiFiles{
		ID:           p.ID,
		VerifikasiID: p.VerifikasiID,
		FilesID:      p.FilesID,
		// UpdatedAt:    p.CreatedAt,
		// CreatedAt:    p.UpdatedAt,
	}
}

func (vf VerifikasiFilesRequest) TableName() string {
	return "verifikasi_lampiran"
}

func (vf VerifikasiFilesResponse) TableName() string {
	return "verifikasi_lampiran"
}
