package models

type AssetsRequest struct {
	ID              int64  `json:"id,string"`
	Type            string `json:"type"`
	KpknlID         int64  `json:"kpknl_id,string"`
	AuctionDate     string `json:"auction_date"`
	AuctionTime     string `json:"auction_time"`
	AuctionLink     string `json:"auction_link"`
	CategoryID      int64  `json:"category_id,string"`
	SubCategoryID   int64  `json:"subCategory_id,string"`
	Name            string `json:"name"`
	Price           int64  `json:"price"`
	Description     string `json:"description"`
	AddressID       int64  `json:"address_id"`
	BuildingAssetID int64  `json:"building_asset_id,string"`
	VehicleAssetID  int64  `json:"vehicle_sset_id,string"`
	ContactID       int64  `json:"contact_id,string"`
	ApprovalID      int64  `json:"approval_id,string"`
}

type AssetsResponse struct {
	ID              int64  `json:"id,string"`
	Type            string `json:"type"`
	KpknlID         int64  `json:"kpknl_id,string"`
	AuctionDate     string `json:"auction_date"`
	AuctionTime     string `json:"auction_time"`
	AuctionLink     string `json:"auction_link"`
	CategoryID      int64  `json:"category_id,string"`
	SubCategoryID   int64  `json:"subCategory_id,string"`
	Name            string `json:"name"`
	Price           int64  `json:"price"`
	Description     string `json:"description"`
	AddressID       int64  `json:"address_id"`
	BuildingAssetID int64  `json:"building_asset_id,string"`
	VehicleAssetID  int64  `json:"vehicle_sset_id,string"`
	ContactID       int64  `json:"contact_id,string"`
	ApprovalID      int64  `json:"approval_id,string"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

func (p AssetsRequest) ParseRequest() Assets {
	return Assets{
		ID:              p.ID,
		Type:            p.Type,
		KpknlID:         p.KpknlID,
		AuctionDate:     p.AuctionDate,
		AuctionTime:     p.AuctionTime,
		AuctionLink:     p.AuctionLink,
		CategoryID:      p.CategoryID,
		SubCategoryID:   p.SubCategoryID,
		Name:            p.Name,
		Price:           p.Price,
		Description:     p.Description,
		AddressID:       p.AddressID,
		BuildingAssetID: p.BuildingAssetID,
		VehicleAssetID:  p.VehicleAssetID,
		ContactID:       p.ContactID,
		ApprovalID:      p.ApprovalID,
	}
}

func (p AssetsResponse) ParseResponse() Assets {
	return Assets{
		ID:              p.ID,
		Type:            p.Type,
		KpknlID:         p.KpknlID,
		AuctionDate:     p.AuctionDate,
		AuctionTime:     p.AuctionTime,
		AuctionLink:     p.AuctionLink,
		CategoryID:      p.CategoryID,
		SubCategoryID:   p.SubCategoryID,
		Name:            p.Name,
		Price:           p.Price,
		Description:     p.Description,
		AddressID:       p.AddressID,
		BuildingAssetID: p.BuildingAssetID,
		VehicleAssetID:  p.VehicleAssetID,
		ContactID:       p.ContactID,
		ApprovalID:      p.ApprovalID,
		CreatedAt:       p.CreatedAt,
		UpdatedAt:       p.UpdatedAt,
	}
}
