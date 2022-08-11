package asset

import (
	"fmt"
	"infolelang/lib"
	"os"
	"reflect"
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

	accessPlaceModel "infolelang/models/access_places"
	buildingModel "infolelang/models/building_assets"
	facilitiesModel "infolelang/models/facilities"
	VehicleModel "infolelang/models/vehicle_assets"

	"github.com/golang-module/carbon"
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
	GetAuctionSchedule(request models.AuctionSchedule) (responses []models.AuctionScheduleResponse, pagination lib.Pagination, err error)
	GetOne(id int64) (responses models.AssetsResponseGetOneString, status bool, err error)
	Store(request models.AssetsRequest) (status bool, err error)
	GetApproval(request models.AssetsRequestMaintain) (responses []models.AssetsResponses, pagination lib.Pagination, err error)
	GetMaintain(request models.AssetsRequestMaintain) (responses []models.AssetsResponses, pagination lib.Pagination, err error)
	UpdateApproval(request *models.AssetsRequestUpdate) (status bool, err error)
	UpdateMaintain(request models.AssetsResponseGetOne) (status bool, err error)
	Delete(request *models.AssetsRequestUpdate) (responses bool, err error)
	DeleteAssetImage(request *models.AssetImageRequest) (status bool, err error)
}
type AssetService struct {
	db               lib.Database
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
	db lib.Database,
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
		db:               db,
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
func (asset AssetService) GetAuctionSchedule(request models.AuctionSchedule) (responses []models.AuctionScheduleResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	responses, totalRows, totalData, err := asset.assetRepo.GetAuctionSchedule(request)
	if err != nil {
		asset.logger.Zap.Error(err)
		return responses, pagination, err
	}
	pagination = lib.SetPaginationResponse(page, limit, int(totalRows), int(totalData))
	return responses, pagination, err
}

// GetOne implements AssetDefinition
func (asset AssetService) GetOne(id int64) (responses models.AssetsResponseGetOneString, status bool, err error) {
	// join table
	assets, err := asset.assetRepo.GetOne(id)
	fmt.Println(assets)

	if assets.ID != 0 {
		fmt.Println("Bukan 0")
		// join table
		address, err := asset.addressRepo.GetOneAsset(assets.ID)
		building := buildingModel.BuildingAssetsResponse{}
		facilities := []facilitiesModel.FacilitiesResponse{}
		accessPlace := []accessPlaceModel.AccessPlacesResponse{}
		vehicle := VehicleModel.VehicleAssetsResponse{}
		fmt.Println("assets.FormType", assets.FormType)
		if assets.FormType == "form-b1" {
			fmt.Println("building")
			// join table
			building, err = asset.buildingRepo.GetOneAsset(assets.ID)
			fmt.Println("building.CertificateNumber", building.CertificateNumber)
			// join table
			facilities, err = asset.assetFacility.GetOneAsset(assets.ID)

			// join table
			accessPlace, err = asset.assetAccessPlace.GetOneAsset(assets.ID)
		} else {
			fmt.Println("vehicle")
			// join table
			vehicle, err = asset.vehicleRepo.GetOneAsset(assets.ID)
		}

		// join table
		contact, err := asset.contactRepo.GetOneAsset(assets.ID)

		// join table
		images, err := asset.imagesRepo.GetOneAsset(assets.ID)

		// join table
		approval, err := asset.approvalRepo.GetOneAsset(assets.ID)
		fmt.Println("assets.AuctionTime=>", assets.AuctionTime)
		responses = models.AssetsResponseGetOneString{
			ID:              assets.ID,
			FormType:        assets.FormType,
			Type:            assets.Type,
			KpknlID:         assets.KpknlID,
			AuctionDate:     carbon.Parse(fmt.Sprint(assets.AuctionDate)).ToDateString(),
			AuctionTime:     assets.AuctionTime,
			AuctionLink:     assets.AuctionLink,
			CategoryID:      assets.CategoryID,
			SubCategoryID:   assets.SubCategoryID,
			Name:            assets.Name,
			Price:           assets.Price,
			Description:     assets.Description,
			Status:          assets.Status,
			MakerID:         assets.MakerID,
			MakerDesc:       assets.MakerDesc,
			MakerDate:       carbon.Parse(fmt.Sprint(assets.MakerDate)).ToDateTimeString(),
			LastMakerID:     assets.LastMakerID,
			LastMakerDesc:   assets.LastMakerDesc,
			LastMakerDate:   carbon.Parse(fmt.Sprint(assets.LastMakerDate)).ToDateTimeString(),
			Published:       assets.Published,
			Deleted:         assets.Deleted,
			PublishDate:     carbon.Parse(fmt.Sprint(assets.PublishDate)).ToDateTimeString(),
			ExpiredDate:     carbon.Parse(fmt.Sprint(assets.ExpiredDate)).ToDateTimeString(),
			Action:          assets.Action,
			KpknlName:       assets.KpknlName,
			CategoryName:    assets.CategoryName,
			SubCategoryName: assets.SubCategoryName,
			StatusName:      assets.StatusName,
			Addresses:       address,
			BuildingAssets:  building,
			VehicleAssets:   vehicle,
			Facilities:      facilities,
			AccessPlaces:    accessPlace,
			Contacts:        contact,
			Images:          images,
			Approvals:       approval,
			DocumentID:      assets.DocumentID,
			UpdatedAt:       carbon.Parse(fmt.Sprint(assets.UpdatedAt)).ToDateTimeString(),
			CreatedAt:       carbon.Parse(fmt.Sprint(assets.CreatedAt)).ToDateTimeString(),
		}

		return responses, true, err
	}
	return responses, false, err
}

// Store implements AssetDefinition
func (asset AssetService) Store(request models.AssetsRequest) (status bool, err error) {
	tx := asset.db.DB.Begin()
	// , tx *gorm.DB
	// tx.Rollback()
	// tx.Commit()
	// create assets
	bucket := os.Getenv("BUCKET_NAME")
	assets := &models.Assets{}

	if request.Type == "Lelang" {
		assets = &models.Assets{
			FormType:      request.FormType,
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
			MakerID:       request.MakerID,
			MakerDesc:     request.MakerDesc,
			MakerDate:     request.MakerDate,
			LastMakerID:   request.LastMakerID,
			LastMakerDesc: request.LastMakerDesc,
			LastMakerDate: request.LastMakerDate,
			Action:        "Create",
			CreatedAt:     &timeNow,
		}
	} else {
		assets = &models.Assets{
			FormType:      request.FormType,
			Type:          request.Type,
			KpknlID:       request.KpknlID,
			AuctionDate:   nil,
			AuctionTime:   nil,
			AuctionLink:   request.AuctionLink,
			CategoryID:    request.CategoryID,
			SubCategoryID: request.SubCategoryID,
			Name:          request.Name,
			Price:         request.Price,
			Description:   request.Description,
			Status:        "01a", // pending checker
			MakerID:       request.MakerID,
			MakerDesc:     request.MakerDesc,
			MakerDate:     request.MakerDate,
			LastMakerID:   request.LastMakerID,
			LastMakerDesc: request.LastMakerDesc,
			LastMakerDate: request.LastMakerDate,
			Action:        "Create",
			CreatedAt:     &timeNow,
		}
	}

	dataAsset, err := asset.assetRepo.Store(assets, tx)

	if err != nil {
		tx.Rollback()
		asset.logger.Zap.Error(err)
		return false, err
	}
	fmt.Println("dataAsset", dataAsset)

	// address
	fmt.Println("request.Addresses.PostalcodeID", request.Addresses)
	address, err := asset.addressRepo.Store(
		&requestAddress.Addresses{
			AssetID:      dataAsset.ID,
			PostalcodeID: request.Addresses.PostalcodeID,
			Address:      request.Addresses.Address,
			Longitude:    request.Addresses.Longitude,
			Langitude:    request.Addresses.Langitude,
			CreatedAt:    &timeNow,
		}, tx)

	fmt.Println("this is address => ", address)

	if err != nil {
		tx.Rollback()
		asset.logger.Zap.Error(err)
		return false, err
	}

	switch request.FormType {
	case "form-b1":
		// buildingasset
		building, err := asset.buildingRepo.Store(&requestBuilding.BuildingAssets{
			AssetID:           dataAsset.ID,
			CertificateTypeID: request.BuildingAssets.CertificateTypeID,
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
		}, tx)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}
		fmt.Println("building=>", building)

		fmt.Println("facilities=>", request.Facilities)

		// check apabila array image error return false
		for _, value := range request.Facilities {
			_, err := asset.assetFacility.Store(
				&models.AssetFacilities{
					AssetID:    dataAsset.ID,
					FacilityID: value.ID,
					Status:     value.Status,
					CreatedAt:  &timeNow,
				}, tx)
			if err != nil {
				tx.Rollback()
				asset.logger.Zap.Error(err)
				return false, err
			}
		}

		// asset_access_places
		// check apabila array facilitie error return false
		fmt.Println("Access Places=>", request.AccessPlaces)
		for _, value := range request.AccessPlaces {
			_, err = asset.assetAccessPlace.Store(
				&models.AssetAccessPlaces{
					AssetID:       dataAsset.ID,
					AccessPlaceID: value.ID,
					Status:        value.Status,
					CreatedAt:     &timeNow,
				}, tx)

			if err != nil {
				tx.Rollback()
				asset.logger.Zap.Error(err)
				return false, err
			}
		}

	default:
		// vehicle asset
		vehicle, err := asset.vehicleRepo.Store(&requestVehicle.VehicleAssets{
			AssetID: dataAsset.ID,
			// VehicleTypeID:     request.VehicleAssets.VehicleTypeID,
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
		}, tx)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}
		fmt.Println("vehicle=>", vehicle)

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
	}, tx)
	fmt.Println("contact", contact)
	if err != nil {
		tx.Rollback()
		asset.logger.Zap.Error(err)
		return false, err
	}

	// images
	// check apabila array image error return false
	if len(request.Images) != 0 {
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
			fmt.Println("bucketExist=>", reflect.TypeOf(bucketExist))
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
			}, tx)

			if err != nil {
				tx.Rollback()
				asset.logger.Zap.Error(err)
				return false, err
			}

			_, err = asset.assetImage.Store(&models.AssetImages{
				AssetID:   dataAsset.ID,
				ImageID:   image.ID,
				CreatedAt: &timeNow,
			}, tx)

			if err != nil {
				tx.Rollback()
				asset.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		asset.logger.Zap.Error("Images Empty")
		return false, err
	}

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
			CreatedAt: &timeNow}, tx)
	if err != nil {
		tx.Rollback()
		asset.logger.Zap.Error(err)
		return false, err
	}
	fmt.Println("approval=>", approval)

	tx.Commit()
	return true, err
}

// UpdateApproval implements AssetDefinition
// get document id and update to table asset not handled
func (asset AssetService) UpdateApproval(request *models.AssetsRequestUpdate) (status bool, err error) {
	tx := asset.db.DB.Begin()
	switch request.Type {
	//===================== Approve Checker =====================
	case "approve checker":
		_, err = asset.assetRepo.UpdateApproval(&models.AssetsUpdateApproval{
			ID:        request.ID,
			Status:    "01c", // pending signer
			Action:    "UpdateApproval",
			UpdatedAt: &timeNow,
		},
			[]string{"status", "action", "updated_at"}, // define field to update
			tx)

		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}

		// approval
		err = asset.approvalRepo.DeleteApprovals(request.ID, tx)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error()
			return false, err
		}

		_, err = asset.approvalRepo.Store(
			&requestApprovals.Approvals{
				AssetID:        request.ID,
				CheckerID:      request.Approvals.CheckerID,
				CheckerDesc:    request.Approvals.CheckerDesc,
				CheckerComment: request.Approvals.CheckerComment,
				CheckerDate:    request.Approvals.CheckerDate,
				SignerID:       request.Approvals.SignerID,
				SignerDesc:     request.Approvals.SignerDesc,
				// SignerComment:  request.Approvals.SignerComment,
				// SignerDate:     request.Approvals.SignerDate,
				UpdatedAt: &timeNow}, tx)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}
		tx.Commit()
		return true, err
		//===================== Approve Checker End =====================

	//===================== Approve Signer =====================
	case "approve signer":
		tx = asset.db.DB.Begin()
		_, err = asset.assetRepo.UpdateApproval(&models.AssetsUpdateApproval{
			ID:          request.ID,
			Published:   true,
			PublishDate: &timeNow,
			ExpiredDate: lib.AddTime(0, 6, 0),
			Status:      "01e", // published
			Action:      "UpdateApproval",
			UpdatedAt:   &timeNow,
		},
			[]string{"published", "publish_date", "expired_date", "status", "action", "updated_at"}, // define field to update
			tx)

		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}

		// approval
		err = asset.approvalRepo.DeleteApprovals(request.ID, tx)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error()
			return false, err
		}

		_, err := asset.approvalRepo.Store(
			&requestApprovals.Approvals{
				AssetID:        request.ID,
				CheckerID:      request.Approvals.CheckerID,
				CheckerDesc:    request.Approvals.CheckerDesc,
				CheckerComment: request.Approvals.CheckerComment,
				CheckerDate:    request.Approvals.CheckerDate,
				SignerID:       request.Approvals.SignerID,
				SignerDesc:     request.Approvals.SignerDesc,
				SignerComment:  request.Approvals.SignerComment,
				SignerDate:     request.Approvals.SignerDate,
				UpdatedAt:      &timeNow}, tx)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}

		documentID := lib.UUID(false)
		update, err := asset.assetRepo.UpdateDocumentID(&models.AssetsRequestUpdateElastic{
			ID:         request.ID,
			DocumentID: documentID,
		},
			[]string{
				"document_id",
			}, tx)
		if !update || err != nil {
			tx.Rollback()
			return false, err
		}

		// create elastic
		_, _, err = asset.GetOne(request.ID)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error()
			return false, err
		}
		tx.Rollback()
		return false, err
		//===================== Approve Signer End =====================

	//===================== Tolak Checker =====================
	case "tolak checker":
		tx := asset.db.DB.Begin()
		// update checker
		_, err = asset.assetRepo.UpdateApproval(&models.AssetsUpdateApproval{
			ID: request.ID,
			// Published: ,
			// PublishDate: ,
			Status:    "01b", // ditolak checker
			Action:    "UpdateApproval",
			UpdatedAt: &timeNow,
		},
			[]string{"status", "action", "updated_at"}, // define field to update
			tx)

		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}

		// approval
		err = asset.approvalRepo.DeleteApprovals(request.ID, tx)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error()
			return false, err
		}

		_, err := asset.approvalRepo.Store(
			&requestApprovals.Approvals{
				AssetID:        request.ID,
				CheckerID:      request.Approvals.CheckerID,
				CheckerDesc:    request.Approvals.CheckerDesc,
				CheckerComment: request.Approvals.CheckerComment,
				CheckerDate:    request.Approvals.CheckerDate,
				SignerID:       request.Approvals.SignerID,
				SignerDesc:     request.Approvals.SignerDesc,
				// SignerComment:  request.Approvals.SignerComment,
				// SignerDate:     request.Approvals.SignerDate,
				UpdatedAt: &timeNow}, tx)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}
		tx.Commit()
		return true, err
	//===================== Tolak Checker End =====================

	//===================== Tolak Signer      =====================
	case "tolak signer":
		tx := asset.db.DB.Begin()
		_, err = asset.assetRepo.UpdateApproval(&models.AssetsUpdateApproval{
			ID: request.ID,
			// Published:   true,
			// PublishDate: &timeNow,
			// ExpiredDate: lib.AddTime(0, 6, 0),
			Status:    "01d", // ditolak signer
			Action:    "UpdateApproval",
			UpdatedAt: &timeNow,
		},
			[]string{
				"status",
				"action",
				"updated_at"}, // define field to update
			tx)

		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}

		// approval
		err = asset.approvalRepo.DeleteApprovals(request.ID, tx)
		_, err = asset.approvalRepo.Store(
			&requestApprovals.Approvals{
				AssetID:        request.ID,
				CheckerID:      request.Approvals.CheckerID,
				CheckerDesc:    request.Approvals.CheckerDesc,
				CheckerComment: request.Approvals.CheckerComment,
				CheckerDate:    request.Approvals.CheckerDate,
				SignerID:       request.Approvals.SignerID,
				SignerDesc:     request.Approvals.SignerDesc,
				SignerComment:  request.Approvals.SignerComment,
				SignerDate:     request.Approvals.SignerDate,
				UpdatedAt:      &timeNow}, tx)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}
		tx.Commit()
		return true, err
		//===================== Tolak Signer End =====================
	default:
		tx.Rollback()
		return false, err
	}

}

// Update implements AssetDefinition
func (asset AssetService) UpdateMaintain(request models.AssetsResponseGetOne) (status bool, err error) {
	tx := asset.db.DB.Begin()
	// update here
	// create assets
	bucket := os.Getenv("BUCKET_NAME")
	assets := &models.AssetsRequestUpdateMaintain{}
	include := []string{
		"type",
		"kpknl_id",
		"auction_date",
		"auction_time",
		"auction_link",
		"category_id",
		"sub_category_id",
		"name",
		"price",
		"description",
		"last_maker_id",
		"last_maker_desc",
		"last_maker_date",
		"action",
		"updated_at",
	}
	if request.Type == "Lelang" {

		assets = &models.AssetsRequestUpdateMaintain{
			ID:   request.ID,
			Type: request.Type,
			// FormType:      request.FormType,
			KpknlID:       request.KpknlID,
			AuctionDate:   request.AuctionDate,
			AuctionTime:   request.AuctionTime,
			AuctionLink:   request.AuctionLink,
			CategoryID:    request.CategoryID,
			SubCategoryID: request.SubCategoryID,
			// Status:        request.Status,
			Name:          request.Name,
			Price:         request.Price,
			Description:   request.Description,
			LastMakerID:   request.LastMakerID,
			LastMakerDesc: request.LastMakerDesc,
			LastMakerDate: request.LastMakerDate,
			DocumentID:    request.DocumentID,
			Action:        "UpdateMaintain",
			UpdatedAt:     &timeNow,
		}
		include = []string{
			"type",
			"kpknl_id",
			"auction_date",
			"auction_time",
			"auction_link",
			"category_id",
			"sub_category_id",
			"name",
			"price",
			"description",
			"last_maker_id",
			"last_maker_desc",
			"last_maker_date",
			"action",
			"updated_at",
		}
	} else {
		assets = &models.AssetsRequestUpdateMaintain{
			ID:      request.ID,
			Type:    request.Type,
			KpknlID: request.KpknlID,
			// AuctionDate:   "",
			// AuctionTime:   "",
			// AuctionLink:   "",
			CategoryID:    request.CategoryID,
			SubCategoryID: request.SubCategoryID,
			// Status:        request.Status,
			Name:          request.Name,
			Price:         request.Price,
			Description:   request.Description,
			LastMakerID:   request.LastMakerID,
			LastMakerDesc: request.LastMakerDesc,
			LastMakerDate: request.LastMakerDate,
			DocumentID:    request.DocumentID,
			Action:        "UpdateMaintain",
			UpdatedAt:     &timeNow,
		}
		include = []string{
			"type",
			"kpknl_id",
			// "auction_date",
			// "auction_time",
			// "auction_link",
			"category_id",
			"sub_category_id",
			"name",
			"price",
			"description",
			"last_maker_id",
			"last_maker_desc",
			"last_maker_date",
			"action",
			"updated_at",
		}
	}

	dataAsset, err := asset.assetRepo.UpdateMaintain(
		assets,
		include, tx)
	if err != nil {
		tx.Rollback()
		asset.logger.Zap.Error(err)
		return false, err
	}

	fmt.Println("dataAsset", dataAsset)

	// address

	address, err := asset.addressRepo.Store(
		&requestAddress.Addresses{
			ID:           request.Addresses.ID,
			AssetID:      dataAsset.ID,
			PostalcodeID: request.Addresses.PostalcodeID,
			Address:      request.Addresses.Address,
			Longitude:    request.Addresses.Longitude,
			Langitude:    request.Addresses.Langitude,
			UpdatedAt:    &timeNow,
		}, tx)

	fmt.Println("this is address => ", address)

	if err != nil {
		tx.Rollback()
		asset.logger.Zap.Error(err)
		return false, err
	}

	switch request.FormType {
	case "form-b1":
		// buildingasset
		building, err := asset.buildingRepo.Store(&requestBuilding.BuildingAssets{
			ID:                request.BuildingAssets.ID,
			AssetID:           dataAsset.ID,
			CertificateTypeID: request.BuildingAssets.CertificateTypeID,
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
			UpdatedAt:         &timeNow,
		}, tx)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}
		fmt.Println("building=>", building)

		fmt.Println("facilities=>", request.Facilities)

		// check apabila array image error return false
		for _, value := range request.Facilities {
			_, err := asset.assetFacility.Store(
				&models.AssetFacilities{
					ID:         value.ID,
					AssetID:    dataAsset.ID,
					FacilityID: value.ID,
					Status:     value.Status,
					UpdatedAt:  &timeNow,
				}, tx)
			if err != nil {
				tx.Rollback()
				asset.logger.Zap.Error(err)
				return false, err
			}
		}

		// asset_access_places
		// check apabila array facilitie error return false
		fmt.Println("Access Places=>", request.AccessPlaces)
		for _, value := range request.AccessPlaces {
			_, err = asset.assetAccessPlace.Store(
				&models.AssetAccessPlaces{
					ID:            value.ID,
					AssetID:       dataAsset.ID,
					AccessPlaceID: value.ID,
					Status:        value.Status,
					UpdatedAt:     &timeNow,
				}, tx)

			if err != nil {
				tx.Rollback()
				asset.logger.Zap.Error(err)
				return false, err
			}
		}
	default:
		// vehicle asset
		vehicle, err := asset.vehicleRepo.Store(&requestVehicle.VehicleAssets{
			ID:      request.BuildingAssets.ID,
			AssetID: dataAsset.ID,
			// VehicleTypeID:     request.VehicleAssets.VehicleTypeID,
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
			UpdatedAt:         &timeNow,
		}, tx)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}
		fmt.Println("vehicle=>", vehicle)

	}

	// contact
	contact, err := asset.contactRepo.Store(&requestContact.Contacts{
		ID:          request.Contacts.ID,
		AssetID:     dataAsset.ID,
		DebiturName: request.Contacts.DebiturName,
		PicName:     request.Contacts.PicName,
		PicPhone:    request.Contacts.PicPhone,
		PicEmail:    request.Contacts.PicEmail,
		Cif:         request.Contacts.Cif,
		UpdatedAt:   &timeNow,
	}, tx)
	if err != nil {
		tx.Rollback()
		asset.logger.Zap.Error(err)
		return false, err
	}
	fmt.Println("contact=>", contact)

	// images
	// check apabila array image error return false
	if len(request.Images) != 0 {
		// Delete images where asset_id
		err := asset.assetImage.DeleteAssetID(assets.ID, tx)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}

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

			if pathSplit[0] == "tmp" {
				asset.logger.Zap.Info("============> new images")
				// copy to origin directory and create image to db
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
					ID:        value.ID,
					Filename:  value.Filename,
					Path:      destinationPath,
					Extension: value.Extension,
					Size:      value.Size,
					UpdatedAt: &timeNow,
				}, tx)

				if err != nil {
					tx.Rollback()
					asset.logger.Zap.Error(err)
					return false, err
				}

				_, err = asset.assetImage.Store(&models.AssetImages{
					ID:        image.ID,
					AssetID:   dataAsset.ID,
					ImageID:   image.ID,
					UpdatedAt: &timeNow,
				}, tx)

				if err != nil {
					tx.Rollback()
					asset.logger.Zap.Error(err)
					return false, err
				}
			} else {
				// else update path to db with relation asset_image
				asset.logger.Zap.Info("============> old images")
				image, err := asset.imagesRepo.Store(&requestImage.Images{
					ID:        value.ID,
					Filename:  value.Filename,
					Path:      value.Path,
					Extension: value.Extension,
					Size:      value.Size,
					UpdatedAt: &timeNow,
				}, tx)

				if err != nil {
					tx.Rollback()
					asset.logger.Zap.Error(err)
					return false, err
				}

				_, err = asset.assetImage.Store(&models.AssetImages{
					ID:        image.ID,
					AssetID:   dataAsset.ID,
					ImageID:   image.ID,
					UpdatedAt: &timeNow,
				}, tx)

				if err != nil {
					tx.Rollback()
					asset.logger.Zap.Error(err)
					return false, err
				}
			}

		}
	} else {
		tx.Rollback()
		// asset.logger.Zap.Error(err)
		return false, err
	}

	// update approval
	_, err = asset.assetRepo.UpdateApproval(&models.AssetsUpdateApproval{
		ID:        request.ID,
		Status:    "01c", // pending signer
		UpdatedAt: &timeNow,
	},
		[]string{"status", "updated_at"}, // define field to update
		tx)

	if err != nil {
		tx.Rollback()
		asset.logger.Zap.Error(err)
		return false, err
	}

	// approval
	err = asset.approvalRepo.DeleteApprovals(request.ID, tx)
	if err != nil {
		tx.Rollback()
		asset.logger.Zap.Error()
		return false, err
	}

	_, err = asset.approvalRepo.Store(
		&requestApprovals.Approvals{
			AssetID:        request.ID,
			CheckerID:      request.Approvals.CheckerID,
			CheckerDesc:    request.Approvals.CheckerDesc,
			CheckerComment: request.Approvals.CheckerComment,
			CheckerDate:    request.Approvals.CheckerDate,
			SignerID:       request.Approvals.SignerID,
			SignerDesc:     request.Approvals.SignerDesc,
			// SignerComment:  request.Approvals.SignerComment,
			// SignerDate:     request.Approvals.SignerDate,
			UpdatedAt: &timeNow}, tx)
	if err != nil {
		tx.Rollback()
		asset.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// GetApproval implements AssetDefinition
func (asset AssetService) GetApproval(request models.AssetsRequestMaintain) (responses []models.AssetsResponses, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataAssets, totalRows, totalData, err := asset.assetRepo.GetApproval(request)
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
			Price:       float32(response.Price.Float64),
			Status:      response.Status.String,
			PicName:     response.PicName.String,
			Published:   response.Published.String,
			CheckerID:   response.CheckerID.String,
			SignerID:    response.SignerID.String,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// GetMaintain implements AssetDefinition
func (asset AssetService) GetMaintain(request models.AssetsRequestMaintain) (responses []models.AssetsResponses, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	dataAssets, totalRows, totalData, err := asset.assetRepo.GetMaintain(request)
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
			Price:       float32(response.Price.Float64),
			Status:      response.Status.String,
			PicName:     response.PicName.String,
			Published:   response.Published.String,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// Delete implements AssetDefinition
func (asset AssetService) Delete(request *models.AssetsRequestUpdate) (status bool, err error) {
	switch request.Type {
	//===================== Approve Checker =====================
	case "approve checker":
		tx := asset.db.DB.Begin()
		_, err = asset.assetRepo.UpdateApproval(&models.AssetsUpdateApproval{
			ID:            request.ID,
			LastMakerID:   request.LastMakerID,
			LastMakerDesc: request.LastMakerDesc,
			LastMakerDate: request.LastMakerDate,
			Status:        "01c", // pending signer
			Action:        "updateDelete",
			UpdatedAt:     &timeNow,
		},
			[]string{
				"last_maker_id",
				"last_maker_desc",
				"last_maker_date",
				"status",
				"action",
				"updated_at"}, // define field to update
			tx)

		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}

		err = asset.approvalRepo.DeleteApprovals(request.ID, tx)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error()
			return false, err
		}
		_, err := asset.approvalRepo.Store(
			&requestApprovals.Approvals{
				AssetID:        request.ID,
				CheckerID:      request.Approvals.CheckerID,
				CheckerDesc:    request.Approvals.CheckerDesc,
				CheckerComment: request.Approvals.CheckerComment,
				CheckerDate:    request.Approvals.CheckerDate,
				SignerID:       request.Approvals.SignerID,
				SignerDesc:     request.Approvals.SignerDesc,
				// SignerComment:  request.Approvals.SignerComment,
				// SignerDate:     request.Approvals.SignerDate,
				UpdatedAt: &timeNow}, tx)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}
		tx.Commit()
		return true, err
		//===================== Approve Checker End =====================

	//===================== Approve Signer =====================
	case "approve signer":
		tx := asset.db.DB.Begin()
		getOneAsset, exist, err := asset.GetOne(request.ID)
		if err != nil {
			asset.logger.Zap.Error(err)
			tx.Rollback()
			return false, err
		}
		fmt.Println("service", getOneAsset)
		fmt.Println("exist", exist)
		fmt.Println("if published false", getOneAsset.Published)

		updateData := &models.AssetsUpdateDelete{}

		if getOneAsset.DocumentID == "" {
			updateData = &models.AssetsUpdateDelete{
				ID:            request.ID,
				LastMakerID:   request.LastMakerID,
				LastMakerDesc: request.LastMakerDesc,
				LastMakerDate: request.LastMakerDate,
				Deleted:       true,
				Action:        "updateDelete",
				Status:        "02b", // selesai
				Published:     false,
				PublishDate:   nil,
				ExpiredDate:   nil,
				UpdatedAt:     &timeNow,
			}
		} else {
			updateData = &models.AssetsUpdateDelete{
				ID:            request.ID,
				LastMakerID:   request.LastMakerID,
				LastMakerDesc: request.LastMakerDesc,
				LastMakerDate: request.LastMakerDate,
				Deleted:       true,
				Action:        "updateDelete",
				Status:        "02b", // selesai
				Published:     false,
				PublishDate:   nil,
				ExpiredDate:   nil,
				UpdatedAt:     &timeNow,
			}
		}

		_, err = asset.assetRepo.Delete(updateData,
			[]string{
				"last_maker_id",
				"last_maker_desc",
				"last_aker_date",
				"deleted",
				"action",
				"status",
				"published",
				"publish_date",
				"expired_date",
				"updated_at",
			}, tx)

		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}

		err = asset.approvalRepo.DeleteApprovals(request.ID, tx)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error()
			return false, err
		}
		_, err = asset.approvalRepo.Store(
			&requestApprovals.Approvals{
				AssetID:        request.ID,
				CheckerID:      request.Approvals.CheckerID,
				CheckerDesc:    request.Approvals.CheckerDesc,
				CheckerComment: request.Approvals.CheckerComment,
				CheckerDate:    request.Approvals.CheckerDate,
				SignerID:       request.Approvals.SignerID,
				SignerDesc:     request.Approvals.SignerDesc,
				SignerComment:  request.Approvals.SignerComment,
				SignerDate:     request.Approvals.SignerDate,
				UpdatedAt:      &timeNow}, tx)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}

		if exist {
			fmt.Println("getOneAsset", getOneAsset)

			tx.Commit()
			return true, err
		}

		return false, err
		//===================== Approve Signer End =====================

	//===================== Tolak Checker =====================
	case "tolak checker":
		tx := asset.db.DB.Begin()
		_, err = asset.assetRepo.UpdateApproval(&models.AssetsUpdateApproval{
			ID:            request.ID,
			LastMakerID:   request.LastMakerID,
			LastMakerDesc: request.LastMakerDesc,
			LastMakerDate: request.LastMakerDate,
			Status:        "01b", // tolak checker
			Action:        "updateDelete",
			UpdatedAt:     &timeNow,
		},
			[]string{
				"last_maker_id",
				"last_maker_desc",
				"last_maker_date",
				"status",
				"action",
				"updated_at"}, // define field to update
			tx)

		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}

		err = asset.approvalRepo.DeleteApprovals(request.ID, tx)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error()
			return false, err
		}
		_, err := asset.approvalRepo.Store(
			&requestApprovals.Approvals{
				AssetID:        request.ID,
				CheckerID:      request.Approvals.CheckerID,
				CheckerDesc:    request.Approvals.CheckerDesc,
				CheckerComment: request.Approvals.CheckerComment,
				CheckerDate:    request.Approvals.CheckerDate,
				SignerID:       request.Approvals.SignerID,
				SignerDesc:     request.Approvals.SignerDesc,
				// SignerComment:  request.Approvals.SignerComment,
				// SignerDate:     request.Approvals.SignerDate,
				UpdatedAt: &timeNow}, tx)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}
		tx.Commit()
		return true, err
		//===================== Tolak Checker End =====================

	//===================== Tolak Signer =====================
	case "tolak signer":
		tx := asset.db.DB.Begin()
		_, err = asset.assetRepo.UpdateApproval(&models.AssetsUpdateApproval{
			ID:            request.ID,
			LastMakerID:   request.LastMakerID,
			LastMakerDesc: request.LastMakerDesc,
			LastMakerDate: request.LastMakerDate,
			Status:        "01d", // tolak signer
			Action:        "updateDelete",
			UpdatedAt:     &timeNow,
		},
			[]string{
				"last_maker_id",
				"last_maker_desc",
				"last_maker_date",
				"status",
				"action",
				"updated_at"}, // define field to update
			tx)

		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}

		err = asset.approvalRepo.DeleteApprovals(request.ID, tx)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}

		_, err := asset.approvalRepo.Store(
			&requestApprovals.Approvals{
				AssetID:        request.ID,
				CheckerID:      request.Approvals.CheckerID,
				CheckerDesc:    request.Approvals.CheckerDesc,
				CheckerComment: request.Approvals.CheckerComment,
				CheckerDate:    request.Approvals.CheckerDate,
				SignerID:       request.Approvals.SignerID,
				SignerDesc:     request.Approvals.SignerDesc,
				SignerComment:  request.Approvals.SignerComment,
				SignerDate:     request.Approvals.SignerDate,
				UpdatedAt:      &timeNow}, tx)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}
		tx.Commit()
		return true, err
		//===================== Tolak Signer End =====================
	default:
		return false, err
	}
}

// DeleteAssetImage implements AssetDefinition
func (asset AssetService) DeleteAssetImage(request *models.AssetImageRequest) (status bool, err error) {
	bucket := os.Getenv("BUCKET_NAME")
	ok := asset.minio.RemoveObject(asset.minio.Client(), bucket, request.Path)
	if !ok {
		return false, err
	} else {
		tx := asset.db.DB.Begin()
		err = asset.imagesRepo.Delete(request.ImageID, tx)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}

		err = asset.assetRepo.DeleteAssetImage(request.AssetImageID, tx)
		if err != nil {
			tx.Rollback()
			asset.logger.Zap.Error(err)
			return false, err
		}
		tx.Commit()
	}

	return true, err
}
