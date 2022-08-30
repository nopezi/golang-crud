package models

type Coaching struct {
	ID            int64
	NoPelaporan   string
	UnitKerja     string
	Peserta       string
	JumlahPeserta int64
	ActivityID    int64
	SubActivityID int64
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

type CoachingUpdateDelete struct {
	ID            int64
	LastMakerID   string
	LastMakerDesc string
	LastMakerDate *string
	Status        string
	Action        string // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted       bool
	UpdatedAt     *string
}

type CoachingUpdateActivity struct {
	ID            int64
	UnitKerja     string
	Peserta       string
	JumlahPeserta int64
	ActivityID    int64
	SubActivityID int64
	LastMakerID   string
	LastMakerDesc string
	LastMakerDate *string
	Deleted       bool
	Action        string
	Status        string
	UpdatedAt     *string
}

type CoachingActRequest struct {
	ID         int64 `json:"id"`
	CoachingID int64 `json:"coaching_id"`
}

func (c Coaching) TableName() string {
	return "coaching"
}

func (c CoachingUpdateDelete) TableName() string {
	return "coaching"
}

func (c CoachingUpdateActivity) TableName() string {
	return "coaching"
}
