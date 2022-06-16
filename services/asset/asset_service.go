package asset

import (
	"fmt"
	"infolelang/lib"
	"os"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"

	requestAddress "infolelang/models/addresses"
	requestApprovals "infolelang/models/approvals"
	requestBuilding "infolelang/models/building_assets"
	requestContact "infolelang/models/contacts"
	requestImage "infolelang/models/images"
	requestVehicle "infolelang/models/vehicle_assets"

	models "infolelang/models/assets"
	accessPlace "infolelang/repository/access_places"
	addressRepo "infolelang/repository/address"
	approvalRepo "infolelang/repository/approvals"
	assetRepo "infolelang/repository/asset"
	contactRepo "infolelang/repository/contacts"
	facilityRepo "infolelang/repository/facilities"
	imageRepo "infolelang/repository/images"

	"gitlab.com/golang-package-library/logger"
	minio "gitlab.com/golang-package-library/minio"
)

var (
	timeNow = lib.GetTimeNow("timestime")
	UUID    = uuid.NewString()
)

type AssetDefinition interface {
	WithTrx(trxHandle *gorm.DB) AssetService
	GetAll() (responses []models.AssetsResponse, err error)
	GetOne(id int64) (responses models.AssetsResponseGetOne, err error)
	Store(request *models.AssetsRequest) (err error)
	Update(request *models.AssetsRequest) (err error)
	Delete(id int64) (err error)
	GetApproval(request models.AssetsRequestMaintain) (responses []models.AssetsResponses, pagination lib.Pagination, err error)
	GetMaintain(request models.AssetsRequestMaintain) (responses []models.AssetsResponses, pagination lib.Pagination, err error)
}
type AssetService struct {
	minio            minio.Minio
	logger           logger.Logger
	assetRepo        assetRepo.AssetDefinition
	addressRepo      addressRepo.AddressDefinition
	buildingRepo     assetRepo.BuildingAssetDefinition
	vehicleRepo      assetRepo.VehicleAssetDefinition
	contactRepo      contactRepo.ContactDefinition
	approvalRepo     approvalRepo.ApprovalDefinition
	imagesRepo       imageRepo.ImageDefinition
	assetImage       assetRepo.AssetImageDefinition
	assetFacility    assetRepo.AssetFacilityDefinition
	facilityRepo     facilityRepo.FacilitiesDefinition
	assetAccessPlace assetRepo.AssetAccessPlaceDefinition
	accessPlace      accessPlace.AccessPlaceDefinition
}

func NewAssetService(
	minio minio.Minio,
	logger logger.Logger,
	assetRepo assetRepo.AssetDefinition,
	addressRepo addressRepo.AddressDefinition,
	buildingRepo assetRepo.BuildingAssetDefinition,
	vehicleRepo assetRepo.VehicleAssetDefinition,
	contactRepo contactRepo.ContactDefinition,
	approvalRepo approvalRepo.ApprovalDefinition,
	imagesRepo imageRepo.ImageDefinition,
	assetImage assetRepo.AssetImageDefinition,
	assetFacility assetRepo.AssetFacilityDefinition,
	facilityRepo facilityRepo.FacilitiesDefinition,
	assetAccessPlace assetRepo.AssetAccessPlaceDefinition,
	accessPlace accessPlace.AccessPlaceDefinition,
) AssetDefinition {
	return AssetService{
		minio:            minio,
		logger:           logger,
		assetRepo:        assetRepo,
		addressRepo:      addressRepo,
		buildingRepo:     buildingRepo,
		vehicleRepo:      vehicleRepo,
		contactRepo:      contactRepo,
		approvalRepo:     approvalRepo,
		imagesRepo:       imagesRepo,
		assetImage:       assetImage,
		assetFacility:    assetFacility,
		facilityRepo:     facilityRepo,
		assetAccessPlace: assetAccessPlace,
		accessPlace:      accessPlace,
	}
}

// WithTrx delegates transaction to repository database
func (asset AssetService) WithTrx(trxHandle *gorm.DB) AssetService {
	asset.assetRepo = asset.assetRepo.WithTrx(trxHandle)
	return asset
}

// GetAll implements AssetDefinition
func (asset AssetService) GetAll() (responses []models.AssetsResponse, err error) {
	return asset.assetRepo.GetAll()
}

// GetOne implements AssetDefinition
func (asset AssetService) GetOne(id int64) (responses models.AssetsResponseGetOne, err error) {
	assets, err := asset.assetRepo.GetOne(id)
	// get address
	address, err := asset.addressRepo.GetOneAsset(assets.ID)
	// get building
	building, err := asset.buildingRepo.GetOneAsset(assets.ID)
	// get vehicle
	vehicle, err := asset.vehicleRepo.GetOneAsset(assets.ID)

	// get asset facilities
	facilities, err := asset.assetFacility.GetOneAsset(assets.ID)
	// get access place
	accessPlace, err := asset.assetAccessPlace.GetOneAsset(assets.ID)
	// get contact
	contact, err := asset.contactRepo.GetOneAsset(assets.ID)

	// get images
	images, err := asset.imagesRepo.GetOneAsset(assets.ID)

	// get approval
	approval, err := asset.approvalRepo.GetOneAsset(assets.ID)

	responses = models.AssetsResponseGetOne{
		ID:             assets.ID,
		Type:           assets.Type,
		KpknlID:        assets.KpknlID,
		AuctionDate:    assets.AuctionDate,
		AuctionTime:    assets.AuctionTime,
		AuctionLink:    assets.AuctionLink,
		CategoryID:     assets.CategoryID,
		SubCategoryID:  assets.SubCategoryID,
		Name:           assets.Name,
		Price:          assets.Price,
		Description:    assets.Description,
		Status:         assets.Status,
		MakerID:        assets.MakerID,
		MakerDesc:      assets.MakerDesc,
		MakerComment:   assets.MakerComment,
		MakerDate:      assets.MakerDate,
		LastMakerID:    assets.LastMakerID,
		LastMakerDesc:  assets.LastMakerDesc,
		LastMakerDate:  assets.LastMakerDate,
		Published:      assets.Published,
		Deleted:        assets.Deleted,
		ExpiredDate:    assets.ExpiredDate,
		Action:         assets.Action,
		Addresses:      address,
		BuildingAssets: building,
		VehicleAssets:  vehicle,
		Facilities:     facilities,
		AccessPlaces:   accessPlace,
		Contacts:       contact,
		Images:         images,
		Approvals:      approval,
		UpdatedAt:      assets.UpdatedAt,
		CreatedAt:      assets.CreatedAt,
	}
	return responses, err
}

// Store implements AssetDefinition
func (asset AssetService) Store(request *models.AssetsRequest) (err error) {
	// create assets
	bucket := os.Getenv("BUCKET_NAME")

	dataAsset, err := asset.assetRepo.Store(&models.Assets{
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
		Status:        "01a", // pending checker
		CreatedAt:     &timeNow,
	})

	if err != nil {
		asset.logger.Zap.Error(err)
		return err
	}
	fmt.Println("dataAsset", dataAsset)

	// address

	address, err := asset.addressRepo.Store(
		&requestAddress.Addresses{
			AssetID:      dataAsset.ID,
			PostalcodeID: request.Addresses.PostalcodeID,
			Address:      request.Addresses.Address,
			Longitude:    request.Addresses.Longitude,
			Langitude:    request.Addresses.Langitude,
			CreatedAt:    &timeNow,
		})

	fmt.Println("this is address => ", address)

	if err != nil {
		asset.logger.Zap.Error(err)
		return err
	}

	// var building *requestBuilding.BuildingAssets
	// var vehicle *requestVehicle.VehicleAssets

	switch request.FormType {
	case "form-b1":
		// buildingasset
		building, err := asset.buildingRepo.Store(&requestBuilding.BuildingAssets{
			AssetID:           dataAsset.ID,
			CertificateType:   request.BuildingAssets.CertificateType,
			CertificateNumber: request.BuildingAssets.CertificateNumber,
			BuildYear:         request.BuildingAssets.BuildYear,
			SurfaceArea:       request.BuildingAssets.SurfaceArea,
			BuildingArea:      request.BuildingAssets.BuildingArea,
			Direction:         request.BuildingAssets.Direction,
			NumberOfFloors:    request.BuildingAssets.NumberOfFloors,
			NumberOfBedrooms:  request.BuildingAssets.NumberOfBedrooms,
			NumberOfBathrooms: request.BuildingAssets.NumberOfBathrooms,
			ElectricalPower:   request.BuildingAssets.ElectricalPower,
			Carport:           request.BuildingAssets.Carport,
			CreatedAt:         &timeNow,
		})
		if err != nil {
			asset.logger.Zap.Error(err)
			return err
		}
		fmt.Println("building=>", building)
	default:
		// vehicle asset
		vehicle, err := asset.vehicleRepo.Store(&requestVehicle.VehicleAssets{
			AssetID:           dataAsset.ID,
			VehicleType:       request.VehicleAssets.VehicleType,
			CertificateTypeID: request.VehicleAssets.CertificateTypeID,
			CertificateNumber: request.VehicleAssets.CertificateNumber,
			Series:            request.VehicleAssets.Series,
			BrandID:           request.VehicleAssets.BrandID,
			Type:              request.VehicleAssets.Type,
			ProductionYear:    request.VehicleAssets.ProductionYear,
			TransmissionID:    request.VehicleAssets.TransmissionID,
			MachineCapacityID: request.VehicleAssets.MachineCapacityID,
			ColorID:           request.VehicleAssets.ColorID,
			NumberOfSeat:      request.VehicleAssets.NumberOfSeat,
			NumberOfUsage:     request.VehicleAssets.NumberOfUsage,
			MachineNumber:     request.VehicleAssets.MachineNumber,
			BodyNumber:        request.VehicleAssets.BodyNumber,
			LicenceDate:       request.VehicleAssets.LicenceDate,
			CreatedAt:         &timeNow,
		})
		if err != nil {
			asset.logger.Zap.Error(err)
			return err
		}
		fmt.Println("vehicle=>", vehicle)

	}

	fmt.Println("facilities=>", request.Facilities)

	for _, value := range request.Facilities {
		_, err := asset.assetFacility.Store(
			&models.AssetFacilities{
				AssetID:    dataAsset.ID,
				FacilityID: value.ID,
				Status:     value.Status,
				CreatedAt:  &timeNow,
			})
		if err != nil {
			asset.logger.Zap.Error(err)
			return err
		}
	}

	// asset_access_places
	fmt.Println("Access Places=>", request.AccessPlaces)
	for _, value := range request.AccessPlaces {
		_, err = asset.assetAccessPlace.Store(
			&models.AssetAccessPlaces{
				AssetID:       dataAsset.ID,
				AccessPlaceID: value.ID,
				CreatedAt:     &timeNow,
			})

		if err != nil {
			asset.logger.Zap.Error(err)
			return err
		}
	}

	// contact
	contact, err := asset.contactRepo.Store(&requestContact.Contacts{
		AssetID:     dataAsset.ID,
		DebiturName: request.Contacts.DebiturName,
		PicName:     request.Contacts.PicName,
		PicPhone:    request.Contacts.PicPhone,
		PicEmail:    request.Contacts.PicEmail,
		Cif:         request.Contacts.Cif,
		CreatedAt:   &timeNow,
	})
	if err != nil {
		asset.logger.Zap.Error(err)
		return err
	}
	fmt.Println("contact=>", contact)

	var images []requestImage.ImagesRequest

	// images
	for _, value := range request.Images {

		var destinationPath string
		bucketExist := asset.minio.BucketExist(asset.minio.Client(), bucket)

		pathSplit := strings.Split(value.Path, "/")
		sourcePath := fmt.Sprint(value.Path)
		destinationPath = pathSplit[1] + "/" +
			dataAsset.Type + "/" +
			lib.GetTimeNow("year") + "/" +
			lib.GetTimeNow("month") + "/" +
			lib.GetTimeNow("day") + "/" +
			pathSplit[2] + "/" +
			value.Filename
		// assets/formb1/2022/June/01/uuid/gambar.jpg

		if bucketExist {
			fmt.Println("Exist")
			fmt.Println(bucket)
			fmt.Println(destinationPath)
			asset.minio.CopyObject(asset.minio.Client(), bucket, sourcePath, bucket, destinationPath)

		} else {
			fmt.Println("Not Exist")
			fmt.Println(bucket)
			fmt.Println(destinationPath)
			asset.minio.MakeBucket(asset.minio.Client(), bucket, "")
			asset.minio.CopyObject(asset.minio.Client(), bucket, sourcePath, bucket, destinationPath)
		}

		image, err := asset.imagesRepo.Store(&requestImage.Images{
			Filename:  value.Filename,
			Path:      destinationPath,
			Extension: value.Extension,
			Size:      value.Size,
			CreatedAt: &timeNow,
		})

		images = append(images, requestImage.ImagesRequest{
			ID:        image.ID,
			Filename:  value.Filename,
			Path:      destinationPath,
			Extension: value.Extension,
			Size:      value.Size,
			CreatedAt: &timeNow,
		})

		if err != nil {
			asset.logger.Zap.Error(err)
			return err
		}

		_, err = asset.assetImage.Store(&models.AssetImages{
			AssetID:   dataAsset.ID,
			ImageID:   image.ID,
			CreatedAt: &timeNow,
		})

		if err != nil {
			asset.logger.Zap.Error(err)
			return err
		}
	}
	fmt.Println("images=>", images)

	// approval
	approval, err := asset.approvalRepo.Store(
		&requestApprovals.Approvals{
			AssetID:     dataAsset.ID,
			CheckerID:   request.Approvals.CheckerID,
			CheckerDesc: request.Approvals.CheckerDesc,
			// CheckerComment: request.Approvals.CheckerComment,
			// CheckerDate:    request.Approvals.CheckerDate,
			SignerID:   request.Approvals.SignerID,
			SignerDesc: request.Approvals.SignerDesc,
			// SignerComment:  request.Approvals.SignerComment,
			// SignerDate:     request.Approvals.SignerDate,
			CreatedAt: &timeNow})
	if err != nil {
		asset.logger.Zap.Error(err)
		return err
	}
	fmt.Println("approval=>", approval)

	// create elastic
	// dataAssets := models.AssetsResponse{
	// 	ID:             dataAsset.ID,
	// 	Type:           dataAsset.Type,
	// 	KpknlID:        dataAsset.KpknlID,
	// 	AuctionDate:    dataAsset.AuctionDate,
	// 	AuctionTime:    dataAsset.AuctionTime,
	// 	AuctionLink:    dataAsset.AuctionLink,
	// 	CategoryID:     dataAsset.CategoryID,
	// 	SubCategoryID:  dataAsset.SubCategoryID,
	// 	Name:           dataAsset.Name,
	// 	Price:          dataAsset.Price,
	// 	Description:    dataAsset.Description,
	// 	Addresses:      *address,
	// 	BuildingAssets: *building,
	// 	VehicleAssets:  *vehicle,
	// 	Facilities:     request.Facilities,
	// 	AccessPlaces:   request.AccessPlaces,
	// 	Contacts:       request.Contacts,
	// 	Images:         images,
	// 	Approvals:      *approvals,
	// 	UpdatedAt:      dataAsset.UpdatedAt,
	// 	CreatedAt:      dataAsset.CreatedAt,
	// }

	// fmt.Println("TES")
	// fmt.Println(dataAssets)
	// _, err = asset.assetRepo.StoreElastic(dataAssets)
	// if err != nil {
	// 	asset.logger.Zap.Error(err)
	// 	return responses, err
	// }
	return err
}

// Update implements AssetDefinition
func (asset AssetService) Update(request *models.AssetsRequest) (err error) {
	_, err = asset.assetRepo.Update(request)
	return err
}

// Delete implements AssetDefinition
func (asset AssetService) Delete(id int64) (err error) {
	return asset.assetRepo.Delete(id)
}

func (asset AssetService) GetApproval(request models.AssetsRequestMaintain) (responses []models.AssetsResponses, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataAssets, totalRows, err := asset.assetRepo.GetApproval(request)
	if err != nil {
		asset.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataAssets {
		responses = append(responses, models.AssetsResponses{
			ID:          response.ID.Int64,
			Type:        response.Type.String,
			Category:    response.Category.String,
			SubCategory: response.SubCategory.String,
			Name:        response.Name.String,
			Price:       response.Price.Int64,
			Status:      response.Status.String,
			PicName:     response.PicName.String,
			Published:   response.Published.String,
			CheckerID:   response.CheckerID.String,
			SignerID:    response.SignerID.String,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows)
	return responses, pagination, err
}

func (asset AssetService) GetMaintain(request models.AssetsRequestMaintain) (responses []models.AssetsResponses, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	dataAssets, totalRows, err := asset.assetRepo.GetMaintain(request)
	if err != nil {
		asset.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataAssets {
		responses = append(responses, models.AssetsResponses{
			ID:          response.ID.Int64,
			Type:        response.Type.String,
			Category:    response.Category.String,
			SubCategory: response.SubCategory.String,
			Name:        response.Name.String,
			Price:       response.Price.Int64,
			Status:      response.Status.String,
			PicName:     response.PicName.String,
			Published:   response.Published.String,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows)
	return responses, pagination, err

}
