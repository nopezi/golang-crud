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
	MakerComment  string
	MakerDate     *string
	LastMakerID   string
	LastMakerDesc string
	LastMakerDate *string
	Published     bool
	Deleted       bool
	ExpiredDate   *string
	Action        string // create, updateApproval, updateMaintain, delete, publish, unpublish
	UpdatedAt     *string
	CreatedAt     *string
}
