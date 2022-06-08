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
	UpdatedAt     *string
	CreatedAt     *string
}
