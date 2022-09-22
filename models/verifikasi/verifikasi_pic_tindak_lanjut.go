package models

type VerifikasiPICTindakLanjut struct {
	ID                    int64
	VerifikasiID          int64
	PICID                 int64
	TanggalTindakLanjut   string
	DeskripsiTindakLanjut string
	Status                string
	// CreatedAt             *string
	// UpdatedAt             *string
}

func (vp VerifikasiPICTindakLanjut) TableName() string {
	return "verifikasi_pic_tindak_lanjut"
}
