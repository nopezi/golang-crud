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
	Store(request *models.AssetsRequest) (status bool, err error)
	GetApproval(request models.AssetsRequestMaintain) (responses []models.AssetsResponses, pagination lib.Pagination, err error)
	GetMaintain(request models.AssetsRequestMaintain) (responses []models.AssetsResponses, pagination lib.Pagination, err error)
	UpdateApproval(request *models.AssetsRequestUpdate) (status bool, err error)
	UpdatePublish(request *models.AssetsRequestUpdate) (status bool, err error)
	UpdateMaintain(request *models.AssetsRequest) (status bool, err error)
	Delete(request *models.AssetsRequestUpdate) (responses bool, err error)
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
	// join table
	assets, err := asset.assetRepo.GetOne(id)

	// join table
	address, err := asset.addressRepo.GetOneAsset(assets.ID)

	// join table
	building, err := asset.buildingRepo.GetOneAsset(assets.ID)

	// join table
	vehicle, err := asset.vehicleRepo.GetOneAsset(assets.ID)

	// join table
	facilities, err := asset.assetFacility.GetOneAsset(assets.ID)

	// join table
	accessPlace, err := asset.assetAccessPlace.GetOneAsset(assets.ID)

	// join table
	contact, err := asset.contactRepo.GetOneAsset(assets.ID)

	// join table
	images, err := asset.imagesRepo.GetOneAsset(assets.ID)

	// join table
	approval, err := asset.approvalRepo.GetOneAsset(assets.ID)

	responses = models.AssetsResponseGetOne{
		ID:              assets.ID,
		Type:            assets.Type,
		KpknlID:         assets.KpknlID,
		AuctionDate:     assets.AuctionDate,
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
		MakerDate:       assets.MakerDate,
		LastMakerID:     assets.LastMakerID,
		LastMakerDesc:   assets.LastMakerDesc,
		LastMakerDate:   assets.LastMakerDate,
		Published:       assets.Published,
		Deleted:         assets.Deleted,
		ExpiredDate:     assets.ExpiredDate,
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
		UpdatedAt:       assets.UpdatedAt,
		CreatedAt:       assets.CreatedAt,
	}
	return responses, err
}

// Store implements AssetDefinition
func (asset AssetService) Store(request *models.AssetsRequest) (status bool, err error) {
	// create assets
	bucket := os.Getenv("BUCKET_NAME")
	// -- / asset
	// --  add maker_id. maker_desc, maker_comment, maker_date, last_maker_id. last_maker_desc, last_maker_date, -- jadi pas create maker = last_maker, update dll input ke last_maker
	// -- done add action varchar(100) misal approve hapus, approve create. approve, dia update terus
	// -- done add field deleted tinyint(1) -- ini buat soft delete
	// -- done add expired_date datetime  = published_date datetime= now() pas approve + 6 bulan ketika update approve
	// -- apabila lebih dari 6 bulan dan status publish maka scheduler jalan untuk me unpublis status
	// -- / asset

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
		MakerID:       request.MakerID,
		MakerDesc:     request.MakerDesc,
		MakerDate:     request.MakerDate,
		LastMakerID:   request.LastMakerID,
		LastMakerDesc: request.LastMakerDesc,
		LastMakerDate: request.LastMakerDate,
		// Published:     request.Published,
		// Deleted:       request.Deleted,
		// ExpiredDate:   request.ExpiredDate,
		Action:    "Create",
		CreatedAt: &timeNow,
	})

	if err != nil {
		asset.logger.Zap.Error(err)
		return false, err
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
		return false, err
	}

	// var building *requestBuilding.BuildingAssets
	// var vehicle *requestVehicle.VehicleAssets

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
		})
		if err != nil {
			asset.logger.Zap.Error(err)
			return false, err
		}
		fmt.Println("building=>", building)
	default:
		// vehicle asset
		vehicle, err := asset.vehicleRepo.Store(&requestVehicle.VehicleAssets{
			AssetID:           dataAsset.ID,
			VehicleTypeID:     request.VehicleAssets.VehicleTypeID,
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
			return false, err
		}
		fmt.Println("vehicle=>", vehicle)

	}

	fmt.Println("facilities=>", request.Facilities)

	// check apabila array image error return false
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
				CreatedAt:     &timeNow,
			})

		if err != nil {
			asset.logger.Zap.Error(err)
			return false, err
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
		return false, err
	}
	fmt.Println("contact=>", contact)

	var images []requestImage.ImagesRequest

	// images
	// check apabila array image error return false
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
			return false, err
		}

		_, err = asset.assetImage.Store(&models.AssetImages{
			AssetID:   dataAsset.ID,
			ImageID:   image.ID,
			CreatedAt: &timeNow,
		})

		if err != nil {
			asset.logger.Zap.Error(err)
			return false, err
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
		return false, err
	}
	fmt.Println("approval=>", approval)

	return true, err
}

// UpdateApproval implements AssetDefinition
func (asset AssetService) UpdateApproval(request *models.AssetsRequestUpdate) (status bool, err error) {

	switch request.Type {
	case "approve checker":
		// update checker
		_, err = asset.assetRepo.UpdateApproval(&models.AssetsUpdateApproval{
			ID: request.ID,
			// Published: ,
			// PublishDate: ,
			Status:    "01c", // approve checker
			Action:    "UpdateApproval",
			UpdatedAt: &timeNow,
		})

		if err != nil {
			asset.logger.Zap.Error(err)
			return false, err
		}

		// approval
		asset.approvalRepo.DeleteApprovals(request.ID)
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
				UpdatedAt: &timeNow})
		if err != nil {
			asset.logger.Zap.Error(err)
			return false, err
		}
		return true, err
	case "approve signer":
		_, err = asset.assetRepo.UpdateApproval(&models.AssetsUpdateApproval{
			ID:          request.ID,
			Published:   true,
			PublishDate: &timeNow,
			ExpiredDate: lib.AddTime(0, 6, 0),
			Status:      "01e", // approve signer
			Action:      "UpdateApproval",
			UpdatedAt:   &timeNow,
		})

		if err != nil {
			asset.logger.Zap.Error(err)
			return false, err
		}

		// approval
		asset.approvalRepo.DeleteApprovals(request.ID)
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
				UpdatedAt:      &timeNow})
		if err != nil {
			asset.logger.Zap.Error(err)
			return false, err
		}

		// create elastic
		//  update document id to asset, untuk menghapus asset ketika di unpublish
		getOneAsset, err := asset.GetOne(request.ID)

		fmt.Println("getOneAsset", getOneAsset)
		_, err = asset.assetRepo.StoreElastic(getOneAsset)

		if err != nil {
			asset.logger.Zap.Error(err)
			return false, err
		}

		return true, err
	case "tolak checker":
		// update checker
		_, err = asset.assetRepo.UpdateApproval(&models.AssetsUpdateApproval{
			ID: request.ID,
			// Published: ,
			// PublishDate: ,
			Status:    "01b", // tolak checker
			Action:    "UpdateApproval",
			UpdatedAt: &timeNow,
		})

		if err != nil {
			asset.logger.Zap.Error(err)
			return false, err
		}

		// approval
		asset.approvalRepo.DeleteApprovals(request.ID)
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
				UpdatedAt: &timeNow})
		if err != nil {
			asset.logger.Zap.Error(err)
			return false, err
		}
		return true, err
		// update tolak checker
	case "tolak signer":
		_, err = asset.assetRepo.UpdateApproval(&models.AssetsUpdateApproval{
			ID: request.ID,
			// Published:   true,
			// PublishDate: &timeNow,
			// ExpiredDate: lib.AddTime(0, 6, 0),
			Status:    "01d", // tolak signer
			Action:    "UpdateApproval",
			UpdatedAt: &timeNow,
		})

		if err != nil {
			asset.logger.Zap.Error(err)
			return false, err
		}

		// approval
		asset.approvalRepo.DeleteApprovals(request.ID)
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
				UpdatedAt:      &timeNow})
		if err != nil {
			asset.logger.Zap.Error(err)
			return false, err
		}
		return true, err
	default:
		return false, err
	}
}

// UpdatePublish implements AssetDefinition
func (asset AssetService) UpdatePublish(request *models.AssetsRequestUpdate) (status bool, err error) {
	switch request.TypeChecker {
	case "approve checker":
		// update checker
		_, err = asset.assetRepo.UpdateApproval(&models.AssetsUpdateApproval{
			ID: request.ID,
			// Published: ,
			// PublishDate: ,
			Status:    "01c", // approve checker
			Action:    "UpdatePublish",
			UpdatedAt: &timeNow,
		})

		if err != nil {
			asset.logger.Zap.Error(err)
			return false, err
		}

		// approval
		asset.approvalRepo.DeleteApprovals(request.ID)
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
				UpdatedAt: &timeNow})
		if err != nil {
			asset.logger.Zap.Error(err)
			return false, err
		}
		return true, err
	case "approve signer":

		if request.Type == "publish" {
			status, err = asset.assetRepo.UpdatePublish(&models.AssetsUpdatePublish{
				ID:            request.ID,
				Published:     true,
				PublishDate:   &timeNow,
				ExpiredDate:   lib.AddTime(0, 6, 0),
				LastMakerID:   request.LastMakerID,
				LastMakerDesc: request.LastMakerDesc,
				LastMakerDate: request.LastMakerDate,
				Action:        "UpdatePublish",
				UpdatedAt:     &timeNow,
			})

			if err != nil {
				asset.logger.Zap.Error(err)
				return false, err
			}

			// approval
			asset.approvalRepo.DeleteApprovals(request.ID)
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
					UpdatedAt:      &timeNow})
			if err != nil {
				asset.logger.Zap.Error(err)
				return false, err
			}

			// update, err := asset.assetRepo.UpdateDocumentID(&models.AssetsRequestUpdateElastic{
			// 	ID:         getOneAsset.ID,
			// 	DocumentID: getOneAsset.DocumentID,
			// })
			// if !update || err != nil {
			// 	return false, err
			// }

		} else if request.Type == "unpublish" {
			status, err = asset.assetRepo.UpdatePublish(&models.AssetsUpdatePublish{
				ID:            request.ID,
				Published:     false,
				LastMakerID:   request.LastMakerID,
				LastMakerDesc: request.LastMakerDesc,
				LastMakerDate: request.LastMakerDate,
				Action:        "UpdateUnPublish",
				UpdatedAt:     &timeNow,
			})

			if err != nil {
				asset.logger.Zap.Error(err)
				return false, err
			}

			// approval
			asset.approvalRepo.DeleteApprovals(request.ID)
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
					UpdatedAt:      &timeNow})
			if err != nil {
				asset.logger.Zap.Error(err)
				return false, err
			}

			// remove, err := asset.assetRepo.UpdateRemoveDocumentID(&models.AssetsRequestUpdateElastic{
			// 	ID:         getOneAsset.ID,
			// 	DocumentID: getOneAsset.DocumentID,
			// })

			// if !remove || err != nil {
			// 	return false, err
			// }
		}

		return true, err
	case "tolak checker":
		// update checker
		_, err = asset.assetRepo.UpdateApproval(&models.AssetsUpdateApproval{
			ID: request.ID,
			// Published: ,
			// PublishDate: ,
			Status:    "01b", // tolak checker
			Action:    "UpdatePublish",
			UpdatedAt: &timeNow,
		})

		if err != nil {
			asset.logger.Zap.Error(err)
			return false, err
		}

		// approval
		asset.approvalRepo.DeleteApprovals(request.ID)
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
				UpdatedAt: &timeNow})
		if err != nil {
			asset.logger.Zap.Error(err)
			return false, err
		}
		return true, err
		// update tolak checker
	case "tolak signer":
		_, err = asset.assetRepo.UpdateApproval(&models.AssetsUpdateApproval{
			ID: request.ID,
			// Published:   true,
			// PublishDate: &timeNow,
			// ExpiredDate: lib.AddTime(0, 6, 0),
			Status:    "01d", // tolak signer
			Action:    "UpdatePublish",
			UpdatedAt: &timeNow,
		})

		if err != nil {
			asset.logger.Zap.Error(err)
			return false, err
		}

		// approval
		asset.approvalRepo.DeleteApprovals(request.ID)
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
				UpdatedAt:      &timeNow})
		if err != nil {
			asset.logger.Zap.Error(err)
			return false, err
		}
		return true, err
	default:
		return false, err
	}
}

// Update implements AssetDefinition
func (asset AssetService) UpdateMaintain(request *models.AssetsRequest) (status bool, err error) {
	status, err = asset.assetRepo.UpdateMaintain(request)
	if err != nil {
		asset.logger.Zap.Error(err)
		return false, err
	}
	return status, err
}

// Delete implements AssetDefinition
func (asset AssetService) Delete(request *models.AssetsRequestUpdate) (status bool, err error) {
	status, err = asset.assetRepo.Delete(&models.AssetsUpdateDelete{
		ID:            request.ID,
		LastMakerID:   request.LastMakerID,
		LastMakerDesc: request.LastMakerDesc,
		LastMakerDate: request.LastMakerDate,
		Deleted:       true,
		Action:        "Delete",
		UpdatedAt:     &timeNow,
	})
	if err != nil {
		asset.logger.Zap.Error(err)
		return false, err
	}
	return status, err
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
