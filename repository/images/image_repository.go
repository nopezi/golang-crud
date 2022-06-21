package images

import (
	"infolelang/lib"
	models "infolelang/models/images"
	"time"

	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type ImageDefinition interface {
	GetAll() (responses []models.ImagesResponse, err error)
	GetOne(id int64) (responses models.ImagesResponse, err error)
	GetOneAsset(id int64) (responses []models.ImagesResponses, err error)
	Store(request *models.Images) (responses *models.Images, err error)
	Update(request *models.ImagesRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) ImageRepository
}
type ImageRepository struct {
	db      lib.Database
	db2     lib.Databases
	elastic elastic.Elasticsearch
	logger  logger.Logger
	timeout time.Duration
}

func NewImageReporitory(
	db lib.Database,
	db2 lib.Databases,
	elastic elastic.Elasticsearch,
	logger logger.Logger) ImageDefinition {
	return ImageRepository{
		db:      db,
		db2:     db2,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements ImageDefinition
func (image ImageRepository) WithTrx(trxHandle *gorm.DB) ImageRepository {
	if trxHandle == nil {
		image.logger.Zap.Error("transaction Database not found in gin context. ")
		return image
	}
	image.db.DB = trxHandle
	return image
}

// GetAll implements ImageDefinition
func (image ImageRepository) GetAll() (responses []models.ImagesResponse, err error) {
	return responses, image.db.DB.Find(&responses).Error
}

// GetOne implements ImageDefinition
func (image ImageRepository) GetOne(id int64) (responses models.ImagesResponse, err error) {
	return responses, image.db.DB.Where("id = ?", id).Find(&responses).Error
}

// GetOneAsset implements ImageDefinition
func (image ImageRepository) GetOneAsset(id int64) (responses []models.ImagesResponses, err error) {
	// return responses, image.db.DB.Where("id = ?", id).Find(&responses).Error
	// rows, err := db.Model(&User{}).Where("name = ?", "jinzhu").Select("name, age, email").Rows() // (*sql.Rows, error)
	// defer rows.Close()

	// var user User
	// for rows.Next() {
	//   // ScanRows scan a row into user
	//   db.ScanRows(rows, &user)

	//   // do something
	// }
	rows, err := image.db.DB.Raw("select ai.id id, i.filename filename, i.`path` path , i.extension extension, i.`size` size  from asset_images ai join images i on ai.image_id = i.id where ai.asset_id  = ? ", id).Rows()

	defer rows.Close()

	var images models.ImagesResponses
	for rows.Next() {
		// ScanRows scan a row into user
		image.db.DB.ScanRows(rows, &images)
		responses = append(responses, images)
		// do something
	}
	return responses, err
}

// Store implements ImageDefinition
func (image ImageRepository) Store(request *models.Images) (responses *models.Images, err error) {
	return request, image.db.DB.Save(&request).Error
}

// Update implements ImageDefinition
func (image ImageRepository) Update(request *models.ImagesRequest) (responses bool, err error) {
	return true, image.db.DB.Save(&request).Error
}

// Delete implements ImageDefinition
func (image ImageRepository) Delete(id int64) (err error) {
	return image.db.DB.Where("id = ?", id).Delete(&models.ImagesResponse{}).Error
}
