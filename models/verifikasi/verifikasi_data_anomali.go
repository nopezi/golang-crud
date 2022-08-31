package models

type VerifikasiAnomaliData struct {
	ID              int64
	VerifikasiID    int64
	TanggalKejadian *string
	NomorRekening   string
	Nominal         int64
	Keterangan      string
	CreatedAt       *string
	UpdatedAt       *string
}
