package models

type VerifikasiPICTindakLanjut struct {
	ID                    int64
	VerifikasiID          int64
	PICID                 *string
	TanggalTindakLanjut   string
	DeskripsiTindakLanjut int64
	CreatedAt             *string
	UpdatedAt             *string
}
