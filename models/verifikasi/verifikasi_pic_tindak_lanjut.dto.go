package models

type VerifikasiPICTindakLanjutRequest struct {
	ID                    int64   `json:"id"`
	VerifikasiID          int64   `json:"verifikasi_id"`
	PICID                 *string `json:"pic_id"`
	TanggalTindakLanjut   string  `json:"tanggal_tindak_lanjut"`
	DeskripsiTindakLanjut int64   `json:"deskripsi_tindak_lanjut"`
	CreatedAt             *string `json:"created_at"`
	UpdatedAt             *string `json:"updated_at"`
}

type VerifikasiPICTindakLanjutResponse struct {
	ID                    int64   `json:"id"`
	VerifikasiID          int64   `json:"verifikasi_id"`
	PICID                 *string `json:"pic_id"`
	TanggalTindakLanjut   string  `json:"tanggal_tindak_lanjut"`
	DeskripsiTindakLanjut int64   `json:"deskripsi_tindak_lanjut"`
	CreatedAt             *string `json:"created_at"`
	UpdatedAt             *string `json:"updated_at"`
}

type VerifikasiPICTindakLanjutResponses struct {
	ID                    int64   `json:"id"`
	VerifikasiID          int64   `json:"verifikasi_id"`
	PICID                 *string `json:"pic_id"`
	TanggalTindakLanjut   string  `json:"tanggal_tindak_lanjut"`
	DeskripsiTindakLanjut int64   `json:"deskripsi_tindak_lanjut"`
}

func (vp VerifikasiPICTindakLanjutRequest) TableName() string {
	return "verifikasi_pic_tindak_lanjut"
}

func (vp VerifikasiPICTindakLanjutResponse) TableName() string {
	return "verifikasi_pic_tindak_lanjut"
}
