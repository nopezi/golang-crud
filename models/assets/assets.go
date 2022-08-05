package models

type Assets struct {
	ID            int64
	FormType      string
	Type          string
	KpknlID       int64
	AuctionDate   *string
	AuctionTime   *string
	AuctionLink   string
	CategoryID    int64
	SubCategoryID int64
	Name          string
	Price         float32
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
	DocumentID    string
	UpdatedAt     *string
	CreatedAt     *string
}

type AssetsUpdateApproval struct {
	ID            int64
	LastMakerID   string
	LastMakerDesc string
	LastMakerDate *string
	Published     bool
	PublishDate   *string
	ExpiredDate   *string
	Status        string
	Action        string
	UpdatedAt     *string
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
	Status        string
	UpdatedAt     *string
}
type AssetsUpdateDelete struct {
	ID            int64
	LastMakerID   string
	LastMakerDesc string
	LastMakerDate *string
	Deleted       bool
	Published     bool
	PublishDate   *string
	ExpiredDate   *string
	Action        string
	Status        string
	UpdatedAt     *string
}

func (a Assets) TableName() string {
	return "assets"
}
func (a AssetsUpdateApproval) TableName() string {
	return "assets"
}

func (a AssetsUpdatePublish) TableName() string {
	return "assets"
}

func (a AssetsUpdateDelete) TableName() string {
	return "assets"
}
