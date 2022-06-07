package asset

import (
	"fmt"
	"infolelang/lib"
	approvals "infolelang/models/approvals"
	models "infolelang/models/assets"
	requestImage "infolelang/models/images"
	addressRepo "infolelang/repository/address"
	approvalRepo "infolelang/repository/approvals"
	assetRepo "infolelang/repository/asset"
	contactRepo "infolelang/repository/contacts"
	imageRepo "infolelang/repository/images"

	minio "gitlab.com/golang-package-library/minio"
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
	assetAccessPlace assetRepo.AssetAccessPlaceDefinition) AssetDefinition {
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
	dataAsset, err := asset.assetRepo.Store(request.ParseRequest(*request))
	if err != nil {
		asset.logger.Zap.Error(err)
		return err
	}
	fmt.Println("dataAsset", dataAsset)

	// address
	_, err = asset.addressRepo.Store(&request.Addresses)
	if err != nil {
		asset.logger.Zap.Error(err)
		return err
	}

	switch request.Type {
	case "FormB1":
		// buildingasset
		_, err = asset.buildingRepo.Store(&request.BuildingAssets)
		if err != nil {
			asset.logger.Zap.Error(err)
			return err
		}
	default:
		// vehicle asset
		_, err = asset.vehicleRepo.Store(&request.VehicleAssets)
		if err != nil {
			asset.logger.Zap.Error(err)
			return err
		}
	}

	for _, value := range request.Facilities {
		_, err = asset.assetFacility.Store(
			&models.AssetFacilitiesRequest{
				AssetID:      dataAsset.ID,
				FacilitiesID: value.ID,
			})

		if err != nil {
			asset.logger.Zap.Error(err)
			return err
		}
	}

	// asset_access_places
	for _, value := range request.AccessPlaces {
		_, err = asset.assetAccessPlace.Store(
			&models.AssetAccessPlacesRequest{
				AssetID:       dataAsset.ID,
				AccessPlaceID: value.ID,
			})

		if err != nil {
			asset.logger.Zap.Error(err)
			return err
		}
	}

	// contact
	_, err = asset.contactRepo.Store(&request.Contacts)
	if err != nil {
		asset.logger.Zap.Error(err)
		return err
	}

	// images
	for _, value := range request.Images {
		image, err := asset.imagesRepo.Store(
			&requestImage.ImagesRequest{
				Filename:  value.Filename,
				Path:      value.Path,
				Extension: value.Extension,
				Size:      value.Size,
			})

		if err != nil {
			asset.logger.Zap.Error(err)
			return err
		}

		_, err = asset.assetImage.Store(&models.AssetImagesRequest{
			AssetID: dataAsset.ID,
			ImageID: image.ID,
		})

		if err != nil {
			asset.logger.Zap.Error(err)
			return err
		}
	}

	// approval
	_, err = asset.approvalRepo.Store(
		&approvals.ApprovalsRequest{
			AssetID:        dataAsset.ID,
			CheckerID:      request.Approvals.CheckerID,
			CheckerDesc:    request.Approvals.CheckerDesc,
			CheckerComment: request.Approvals.CheckerComment,
			CheckerDate:    request.Approvals.CheckerDate,
			SignerID:       request.Approvals.SignerID,
			SignerDesc:     request.Approvals.SignerDesc,
			SignerComment:  request.Approvals.SignerComment,
			SignerDate:     request.Approvals.SignerDate})
	if err != nil {
		asset.logger.Zap.Error(err)
		return err
	}
	// fmt.Println(request)
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
