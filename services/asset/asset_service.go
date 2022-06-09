package asset

import (
	"fmt"
	"infolelang/lib"
	"os"
	"strings"

	"github.com/google/uuid"

	requestAddress "infolelang/models/addresses"
	requestApprovals "infolelang/models/approvals"
	requestBuilding "infolelang/models/building_assets"
	requestContact "infolelang/models/contacts"
	requestImage "infolelang/models/images"
	requestVehicle "infolelang/models/vehicle_assets"

	models "infolelang/models/assets"
	addressRepo "infolelang/repository/address"
	approvalRepo "infolelang/repository/approvals"
	assetRepo "infolelang/repository/asset"
	contactRepo "infolelang/repository/contacts"
	imageRepo "infolelang/repository/images"

	minio "gitlab.com/golang-package-library/minio"
)

var (
	timeNow = lib.GetTimeNow("timestime")
	UUID    = uuid.NewString()
)

type AssetDefinition interface {
	GetAll() (responses []models.AssetsResponse, err error)
	GetOne(id int64) (responses models.AssetsResponse, err error)
	Store(request *models.AssetsRequest) (err error)
	Update(request *models.AssetsRequest) (err error)
	Delete(id int64) (err error)
}
type AssetService struct {
	minio            minio.Minio
	logger           lib.Logger
	assetRepo        assetRepo.AssetDefinition
	addressRepo      addressRepo.AddressDefinition
	buildingRepo     assetRepo.BuildingAssetDefinition
	vehicleRepo      assetRepo.VehicleAssetDefinition
	contactRepo      contactRepo.ContactDefinition
	approvalRepo     approvalRepo.ApprovalDefinition
	imagesRepo       imageRepo.ImageDefinition
	assetImage       assetRepo.AssetImageDefinition
	assetFacility    assetRepo.AssetFacilityDefinition
	assetAccessPlace assetRepo.AssetAccessPlaceDefinition
}

func NewAssetService(
	minio minio.Minio,
	logger lib.Logger,
	assetRepo assetRepo.AssetDefinition,
	addressRepo addressRepo.AddressDefinition,
	buildingRepo assetRepo.BuildingAssetDefinition,
	vehicleRepo assetRepo.VehicleAssetDefinition,
	contactRepo contactRepo.ContactDefinition,
	approvalRepo approvalRepo.ApprovalDefinition,
	imagesRepo imageRepo.ImageDefinition,
	assetImage assetRepo.AssetImageDefinition,
	assetFacility assetRepo.AssetFacilityDefinition,
	assetAccessPlace assetRepo.AssetAccessPlaceDefinition,
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
		assetAccessPlace: assetAccessPlace,
	}
}

// GetAll implements AssetDefinition
func (asset AssetService) GetAll() (responses []models.AssetsResponse, err error) {
	return asset.assetRepo.GetAll()
}

// GetOne implements AssetDefinition
func (asset AssetService) GetOne(id int64) (responses models.AssetsResponse, err error) {
	return asset.assetRepo.GetOne(id)
}

// Store implements AssetDefinition
func (asset AssetService) Store(request *models.AssetsRequest) (err error) {
	// create assets
	bucket := os.Getenv("BUCKET_NAME")
	dataAsset, err := asset.assetRepo.Store(request.ParseCreate(*request))
	if err != nil {
		asset.logger.Zap.Error(err)
		return err
	}
	fmt.Println("dataAsset", dataAsset)

	// address

	_, err = asset.addressRepo.Store(
		&requestAddress.Addresses{
			AssetID:      dataAsset.ID,
			PostalcodeID: request.Addresses.PostalcodeID,
			Address:      request.Addresses.Address,
			Longitude:    request.Addresses.Longitude,
			Langitude:    request.Addresses.Langitude,
			CreatedAt:    &timeNow,
		})
	if err != nil {
		asset.logger.Zap.Error(err)
		return err
	}

	switch request.Type {
	case "FormB1":
		// buildingasset
		_, err = asset.buildingRepo.Store(&requestBuilding.BuildingAssets{
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
	default:
		// vehicle asset
		_, err = asset.vehicleRepo.Store(&requestVehicle.VehicleAssets{
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

	}

	for _, value := range request.Facilities {
		_, err = asset.assetFacility.Store(
			&models.AssetFacilities{
				AssetID:    dataAsset.ID,
				FacilityID: value.ID,
				CreatedAt:  &timeNow,
			})

		if err != nil {
			asset.logger.Zap.Error(err)
			return err
		}
	}

	// asset_access_places
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
	_, err = asset.contactRepo.Store(&requestContact.Contacts{
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

		image, err := asset.imagesRepo.Store(
			&requestImage.Images{
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

	// approval
	_, err = asset.approvalRepo.Store(
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
	fmt.Println(request)
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
