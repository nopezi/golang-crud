package models

type Assets struct {
	ID            int64
	Type          string
	KpknlID       int64
	AuctionDate   string
	AuctionTime   string
	AuctionLink   string
	CategoryID    int64
	SubCategoryID int64
	Name          string
	Price         int64
	Description   string
	Status        string
	MakerID       string
	MakerDesc     string
	MakerDate     *string
	LastMakerID   string
	LastMakerDesc string
	LastMakerDate *string
	Published     bool
	PublishDate   *string
	Deleted       bool
	ExpiredDate   *string
	Action        string // create, updateApproval, updateMaintain, delete, publish, unpublish
	UpdatedAt     *string
	CreatedAt     *string
}

type AssetsUpdateApproval struct {
	ID          int64
	Published   bool
	PublishDate *string
	ExpiredDate *string
	Status      string
	Action      string
	UpdatedAt   *string
}

type AssetsUpdatePublish struct {
	ID            int64
	LastMakerID   string
	LastMakerDesc string
	LastMakerDate *string
	Published     bool
	PublishDate   *string
	ExpiredDate   *string
	Action        string
	UpdatedAt     *string
}
type AssetsUpdateDelete struct {
	ID            int64
	LastMakerID   string
	LastMakerDesc string
	LastMakerDate *string
	Deleted       bool
	Action        string
	UpdatedAt     *string
}

func (a AssetsUpdateApproval) TableName() string {
	return "assets"
}
