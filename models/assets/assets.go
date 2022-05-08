package models

type Assets struct {
	ID              int64
	Type            string
	KpknlID         int64
	AuctionDate     string
	AuctionTime     string
	AuctionLink     string
	CategoryID      int64
	SubCategoryID   int64
	Name            string
	Price           int64
	Description     string
	AddressID       int64
	BuildingAssetID int64
	VehicleAssetID  int64
	ContactID       int64
	ApprovalID      int64
	UpdatedAt       string
	CreatedAt       string
}
