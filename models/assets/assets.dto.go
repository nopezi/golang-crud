package models

import (
	access "infolelang/models/access_places"
	address "infolelang/models/addresses"
	approvals "infolelang/models/approvals"
	building "infolelang/models/building_assets"
	contact "infolelang/models/contacts"
	facilities "infolelang/models/facilities"
	images "infolelang/models/images"
	vehicle "infolelang/models/vehicle_assets"
)

type AssetsRequest struct {
	ID             int64                          `json:"id,string"`
	FormType       string                         `json:"formType"`
	Type           string                         `json:"type"`
	KpknlID        int64                          `json:"kpknl_id,string"`
	AuctionDate    string                         `json:"auction_date"`
	AuctionTime    string                         `json:"auction_time"`
	AuctionLink    string                         `json:"auction_link"`
	CategoryID     int64                          `json:"category_id,string"`
	SubCategoryID  int64                          `json:"subCategory_id,string"`
	Name           string                         `json:"name"`
	Price          int64                          `json:"price"`
	Description    string                         `json:"description"`
	Addresses      address.AddressesRequest       `json:"addresses"`
	BuildingAssets building.BuildingAssetsRequest `json:"building_assets"`
	VehicleAssets  vehicle.VehicleAssetsRequest   `json:"vehicle_assets"`
	Facilities     []facilities.FacilitiesRequest `json:"facilities"`
	AccessPlaces   []access.AccessPlacesRequest   `json:"access_places"`
	Contacts       contact.ContactsRequest        `json:"contacts"`
	Images         []images.ImagesRequest         `json:"images"`
	Approvals      approvals.ApprovalsRequest     `json:"approvals"`
}

type AssetsResponse struct {
	ID            int64  `json:"id,string"`
	Type          string `json:"type"`
	KpknlID       int64  `json:"kpknl_id,string"`
	AuctionDate   string `json:"auction_date"`
	AuctionTime   string `json:"auction_time"`
	AuctionLink   string `json:"auction_link"`
	CategoryID    int64  `json:"category_id,string"`
	SubCategoryID int64  `json:"subCategory_id,string"`
	Name          string `json:"name"`
	Price         int64  `json:"price"`
	Description   string `json:"description"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

func (p AssetsRequest) ParseRequest(request AssetsRequest) *Assets {
	return &Assets{
		ID:            request.ID,
		Type:          request.Type,
		KpknlID:       request.KpknlID,
		AuctionDate:   request.AuctionDate,
		AuctionTime:   request.AuctionTime,
		AuctionLink:   request.AuctionLink,
		CategoryID:    request.CategoryID,
		SubCategoryID: request.SubCategoryID,
		Name:          request.Name,
		Price:         request.Price,
		Description:   request.Description,
	}
}

func (p AssetsResponse) ParseResponse() Assets {
	return Assets{
		ID:            p.ID,
		Type:          p.Type,
		KpknlID:       p.KpknlID,
		AuctionDate:   p.AuctionDate,
		AuctionTime:   p.AuctionTime,
		AuctionLink:   p.AuctionLink,
		CategoryID:    p.CategoryID,
		SubCategoryID: p.SubCategoryID,
		Name:          p.Name,
		Price:         p.Price,
		Description:   p.Description,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

func (ar AssetsRequest) TableName() string {
	return "assets"
}
