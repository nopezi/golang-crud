package user

import (
	"infolelang/lib"
	models "infolelang/models/user"
	repository "infolelang/repository/user"

	// minio "gitlab.com/golang-package-library/minio"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

// UserService service layer
type UserService struct {
	// minio      minio.Minio
	logger     lib.Logger
	repository repository.UserRepository
}

// NewUserService creates a new userservice
func NewUserService(
	// minio minio.Minio,
	logger lib.Logger,
	repository repository.UserRepository) UserService {
	return UserService{
		// minio:      minio,
		logger:     logger,
		repository: repository,
	}
}

// WithTrx delegates transaction to repository database
func (s UserService) WithTrx(trxHandle *gorm.DB) UserService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// GetOneUser gets one user
func (s UserService) GetOneUser(id uint) (models.User, error) {
	user, err := s.repository.GetOne(id)
	return user, err
}

// GetOneUser gets one user
func (s UserService) GetOneUserEmail(email *string) (models.User, error) {
	user, err := s.repository.GetUserByEmail(email)
	return user, err
}

// GetAllUser get all the user
func (s UserService) GetAllUser() ([]models.User, error) {
	users, err := s.repository.GetAll()
	return users, err
}

// CreateUser call to create the user
func (s UserService) CreateUser(user models.User) error {
	_, err := s.repository.Save(user)
	return err
}

// UpdateUser updates the user
func (s UserService) UpdateUser(id uint, user models.User) error {

	userDB, err := s.GetOneUser(id)
	if err != nil {
		return err
	}

	err = copier.Copy(&userDB, &user)
	if err != nil {
		return err
	}
	userDB.ID = id

	_, err = s.repository.Update(userDB)
	return err
}

// DeleteUser deletes the user
func (s UserService) DeleteUser(id uint) error {
	return s.repository.Delete(id)
}
