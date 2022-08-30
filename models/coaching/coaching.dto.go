package models

type CoachingRequest struct {
	ID            int64                     `json:"id"`
	NoPelaporan   string                    `json:"no_pelaporan"`
	UnitKerja     string                    `json:"unit_kerja"`
	Peserta       string                    `json:"peserta"`
	JumlahPeserta int64                     `json:"jumlah_peserta"`
	ActivityID    int64                     `json:"activity_id"`
	SubActivityID int64                     `json:"sub_activity_id"`
	MakerID       string                    `json:"maker_id"`
	MakerDesc     string                    `json:"maker_desc"`
	MakerDate     *string                   `json:"maker_date"`
	LastMakerID   string                    `json:"last_maker_id"`
	LastMakerDesc string                    `json:"last_maker_desc"`
	LastMakerDate *string                   `json:"last_maker_Date"`
	Status        string                    `json:"status"`
	Action        string                    `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted       bool                      `json:"deleted"`
	Activity      []CoachingActivityRequest `json:"activity"`
}

type CoachingResponse struct {
	ID            int64   `json:"id"`
	NoPelaporan   string  `json:"no_pelaporan"`
	UnitKerja     string  `json:"unit_kerja"`
	Peserta       string  `json:"peserta"`
	JumlahPeserta int64   `json:"jumlah_peserta"`
	ActivityID    int64   `json:"activity_id"`
	SubActivityID int64   `json:"sub_activity_id"`
	MakerID       string  `json:"maker_id"`
	MakerDesc     string  `json:"maker_desc"`
	MakerDate     *string `json:"maker_date"`
	LastMakerID   string  `json:"last_maker_id"`
	LastMakerDesc string  `json:"last_maker_desc"`
	LastMakerDate *string `json:"last_maker_Date"`
	Status        string  `json:"status"`
	Action        string  `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted       bool    `json:"deleted"`
	UpdatedAt     *string `json:"updated_at"`
	CreatedAt     *string `json:"created_at"`
}

type CoachingResponses struct {
	ID            int64   `json:"id"`
	NoPelaporan   string  `json:"no_pelaporan"`
	UnitKerja     string  `json:"unit_kerja"`
	Peserta       string  `json:"peserta"`
	JumlahPeserta int64   `json:"jumlah_peserta"`
	ActivityID    int64   `json:"activity_id"`
	SubActivityID int64   `json:"sub_activity_id"`
	MakerID       string  `json:"maker_id"`
	MakerDesc     string  `json:"maker_desc"`
	MakerDate     *string `json:"maker_date"`
	LastMakerID   string  `json:"last_maker_id"`
	LastMakerDesc string  `json:"last_maker_desc"`
	LastMakerDate *string `json:"last_maker_Date"`
	Status        string  `json:"status"`
	Action        string  `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted       bool    `json:"deleted"`
}

type CoachingResponsesGetOneString struct {
	ID            int64                       `json:"id"`
	NoPelaporan   string                      `json:"no_pelaporan"`
	UnitKerja     string                      `json:"unit_kerja"`
	Peserta       string                      `json:"peserta"`
	JumlahPeserta int64                       `json:"jumlah_peserta"`
	ActivityID    int64                       `json:"activity_id"`
	SubActivityID int64                       `json:"sub_activity_id"`
	MakerID       string                      `json:"maker_id"`
	MakerDesc     string                      `json:"maker_desc"`
	MakerDate     *string                     `json:"maker_date"`
	LastMakerID   string                      `json:"last_maker_id"`
	LastMakerDesc string                      `json:"last_maker_desc"`
	LastMakerDate *string                     `json:"last_maker_Date"`
	Status        string                      `json:"status"`
	Action        string                      `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted       bool                        `json:"deleted"`
	Activity      []CoachingActivityResponses `json:"activity"`
	UpdatedAt     *string                     `json:"updated_at"`
	CreatedAt     *string                     `json:"created_at"`
}

type CoachingRequestUpdate struct {
	ID            int64   `json:"id"`
	LastMakerID   string  `json:"last_maker_id"`
	LastMakerDesc string  `json:"last_maker_desc"`
	LastMakerDate *string `json:"last_maker_Date"`
}

type CoachingResponseMaintain struct {
	ID            int64                      `json:"id"`
	NoPelaporan   string                     `json:"no_pelaporan"`
	UnitKerja     string                     `json:"unit_kerja"`
	Peserta       string                     `json:"peserta"`
	JumlahPeserta int64                      `json:"jumlah_peserta"`
	ActivityID    int64                      `json:"activity_id"`
	SubActivityID int64                      `json:"sub_activity_id"`
	LastMakerID   string                     `json:"last_maker_id"`
	LastMakerDesc string                     `json:"last_maker_desc"`
	LastMakerDate *string                    `json:"last_maker_Date"`
	Status        string                     `json:"status"`
	Action        string                     `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted       bool                       `json:"deleted"`
	Activity      []CoachingActivityResponse `json:"activity"`
	UpdatedAt     *string                    `json:"updated_at"`
}

func (c CoachingRequest) TableName() string {
	return "coaching"
}

func (c CoachingResponse) TableName() string {
	return "coaching"
}
