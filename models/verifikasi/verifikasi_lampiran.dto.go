package models

import (
	files "riskmanagement/models/files"
)

type VerifikasiFilesRequest struct {
	ID           int64                `json:"id"`
	Name         string               `json:"name"`
	CreatedAt    string               `json:"created_at"`
	UpdatedAt    string               `json:"updated_at"`
	VerifikasiID int64                `json:"materi_id"`
	FilesID      int64                `json:"files_id"`
	Files        []files.FilesRequest `json:"files"`
}

type VerifikasiFilesResponse struct {
	ID           int64                  `json:"id"`
	Name         string                 `json:"name"`
	CreatedAt    string                 `json:"created_at"`
	UpdatedAt    string                 `json:"updated_at"`
	VerifikasiID int64                  `json:"materi_id"`
	FilesID      int64                  `json:"files_id"`
	Files        []files.FilesResponses `json:"files"`
}

func (cr VerifikasiFiles) TableName() string {
	return "verifikasi_lampiran"
}
