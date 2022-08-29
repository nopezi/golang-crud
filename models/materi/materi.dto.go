package models

import (
	files "riskmanagement/models/files"
)

type MateriRequest struct {
	ID        int64                `json:"id"`
	Name      string               `json:"name"`
	CreatedAt string               `json:"created_at"`
	UpdatedAt string               `json:"updated_at"`
	MateriID  int64                `json:"materi_id"`
	FilesID   int64                `json:"files_id"`
	Files     []files.FilesRequest `json:"files"`
}

type MateriFilesRequest struct {
	FilesID int64 `json:"files_id"`
}

type MateriResponse struct {
	MateriID      int64                `json:"materi_id"`
	MateriFilesID int64                `json:"materi_files_id"`
	CreatedAt     *string              `json:"created_at"`
	UpdatedAt     *string              `json:"updated_at"`
	Files         []files.FilesRequest `json:"files"`
}

type MateriFilesResponse struct {
	MateriFilesID int64  `json:"materi_files_id"`
	MateriID      int64  `json:"materi_id"`
	FilesID       int64  `json:"files_id"`
	Filename      string `json:"filename"`
	Path          string `json:"path"`
	Extension     string `json:"extension"`
	Size          int64  `json:"size"`
}

type MateriAllResponse struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	UpdatedAt *string `json:"updated_at"`
	CreatedAt *string `json:"created_at"`
}

func (cr MateriAllResponse) TableName() string {
	return "materi"
}

func (cr Materi) TableName() string {
	return "materi"
}

func (cr MateriFiles) TableName() string {
	return "materi_files"
}
