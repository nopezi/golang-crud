package content

import (
	"crud/lib"
	models "crud/models/content"
	"fmt"

	"gitlab.com/golang-package-library/logger"
	// "gorm.io/gorm"
)

type ContentRepository struct {
	db     lib.Database
	logger logger.Logger
}

func NewContentRepository(
	db lib.Database,
	logger logger.Logger,
) ContentRepository {
	return ContentRepository{
		db:     db,
		logger: logger,
	}
}

func (r ContentRepository) GetAll() (datas []models.Content, err error) {
	return datas, r.db.DB.Find(&datas).Error
}

func (r ContentRepository) Save(content models.Content) (models.Content, error) {
	fmt.Println("masok repository >> ", content)
	return content, r.db.DB.Create(&content).Error
}

func (r ContentRepository) Update(content models.Content) (models.Content, error) {
	return content, r.db.DB.Save(&content).Error
}

func (r ContentRepository) Delete(id uint) error {
	return r.db.DB.Where("id = ?", id).Delete(&models.Content{}).Error
}
