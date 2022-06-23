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
	Status         string                         `json:"status"`
	MakerID        string                         `json:"maker_id"`
	MakerDesc      string                         `json:"maker_desc"`
	MakerDate      *string                        `json:"maker_date"`
	LastMakerID    string                         `json:"last_maker_id"`
	LastMakerDesc  string                         `json:"last_maker_desc"`
	LastMakerDate  *string                        `json:"last_maker_date"`
	Published      bool                           `json:"published"`
	Deleted        bool                           `json:"deleted"`
	ExpiredDate    *string                        `json:"expired_date"`
	Action         string                         `json:"action"`
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
	// FormType      string  `json:"form_type"`
	ID              int64   `json:"id"`
	Type            string  `json:"type"`
	KpknlID         int64   `json:"kpknl_id"`
	AuctionDate     string  `json:"auction_date"`
	AuctionTime     string  `json:"auction_time"`
	AuctionLink     string  `json:"auction_link"`
	CategoryID      int64   `json:"category_id"`
	SubCategoryID   int64   `json:"sub_category_id"`
	Name            string  `json:"name"`
	Price           int64   `json:"price"`
	Description     string  `json:"description"`
	MakerID         string  `json:"maker_id"`
	MakerDesc       string  `json:"maker_desc"`
	MakerDate       *string `json:"maker_date"`
	LastMakerID     string  `json:"last_maker_id"`
	LastMakerDesc   string  `json:"last_maker_desc"`
	LastMakerDate   *string `json:"last_maker_date"`
	Published       bool    `json:"published"`
	Deleted         bool    `json:"deleted"`
	PublishDate     *string `json:"publish_date"`
	ExpiredDate     *string `json:"expired_date"`
	Status          string  `json:"status"`
	Action          string  `json:"action"`
	UpdatedAt       *string `json:"updated_at"`
	CreatedAt       *string `json:"created_at"`
	KpknlName       string  `json:"kpknl_name"`
	CategoryName    string  `json:"category_name"`
	SubCategoryName string  `json:"sub_category_name"`
	StatusName      string  `json:"status_name"`
	DocumentID      string  `json:"document_id"`

	// Addresses      address.Addresses              `json:"addresses"`
	// BuildingAssets building.BuildingAssets        `json:"building_assets"`
	// VehicleAssets  vehicle.VehicleAssets          `json:"vehicle_assets"`
	// Facilities     []facilities.FacilitiesRequest `json:"facilities"`
	// AccessPlaces   []access.AccessPlacesRequest   `json:"access_places"`
	// Contacts       contact.ContactsRequest        `json:"contacts"`
	// Images         []images.ImagesRequest         `json:"images"`
	// Approvals      approvals.Approvals            `json:"approvals"`
}

type AssetsResponseGetOne struct {
	ID              int64                           `json:"id"`
	FormType        string                          `json:"form_type"`
	Type            string                          `json:"type"`
	KpknlID         int64                           `json:"kpknl_id"`
	AuctionDate     string                          `json:"auction_date"`
	AuctionTime     string                          `json:"auction_time"`
	AuctionLink     string                          `json:"auction_link"`
	CategoryID      int64                           `json:"category_id"`
	SubCategoryID   int64                           `json:"sub_category_id"`
	Name            string                          `json:"name"`
	Price           int64                           `json:"price"`
	Description     string                          `json:"description"`
	Status          string                          `json:"status"`
	MakerID         string                          `json:"maker_id"`
	MakerDesc       string                          `json:"maker_desc"`
	MakerDate       *string                         `json:"maker_date"`
	LastMakerID     string                          `json:"last_maker_id"`
	LastMakerDesc   string                          `json:"last_maker_desc"`
	LastMakerDate   *string                         `json:"last_maker_date"`
	Published       bool                            `json:"published"`
	Deleted         bool                            `json:"deleted"`
	ExpiredDate     *string                         `json:"expired_date"`
	Action          string                          `json:"action"`
	KpknlName       string                          `json:"kpknl_name"`
	CategoryName    string                          `json:"category_name"`
	SubCategoryName string                          `json:"sub_category_name"`
	StatusName      string                          `json:"status_name"`
	Addresses       address.AddressesResponse       `json:"addresses"`
	BuildingAssets  building.BuildingAssetsResponse `json:"building_assets"`
	VehicleAssets   vehicle.VehicleAssetsResponse   `json:"vehicle_assets"`
	Facilities      []facilities.FacilitiesResponse `json:"facilities"`
	AccessPlaces    []access.AccessPlacesResponse   `json:"access_places"`
	Contacts        contact.ContactsResponse        `json:"contacts"`
	Images          []images.ImagesResponses        `json:"images"`
	DocumentID      string                          `json:"document_id"`
	Approvals       approvals.ApprovalsResponse     `json:"approvals"`
	UpdatedAt       *string                         `json:"updated_at"`
	CreatedAt       *string                         `json:"created_at"`
}

type AssetsRequestMaintain struct {
	Order     string `json:"order"`
	Sort      string `json:"sort"`
	Offset    int    `json:"offset"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	Type      string `json:"type"`
	Category  string `json:"category"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	PicName   string `json:"pic_name"`
	Status    string `json:"status"`
	CheckerID string `json:"checker_id"`
	SignerID  string `json:"signer_id"`
}

type AssetsResponseMaintain struct {
	ID          lib.NullInt64  `json:"id"`
	Type        lib.NullString `json:"type"`
	Category    lib.NullString `json:"category"`
	SubCategory lib.NullString `json:"sub_category"`
	Name        lib.NullString `json:"name"`
	Price       lib.NullInt64  `json:"price"`
	Status      lib.NullString `json:"status"`
	PicName     lib.NullString `json:"pic_name"`
	Published   lib.NullString `json:"published"`
	CheckerID   lib.NullString `json:"checker_id"`
	SignerID    lib.NullString `json:"signer_id"`
	UpdatedAt   lib.NullTime   `json:"updated_at"`
	CreatedAt   lib.NullTime   `json:"created_at"`
}

type AssetsResponses struct {
	ID          int64  `json:"id"`
	Type        string `json:"type"`
	Category    string `json:"category"`
	SubCategory string `json:"sub_category"`
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	PicName     string `json:"pic_name"`
	Status      string `json:"status"`
	Published   string `json:"published"`
	CheckerID   string `json:"checker_id"`
	SignerID    string `json:"signer_id"`
	UpdatedAt   string `json:"updated_at"`
	CreatedAt   string `json:"created_at"`
}

type AssetsRequestUpdate struct {
	ID            int64                      `json:"id"`
	Type          string                     `json:"type"`
	TypePublish   string                     `json:"type_publish"`
	LastMakerID   string                     `json:"last_maker_id"`
	LastMakerDesc string                     `json:"last_maker_desc"`
	LastMakerDate *string                    `json:"last_maker_date"`
	Approvals     approvals.ApprovalsRequest `json:"approvals"`
	DocumentID    string                     `json:"document_id"`
}

type AssetsRequestUpdateElastic struct {
	ID         int64  `json:"id"`
	DocumentID string `json:"document_id"`
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

func (a AssetsRequestMaintain) TableName() string {
	return "assets"
}

func (a AssetsRequestUpdateElastic) TableName() string {
	return "assets"
}
