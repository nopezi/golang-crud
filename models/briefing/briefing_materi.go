package models

type BriefingMateri struct {
	ID                int64
	BriefingID        int64
	ActivityID        int64
	SubActivityID     int64
	ProductID         int64
	JudulMateri       string
	RekomendasiMateri string
	MateriTambahan    string
	UpdatedAt         *string
	CreatedAt         *string
}
