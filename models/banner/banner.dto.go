package models

import (
	images "infolelang/models/images"
)

type BannerRequest struct {
	ID         int64                  `json:"id"`
	BannerName string                 `json:"banner_name"`
	BannerID   int64                  `json:"banner_id"`
	ImageID    int64                  `json:"image_id"`
	Images     []images.ImagesRequest `json:"images"`
}

type BannerImageRequest struct {
	ImageID int64 `json:"image_id"`
}

type BannerResponse struct {
	BannerID      int64                  `json:"banner_id"`
	BannerImageID int64                  `json:"banner_image_id"`
	CreatedAt     *string                `json:"created_at"`
	UpdatedAt     *string                `json:"updated_at"`
	Images        []images.ImagesRequest `json:"images"`
}

type BannerImageResponse struct {
	BannerImageID int64  `json:"banner_image_id"`
	BannerID      int64  `json:"banner_id"`
	ImageID       int64  `json:"image_id"`
	Filename      string `json:"filename"`
	Path          string `json:"path"`
	Extension     string `json:"extension"`
	Size          int64  `json:"size"`
}

func (cr Banner) TableName() string {
	return "banners"
}

func (cr BannerImage) TableName() string {
	return "banner_images"
}
