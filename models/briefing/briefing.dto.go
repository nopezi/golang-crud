package models

type BriefingRequest struct {
	ID            int64                   `json:"id"`
	NoPelaporan   string                  `json:"no_pelaporan"`
	UnitKerja     string                  `json:"unit_kerja"`
	Peserta       string                  `json:"peserta"`
	JumlahPeserta int64                   `json:"jumlah_peserta"`
	MakerID       string                  `json:"maker_id"`
	MakerDesc     string                  `json:"maker_desc"`
	MakerDate     *string                 `json:"maker_date"`
	LastMakerID   string                  `json:"last_maker_id"`
	LastMakerDesc string                  `json:"last_maker_desc"`
	LastMakerDate *string                 `json:"last_maker_date"`
	Status        string                  `json:"status"`
	Action        string                  `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted       bool                    `json:"deleted"`
	Materi        []BriefingMateriRequest `json:"materi"`
}

type BriefingResponse struct {
	ID            int64   `json:"id"`
	NoPelaporan   string  `json:"no_pelaporan"`
	UnitKerja     string  `json:"unit_kerja"`
	Peserta       string  `json:"peserta"`
	JumlahPeserta int64   `json:"jumlah_peserta"`
	MakerID       string  `json:"maker_id"`
	MakerDesc     string  `json:"maker_desc"`
	MakerDate     *string `json:"maker_date"`
	LastMakerID   string  `json:"last_maker_id"`
	LastMakerDesc string  `json:"last_maker_desc"`
	LastMakerDate *string `json:"last_maker_date"`
	Status        string  `json:"status"`
	Action        string  `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted       bool    `json:"deleted"`
	CreatedAt     *string `json:"created_at"`
	UpdatedAt     *string `json:"updated_at"`
}

type BriefingResponses struct {
	ID            int64   `json:"id"`
	NoPelaporan   string  `json:"no_pelaporan"`
	UnitKerja     string  `json:"unit_kerja"`
	Peserta       string  `json:"peserta"`
	JumlahPeserta int64   `json:"jumlah_peserta"`
	MakerID       string  `json:"maker_id"`
	MakerDesc     string  `json:"maker_desc"`
	MakerDate     *string `json:"maker_date"`
	LastMakerID   string  `json:"last_maker_id"`
	LastMakerDesc string  `json:"last_maker_desc"`
	LastMakerDate *string `json:"last_maker_date"`
	Status        string  `json:"status"`
	Action        string  `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted       bool    `json:"deleted"`
}

type BriefingResponseData struct {
	ID          int64  `json:"id"`
	NoPelaporan string `json:"no_pelaporan"`
	UnitKerja   string `json:"unit_kerja"`
	JudulMateri string `json:"judul_materi"`
	Aktifitas   string `json:"aktifitas"`
	StatusBrf   string `json:"status_brf"`
}

type BriefingResponseGetOneString struct {
	ID            int64                     `json:"id"`
	NoPelaporan   string                    `json:"no_pelaporan"`
	UnitKerja     string                    `json:"unit_kerja"`
	Peserta       string                    `json:"peserta"`
	JumlahPeserta int64                     `json:"jumlah_peserta"`
	MakerID       string                    `json:"maker_id"`
	MakerDesc     string                    `json:"maker_desc"`
	MakerDate     *string                   `json:"maker_date"`
	LastMakerID   string                    `json:"last_maker_id"`
	LastMakerDesc string                    `json:"last_maker_desc"`
	LastMakerDate *string                   `json:"last_maker_date"`
	Status        string                    `json:"status"`
	Action        string                    `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted       bool                      `json:"deleted"`
	Materi        []BriefingMateriResponses `json:"materi"`
	CreatedAt     *string                   `json:"created_at"`
	UpdatedAt     *string                   `json:"updated_at"`
}

type BriefingRequestUpdate struct {
	ID            int64   `json:"id"`
	LastMakerID   string  `json:"last_maker_id"`
	LastMakerDesc string  `json:"last_maker_desc"`
	LastMakerDate *string `json:"last_maker_date"`
}

type BriefingResponseMaintain struct {
	ID            int64                    `json:"id"`
	NoPelaporan   string                   `json:"no_pelaporan"`
	UnitKerja     string                   `json:"unit_kerja"`
	Peserta       string                   `json:"peserta"`
	JumlahPeserta int64                    `json:"jumlah_peserta"`
	LastMakerID   string                   `json:"last_maker_id"`
	LastMakerDesc string                   `json:"last_maker_desc"`
	LastMakerDate *string                  `json:"last_maker_date"`
	Status        string                   `json:"status"`
	Action        string                   `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted       bool                     `json:"deleted"`
	Materi        []BriefingMateriResponse `json:"materi"`
	UpdatedAt     *string                  `json:"updated_at"`
}

func (b BriefingRequest) TableName() string {
	return "briefing"
}

func (b BriefingResponse) TableName() string {
	return "briefing"
}