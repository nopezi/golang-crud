package repoBanner

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/banner"
	repository "infolelang/repository/banner"
	"os"
	"strings"

	imageRepo "infolelang/repository/images"

	requestImage "infolelang/models/images"

	"github.com/google/uuid"
	"gitlab.com/golang-package-library/logger"
	"gitlab.com/golang-package-library/minio"
)

var (
	timeNow = lib.GetTimeNow("timestime")
	UUID    = uuid.NewString()
)

type BannerDefinition interface {
	GetAll() (responses []models.BannerImageResponse, err error)
	Store(request *models.BannerRequest) (status bool, err error)
	Delete(request models.BannerImageRequest) (status bool, err error)
}
type BannerService struct {
	minio      minio.Minio
	logger     logger.Logger
	repository repository.BannerDefinition
	imagesRepo imageRepo.ImageDefinition
}

func NewBannerService(
	minio minio.Minio,
	logger logger.Logger,
	repository repository.BannerDefinition,
	imagesRepo imageRepo.ImageDefinition,
) BannerDefinition {
	return BannerService{
		minio:      minio,
		logger:     logger,
		repository: repository,
		imagesRepo: imagesRepo,
	}
}

// GetAll implements BannerDefinition
func (banner BannerService) GetAll() (responses []models.BannerImageResponse, err error) {
	return banner.repository.GetAllBannerImage(1)
}

// Store implements BannerDefinition
func (banner BannerService) Store(request *models.BannerRequest) (status bool, err error) {
	bucket := os.Getenv("BUCKET_NAME")
	banners, err := banner.repository.Store(&models.Banner{
		ID:        1,
		Name:      "banner",
		CreatedAt: &timeNow,
		UpdatedAt: &timeNow,
	})
	for _, value := range request.Images {

		var destinationPath string
		bucketExist := banner.minio.BucketExist(banner.minio.Client(), bucket)

		pathSplit := strings.Split(value.Path, "/")
		sourcePath := fmt.Sprint(value.Path)
		destinationPath = pathSplit[1] + "/banners/" +
			lib.GetTimeNow("year") + "/" +
			lib.GetTimeNow("month") + "/" +
			lib.GetTimeNow("day") + "/" +
			pathSplit[2] + "/" +
			value.Filename

		if bucketExist {
			fmt.Println("Exist")
			fmt.Println(bucket)
			fmt.Println(destinationPath)
			banner.minio.CopyObject(banner.minio.Client(), bucket, sourcePath, bucket, destinationPath)

		} else {
			fmt.Println("Not Exist")
			fmt.Println(bucket)
			fmt.Println(destinationPath)
			banner.minio.MakeBucket(banner.minio.Client(), bucket, "")
			banner.minio.CopyObject(banner.minio.Client(), bucket, sourcePath, bucket, destinationPath)
		}

		image, err := banner.imagesRepo.Store(&requestImage.Images{
			Filename:  value.Filename,
			Path:      destinationPath,
			Extension: value.Extension,
			Size:      value.Size,
			CreatedAt: &timeNow,
		})

		if err != nil {
			banner.logger.Zap.Error(err)
			return false, err
		}

		_, err = banner.repository.StoreBannerImage(&models.BannerRequest{
			BannerID:      banners.ID,
			BannerImageID: image.ID,
		})

		if err != nil {
			banner.logger.Zap.Error(err)
			return false, err
		}
	}
	return true, err
}

// Delete implements BannerDefinition
func (banner BannerService) Delete(request models.BannerImageRequest) (status bool, err error) {
	status, err = banner.repository.Delete(request)
	if err != nil {
		banner.logger.Zap.Error(err)
		return false, err
	}

	err = banner.imagesRepo.Delete(request.ImageID)
	if err != nil {
		banner.logger.Zap.Error(err)
		return false, err
	}
	return true, err
}
