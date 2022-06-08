package images

type ImagesRequest struct {
	ID        int64  `json:"id"`
	Filename  string `json:"filename"`
	Path      string `json:"path"`
	Extension string `json:"extension"`
	Size      int64  `json:"size"`
}

type ImagesResponse struct {
	ID        int64   `json:"id"`
	Filename  string  `json:"filename"`
	Path      string  `json:"path"`
	Extension string  `json:"extension"`
	Size      int64   `json:"size"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

func (p ImagesRequest) ParseRequest() Images {
	return Images{
		ID:        p.ID,
		Filename:  p.Filename,
		Path:      p.Path,
		Extension: p.Extension,
		Size:      p.Size,
	}
}

func (p ImagesResponse) ParseResponse() Images {
	return Images{
		ID:        p.ID,
		Filename:  p.Filename,
		Path:      p.Path,
		Extension: p.Extension,
		Size:      p.Size,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (i ImagesRequest) TableName() string {
	return "images"
}

func (i ImagesResponse) TableName() string {
	return "images"
}
