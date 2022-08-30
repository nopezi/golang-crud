package models

type CoachingActivityRequest struct {
	ID                int64   `json:"id"`
	CoachingID        int64   `json:"coaching_id"`
	RiskIssueID       int64   `json:"risk_issue_id"`
	JudulMateri       string  `json:"judul_materi"`
	RekomendasiMateri string  `json:"rekomendasi_materi"`
	MateriTambahan    string  `json:"materi_tambahan"`
	UpdatedAt         *string `json:"updated_at"`
	CreatedAt         *string `json:"created_at"`
}

type CoachingActivityResponse struct {
	ID                int64   `json:"id"`
	CoachingID        int64   `json:"coaching_id"`
	RiskIssueID       int64   `json:"risk_issue_id"`
	JudulMateri       string  `json:"judul_materi"`
	RekomendasiMateri string  `json:"rekomendasi_materi"`
	MateriTambahan    string  `json:"materi_tambahan"`
	UpdatedAt         *string `json:"updated_at"`
	CreatedAt         *string `json:"created_at"`
}

type CoachingActivityResponses struct {
	ID                int64  `json:"id"`
	CoachingID        int64  `json:"coaching_id"`
	RiskIssueID       int64  `json:"risk_issue_id"`
	JudulMateri       string `json:"judul_materi"`
	RekomendasiMateri string `json:"rekomendasi_materi"`
	MateriTambahan    string `json:"materi_tambahan"`
}

func (ca CoachingActivityRequest) TableName() string {
	return "coaching_activity"
}

func (ca CoachingActivityResponse) TableName() string {
	return "coaching_activity"
}
