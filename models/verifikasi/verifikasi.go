package models

type Verifikasi struct {
	ID                        int64
	NoPelaporan               string
	UnitKerja                 string
	ActivityID                int64
	SubActivityID             int64
	ProductID                 int64
	RiskIssueID               int64
	RiskIndicatorID           int64
	IncidentCauseID           int64
	SubIncidentCauseID        int64
	ApplicationID             int64
	HasilVerifikasi           string
	KunjunganNasabah          bool
	IndikasiFraud             bool
	JenisKerugianFinansial    bool
	JumlahPerkiraanKerugian   int64
	JenisKerugianNonFinansial string
	RekomendasiTindakLanjut   string
	RencanaTindakLanjut       string
	RiskTypeID                int64
	TanggalDitemukan          *string
	TanggalMulaiRTL           *string
	TanggalTargetSelesai      *string
	MakerID                   string
	MakerDesc                 string
	MakerDate                 *string
	LastMakerID               string
	LastMakerDesc             string
	LastMakerDate             *string
	Status                    string
	Action                    string // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted                   bool
	UpdatedAt                 *string
	CreatedAt                 *string
}

type VerifikasiUpdateDelete struct {
	ID            int64
	LastMakerID   string
	LastMakerDesc string
	LastMakerDate *string
	Status        string
	Action        string // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted       bool
	UpdatedAt     *string
}

type VerifikasiUpdateMaintain struct {
	ID            int64
	LastMakerID   string
	LastMakerDesc string
	LastMakerDate *string
	Status        string
	Action        string // create, updateApproval, updateMaintain, delete, publish, unpublish
	UpdatedAt     *string
}

type VerifyDataRequest struct {
	ID           int64 `json:"id"`
	VerifikasiID int64 `json:"verifikasi_id"`
}

func (v Verifikasi) TableName() string {
	return "verifikasi"
}

func (v VerifikasiUpdateDelete) TableName() string {
	return "verifikasi"
}
