package access_places

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/banner"
	"time"

	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type BannerDefinition interface {
	GetAll() (responses []models.Banner, err error)
	GetAllBannerImage(bannerID int64) (responses []models.BannerImageResponse, err error)
	Store(request *models.Banner) (responses models.Banner, err error)
	StoreBannerImage(request *models.BannerRequest) (responses bool, err error)
	Delete(request models.BannerImageRequest) (status bool, err error)
	WithTrx(trxHandle *gorm.DB) BannerRepository
}
type BannerRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	elastic elastic.Elasticsearch
	logger  logger.Logger
	timeout time.Duration
}

func NewBannerReporitory(
	db lib.Database,
	dbRaw lib.Databases,
	elastic elastic.Elasticsearch,
	logger logger.Logger) BannerDefinition {
	return BannerRepository{
		db:      db,
		dbRaw:   dbRaw,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements BannerDefinition
func (banner BannerRepository) WithTrx(trxHandle *gorm.DB) BannerRepository {
	if trxHandle == nil {
		banner.logger.Zap.Error("transaction Database not found in gin context. ")
		return banner
	}
	banner.db.DB = trxHandle
	return banner
}

// GetAll implements BannerDefinition
func (banner BannerRepository) GetAll() (responses []models.Banner, err error) {
	return responses, banner.db.DB.Find(&responses).Error
}

// GetAll implements BannerDefinition
func (banner BannerRepository) GetAllBannerImage(bannerID int64) (responses []models.BannerImageResponse, err error) {
	// return responses, banner.db.DB.Where("banner_id = ?", bannerID).Find(&responses).Error
	rows, err := banner.db.DB.Raw(`
				SELECT bi.id benner_image_id,bi.banner_id, 
				bi.image_id, i.filename, 
				i.path ,i.extension,i.size  
				from banner_images bi
				LEFT JOIN images i on bi.image_id = i.id 
				LEFT JOIN banners b on bi.banner_id = b.id
				WHERE b.id = ? `, bannerID).Rows()

	defer rows.Close()

	var bannerImage models.BannerImageResponse
	for rows.Next() {
		banner.db.DB.ScanRows(rows, &bannerImage)
		responses = append(responses, bannerImage)
	}
	return responses, err
}

// Store implements BannerDefinition
func (banner BannerRepository) Store(request *models.Banner) (responses models.Banner, err error) {
	return *request, banner.db.DB.Save(&request).Error
}

// StoreBannerImage implements BannerDefinition
func (banner BannerRepository) StoreBannerImage(request *models.BannerRequest) (responses bool, err error) {
	err = banner.db.DB.Save(&models.BannerImage{
		BannerID:      request.BannerID,
		BannerImageID: request.BannerImageID,
	}).Error
	fmt.Println(err)
	return true, err
}

// Delete implements BannerDefinition
func (banner BannerRepository) Delete(request models.BannerImageRequest) (status bool, err error) {
	err = banner.db.DB.Where("image_id = ?", request.ImageID).Delete(&models.BannerImage{}).Error
	if err != nil {
		banner.logger.Zap.Error(err)
		return false, err
	}

	return true, err

}
