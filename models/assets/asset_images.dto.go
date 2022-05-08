package models

type AssetImagesRequest struct {
	ID      int64 `json:"id,string"`
	AssetID int64 `json:"asset_id,string"`
	ImageID int64 `json:"image_id,string"`
}

type AssetImagesResponse struct {
	ID        int64  `json:"id,string"`
	AssetID   int64  `json:"asset_id,string"`
	ImageID   int64  `json:"image_id,string"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (p AssetImagesRequest) ParseRequest() AssetImages {
	return AssetImages{
		ID:      p.ID,
		AssetID: p.AssetID,
		ImageID: p.ImageID,
	}
}

func (p AssetImagesResponse) ParseResponse() AssetImages {
	return AssetImages{
		ID:        p.ID,
		AssetID:   p.AssetID,
		ImageID:   p.ImageID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
