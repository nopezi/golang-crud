package models

type Briefing struct {
	ID            int64
	NoPelaporan   string
	UnitKerja     string
	Peserta       string
	JumlahPeserta int64
	MakerID       string
	MakerDesc     string
	MakerDate     *string
	LastMakerID   string
	LastMakerDesc string
	LastMakerDate *string
	Status        string
	Action        string // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted       bool
	UpdatedAt     *string
	CreatedAt     *string
}

type BriefingUpdateDelete struct {
	ID            int64
	LastMakerID   string
	LastMakerDesc string
	LastMakerDate *string
	Deleted       bool
	Action        string
	Status        string
	UpdatedAt     *string
}

type BriefingUpdateMateri struct {
	ID            int64
	UnitKerja     string
	Peserta       string
	JumlahPeserta int64
	LastMakerID   string
	LastMakerDesc string
	LastMakerDate *string
	Deleted       bool
	Action        string
	Status        string
	UpdatedAt     *string
}

type BriefMateriRequest struct {
	ID         int64 `json:"id"`
	BriefingID int64 `json:"briefing_id"`
}

func (b Briefing) TableName() string {
	return "briefing"
}

func (b BriefingUpdateDelete) TableName() string {
	return "briefing"
}

func (b BriefingUpdateMateri) TableName() string {
	return "briefing"
}
