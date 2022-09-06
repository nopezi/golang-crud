package models

import files "riskmanagement/models/files"

type VerifikasiRequest struct {
	ID                        int64                              `json:"id"`
	NoPelaporan               string                             `json:"no_pelaporan"`
	UnitKerja                 string                             `json:"unit_kerja"`
	ActivityID                int64                              `json:"activity_id"`
	SubActivityID             int64                              `json:"sub_activity_id"`
	ProductID                 int64                              `json:"product_id"`
	RiskIssueID               int64                              `json:"risk_issue_id"`
	RiskIndicatorID           int64                              `json:"risk_indicator_id"`
	IncidentCauseID           int64                              `json:"incident_cause_id"`
	SubIncidentCauseID        int64                              `json:"sub_incident_cause_id"`
	ApplicationID             int64                              `json:"application_id"`
	HasilVerifikasi           string                             `json:"hasil_verifikasi"`
	KunjunganNasabah          bool                               `json:"kunjungan_nasabah"`
	IndikasiFraud             bool                               `json:"indikasi_fraud"`
	JenisKerugianFinansial    bool                               `json:"jenis_kerugian_finansial"`
	JumlahPerkiraanKerugian   int64                              `json:"jumlah_perkiraan_kerugian"`
	JenisKerugianNonFinansial string                             `json:"jenis_kerugian_non_finansial"`
	RekomendasiTindakLanjut   string                             `json:"rekomendasi_tindak_lanjut"`
	RencanaTindakLanjut       string                             `json:"rencana_tindak_lanjut"`
	RiskTypeID                int64                              `json:"risk_type_id"`
	TanggalDitemukan          *string                            `json:"tanggal_ditemukan"`
	TanggalMulaiRTL           *string                            `json:"tanggal_mulai_rtl"`
	TanggalTargetSelesai      *string                            `json:"tanggal_target_selesai"`
	MakerID                   string                             `json:"maker_id"`
	MakerDesc                 string                             `json:"maker_desc"`
	MakerDate                 *string                            `json:"maker_date"`
	LastMakerID               string                             `json:"last_maker_id"`
	LastMakerDesc             string                             `json:"last_maker_desc"`
	LastMakerDate             *string                            `json:"last_maker_date"`
	Status                    string                             `json:"status"`
	Action                    string                             `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted                   bool                               `json:"deleted"`
	DataAnomali               []VerifikasiAnomaliDataRequest     `json:"data_anomali"`
	PICTindakLanjut           []VerifikasiPICTindakLanjutRequest `json:"pic_tindak_lanjut"`
	Files                     []files.FilesRequest               `json:"files"`
	UpdatedAt                 *string                            `json:"updated_at"`
	CreatedAt                 *string                            `json:"created_at"`
}

type VerifikasiRequestUpdateMaintain struct {
	ID                        int64   `json:"id"`
	NoPelaporan               string  `json:"no_pelaporan"`
	UnitKerja                 string  `json:"unit_kerja"`
	ActivityID                int64   `json:"activity_id"`
	SubActivityID             int64   `json:"sub_activity_id"`
	ProductID                 int64   `json:"product_id"`
	RiskIssueID               int64   `json:"risk_issue_id"`
	RiskIndicatorID           int64   `json:"risk_indicator_id"`
	IncidentCauseID           int64   `json:"incident_cause_id"`
	SubIncidentCauseID        int64   `json:"sub_incident_cause_id"`
	ApplicationID             int64   `json:"application_id"`
	HasilVerifikasi           string  `json:"hasil_verifikasi"`
	KunjunganNasabah          bool    `json:"kunjungan_nasabah"`
	IndikasiFraud             bool    `json:"indikasi_fraud"`
	JenisKerugianFinansial    bool    `json:"jenis_kerugian_finansial"`
	JumlahPerkiraanKerugian   int64   `json:"jumlah_perkiraan_kerugian"`
	JenisKerugianNonFinansial string  `json:"jenis_kerugian_non_finansial"`
	RekomendasiTindakLanjut   string  `json:"rekomendasi_tindak_lanjut"`
	RencanaTindakLanjut       string  `json:"rencana_tindak_lanjut"`
	RiskTypeID                int64   `json:"risk_type_id"`
	TanggalDitemukan          *string `json:"tanggal_ditemukan"`
	TanggalMulaiRTL           *string `json:"tanggal_mulai_rtl"`
	TanggalTargetSelesai      *string `json:"tanggal_target_selesai"`
	MakerID                   string  `json:"maker_id"`
	MakerDesc                 string  `json:"maker_desc"`
	MakerDate                 *string `json:"maker_date"`
	LastMakerID               string  `json:"last_maker_id"`
	LastMakerDesc             string  `json:"last_maker_desc"`
	LastMakerDate             *string `json:"last_maker_date"`
	Status                    string  `json:"status"`
	Action                    string  `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted                   bool    `json:"deleted"`
	UpdatedAt                 *string `json:"updated_at"`
}

type VerifikasiResponse struct {
	ID                        int64   `json:"id"`
	NoPelaporan               string  `json:"no_pelaporan"`
	UnitKerja                 string  `json:"unit_kerja"`
	ActivityID                int64   `json:"activity_id"`
	SubActivityID             int64   `json:"sub_activity_id"`
	ProductID                 int64   `json:"product_id"`
	RiskIssueID               int64   `json:"risk_issue_id"`
	RiskIndicatorID           int64   `json:"risk_indicator_id"`
	IncidentCauseID           int64   `json:"incident_cause_id"`
	SubIncidentCauseID        int64   `json:"sub_incident_cause_id"`
	ApplicationID             int64   `json:"application_id"`
	HasilVerifikasi           string  `json:"hasil_verifikasi"`
	KunjunganNasabah          bool    `json:"kunjungan_nasabah"`
	IndikasiFraud             bool    `json:"indikasi_fraud"`
	JenisKerugianFinansial    bool    `json:"jenis_kerugian_finansial"`
	JumlahPerkiraanKerugian   int64   `json:"jumlah_perkiraan_kerugian"`
	JenisKerugianNonFinansial string  `json:"jenis_kerugian_non_finansial"`
	RekomendasiTindakLanjut   string  `json:"rekomendasi_tindak_lanjut"`
	RencanaTindakLanjut       string  `json:"rencana_tindak_lanjut"`
	RiskTypeID                int64   `json:"risk_type_id"`
	TanggalDitemukan          *string `json:"tanggal_ditemukan"`
	TanggalMulaiRTL           *string `json:"tanggal_mulai_rtl"`
	TanggalTargetSelesai      *string `json:"tanggal_target_selesai"`
	MakerID                   string  `json:"maker_id"`
	MakerDesc                 string  `json:"maker_desc"`
	MakerDate                 *string `json:"maker_date"`
	LastMakerID               string  `json:"last_maker_id"`
	LastMakerDesc             string  `json:"last_maker_desc"`
	LastMakerDate             *string `json:"last_maker_date"`
	Status                    string  `json:"status"`
	Action                    string  `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted                   bool    `json:"deleted"`
	UpdatedAt                 *string `json:"updated_at"`
	CreatedAt                 *string `json:"created_at"`
}

type VerifikasiResponseGetOne struct {
	ID                        int64                                `json:"id"`
	NoPelaporan               string                               `json:"no_pelaporan"`
	UnitKerja                 string                               `json:"unit_kerja"`
	ActivityID                int64                                `json:"activity_id"`
	SubActivityID             int64                                `json:"sub_activity_id"`
	ProductID                 int64                                `json:"product_id"`
	RiskIssueID               int64                                `json:"risk_issue_id"`
	RiskIndicatorID           int64                                `json:"risk_indicator_id"`
	IncidentCauseID           int64                                `json:"incident_cause_id"`
	SubIncidentCauseID        int64                                `json:"sub_incident_cause_id"`
	ApplicationID             int64                                `json:"application_id"`
	HasilVerifikasi           string                               `json:"hasil_verifikasi"`
	KunjunganNasabah          bool                                 `json:"kunjungan_nasabah"`
	IndikasiFraud             bool                                 `json:"indikasi_fraud"`
	JenisKerugianFinansial    bool                                 `json:"jenis_kerugian_finansial"`
	JumlahPerkiraanKerugian   int64                                `json:"jumlah_perkiraan_kerugian"`
	JenisKerugianNonFinansial string                               `json:"jenis_kerugian_non_finansial"`
	RekomendasiTindakLanjut   string                               `json:"rekomendasi_tindak_lanjut"`
	RencanaTindakLanjut       string                               `json:"rencana_tindak_lanjut"`
	RiskTypeID                int64                                `json:"risk_type_id"`
	TanggalDitemukan          *string                              `json:"tanggal_ditemukan"`
	TanggalMulaiRTL           *string                              `json:"tanggal_mulai_rtl"`
	TanggalTargetSelesai      *string                              `json:"tanggal_target_selesai"`
	MakerID                   string                               `json:"maker_id"`
	MakerDesc                 string                               `json:"maker_desc"`
	MakerDate                 *string                              `json:"maker_date"`
	LastMakerID               string                               `json:"last_maker_id"`
	LastMakerDesc             string                               `json:"last_maker_desc"`
	LastMakerDate             *string                              `json:"last_maker_date"`
	Status                    string                               `json:"status"`
	Action                    string                               `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted                   bool                                 `json:"deleted"`
	DataAnomali               []VerifikasiAnomaliDataResponses     `json:"data_anomali"`
	PICTindakLanjut           []VerifikasiPICTindakLanjutResponses `json:"pic_tindak_lanjut"`
	Files                     []VerifikasiFilesResponses           `json:"files"` // Files                     []files.FilesResponses               `json:"files"`
	UpdatedAt                 *string                              `json:"updated_at"`
	CreatedAt                 *string                              `json:"created_at"`
}

type VerifikasiRequestMaintain struct {
	ID                        int64                              `json:"id"`
	NoPelaporan               string                             `json:"no_pelaporan"`
	UnitKerja                 string                             `json:"unit_kerja"`
	ActivityID                int64                              `json:"activity_id"`
	SubActivityID             int64                              `json:"sub_activity_id"`
	ProductID                 int64                              `json:"product_id"`
	RiskIssueID               int64                              `json:"risk_issue_id"`
	RiskIndicatorID           int64                              `json:"risk_indicator_id"`
	IncidentCauseID           int64                              `json:"incident_cause_id"`
	SubIncidentCauseID        int64                              `json:"sub_incident_cause_id"`
	ApplicationID             int64                              `json:"application_id"`
	HasilVerifikasi           string                             `json:"hasil_verifikasi"`
	KunjunganNasabah          bool                               `json:"kunjungan_nasabah"`
	IndikasiFraud             bool                               `json:"indikasi_fraud"`
	JenisKerugianFinansial    bool                               `json:"jenis_kerugian_finansial"`
	JumlahPerkiraanKerugian   int64                              `json:"jumlah_perkiraan_kerugian"`
	JenisKerugianNonFinansial string                             `json:"jenis_kerugian_non_finansial"`
	RekomendasiTindakLanjut   string                             `json:"rekomendasi_tindak_lanjut"`
	RencanaTindakLanjut       string                             `json:"rencana_tindak_lanjut"`
	RiskTypeID                int64                              `json:"risk_type_id"`
	TanggalDitemukan          *string                            `json:"tanggal_ditemukan"`
	TanggalMulaiRTL           *string                            `json:"tanggal_mulai_rtl"`
	TanggalTargetSelesai      *string                            `json:"tanggal_target_selesai"`
	MakerID                   string                             `json:"maker_id"`
	MakerDesc                 string                             `json:"maker_desc"`
	MakerDate                 *string                            `json:"maker_date"`
	LastMakerID               string                             `json:"last_maker_id"`
	LastMakerDesc             string                             `json:"last_maker_desc"`
	LastMakerDate             *string                            `json:"last_maker_date"`
	Status                    string                             `json:"status"`
	Action                    string                             `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted                   bool                               `json:"deleted"`
	DataAnomali               []VerifikasiAnomaliDataRequest     `json:"data_anomali"`
	PICTindakLanjut           []VerifikasiPICTindakLanjutRequest `json:"pic_tindak_lanjut"`
	Files                     []files.FilesRequest               `json:"files"`
	UpdatedAt                 *string                            `json:"updated_at"`
	CreatedAt                 *string                            `json:"created_at"`
}

type VerifikasiFileRequest struct {
	FilesID              int64  `json:"files_id"`
	VerifikasiLampiranID int64  `json:"verifikasi_lampiran_id"`
	Path                 string `json:"path"`
}

func (v VerifikasiRequest) TableName() string {
	return "verifikasi"
}

func (v VerifikasiResponse) TableName() string {
	return "verifikasi"
}

func (v VerifikasiRequestUpdateMaintain) TableName() string {
	return "verifikasi"
}
