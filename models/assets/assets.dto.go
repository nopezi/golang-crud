package models

import (
	"infolelang/lib"
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
	ID             int64                          `json:"id"`
	FormType       string                         `json:"form_type"`
	Type           string                         `json:"type"`
	KpknlID        int64                          `json:"kpknl_id"`
	AuctionDate    string                         `json:"auction_date"`
	AuctionTime    string                         `json:"auction_time"`
	AuctionLink    string                         `json:"auction_link"`
	CategoryID     int64                          `json:"category_id"`
	SubCategoryID  int64                          `json:"sub_category_id"`
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
	ID             int64                          `json:"id"`
	FormType       string                         `json:"form_type"`
	Type           string                         `json:"type"`
	KpknlID        int64                          `json:"kpknl_id"`
	AuctionDate    string                         `json:"auction_date"`
	AuctionTime    string                         `json:"auction_time"`
	AuctionLink    string                         `json:"auction_link"`
	CategoryID     int64                          `json:"category_id"`
	SubCategoryID  int64                          `json:"sub_category_id"`
	Name           string                         `json:"name"`
	Price          int64                          `json:"price"`
	Description    string                         `json:"description"`
	Addresses      address.Addresses              `json:"addresses"`
	BuildingAssets building.BuildingAssets        `json:"building_assets"`
	VehicleAssets  vehicle.VehicleAssets          `json:"vehicle_assets"`
	Facilities     []facilities.FacilitiesRequest `json:"facilities"`
	AccessPlaces   []access.AccessPlacesRequest   `json:"access_places"`
	Contacts       contact.ContactsRequest        `json:"contacts"`
	Images         []images.ImagesRequest         `json:"images"`
	Approvals      approvals.Approvals            `json:"approvals"`
	UpdatedAt      *string                        `json:"updated_at"`
	CreatedAt      *string                        `json:"created_at"`
}

func (p AssetsRequest) ParseCreate(request AssetsRequest) *Assets {
	timeNow := lib.GetTimeNow("timestime")
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
		CreatedAt:     &timeNow,
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

func (a AssetsRequest) TableName() string {
	return "assets"
}

func (a AssetsResponse) TableName() string {
	return "assets"
}
