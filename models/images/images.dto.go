package models

type ImagesRequest struct {
	ID   int64  `json:"id,string"`
	Path string `json:"path"`
	Size int64  `json:"size,string"`
}

type ImagesResponse struct {
	ID        int64  `json:"id,string"`
	Path      string `json:"path"`
	Size      int64  `json:"size,string"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (p ImagesRequest) ParseRequest() Images {
	return Images{
		ID:   p.ID,
		Path: p.Path,
		Size: p.Size,
	}
}

func (p ImagesResponse) ParseResponse() Images {
	return Images{
		ID:        p.ID,
		Path:      p.Path,
		Size:      p.Size,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
