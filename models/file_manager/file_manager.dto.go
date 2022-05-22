package file_manager

import "mime/multipart"

type FileManagerResponse struct {
	Subdir   string `json:"subdir"`
	Size     string `json:"size"`
	Filename string `json:"filename"`
}

type FileManagerRequest struct {
	File       *multipart.FileHeader `json:"file"`
	BucketName string                `json:"bucketName"`
	ObjectName string                `json:"objectName"`
	Subdir     string                `json:"subdir"`
}
