package models

type Banner struct {
	ID        int64
	Name      string
	UpdatedAt *string
	CreatedAt *string
}

type BannerImage struct {
	ID       int64
	BannerID int64
	ImageID  int64
}
