package contents

import (
	models "crud/models/content"
	repository "crud/repository/content"

	"gitlab.com/golang-package-library/logger"
)

type ContentService struct {
	logger     logger.Logger
	repository repository.ContentRepository
}

func NewContentService(
	logger logger.Logger,
	repository repository.ContentRepository,
) ContentService {
	return ContentService{
		logger:     logger,
		repository: repository,
	}
}

func (s ContentService) GetAll() (responses []models.Content, err error) {
	responses, err = s.repository.GetAll()
	return responses, err
}

func (s ContentService) CreateContent(content models.Content) error {
	_, err := s.repository.Save(content)
	return err
}

func (s ContentService) UpdateContent(content models.Content) error {
	_, err := s.repository.Update(content)
	return err
}

func (s ContentService) DeleteContent(id uint) error {
	return s.repository.Delete(id)
}
