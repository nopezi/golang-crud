package user

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/user"
	repository "infolelang/repository/user"
	"os"

	// minio "gitlab.com/golang-package-library/minio"

	"github.com/jinzhu/copier"
	// "gitlab.com/golang-package-library/goresums"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

// UserService service layer
type UserService struct {
	// minio      minio.Minio
	logger     logger.Logger
	repository repository.UserRepository
}

// NewUserService creates a new userservice
func NewUserService(
	// minio minio.Minio,
	logger logger.Logger,
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

// Get User Login
// AuthBearer(options Options, auth Auth)
func (s UserService) Login(request models.Login) (response bool, err error) {
	// ===============================
	// Get Session API
	type Payload struct {
		clientid     string
		clientsecret string
	}

	jwt := ""
	options := lib.Options{
		BaseUrl: os.Getenv("OnegateURL"),
		SSL:     false,
		Payload: Payload{
			clientid:     os.Getenv("OnegateClientID"),
			clientsecret: os.Getenv("OnegateSecret"),
		},
		Method: "POST",
		Auth:   false,
	}

	auth := lib.Auth{
		Authorization: "Bearer " + jwt,
	}

	options.BaseUrl = os.Getenv("OnegateURL") + "onegateapi/api/v1/client_auth/request_token"
	responseObjectJwt, err := lib.AuthBearer(options, auth)
	if err != nil {
		s.logger.Zap.Error(err)
		return false, err
	}

	statusResponseJwt := responseObjectJwt["success"]
	dataResponseJwt := responseObjectJwt["message"].(map[string]interface{})["token"].(map[string]interface{})["token"]

	fmt.Println("statusResponseJwt", statusResponseJwt)
	fmt.Println("dataResponseJwt", dataResponseJwt)

	// ===============================
	// Get User Login Session
	auth = lib.Auth{
		Authorization: "Bearer " + fmt.Sprint(dataResponseJwt),
	}

	options.BaseUrl = os.Getenv("OnegateURL") + "onegateapi/api/v1/pekerja/loginPekerja"
	responseObjectSession, err := lib.AuthBearer(options, auth)
	if err != nil {
		s.logger.Zap.Error(err)
		return false, err
	}

	statusResponseSession := responseObjectSession["success"]
	dataResponseSession := responseObjectSession["message"]
	// .(map[string]interface{})["token"].(map[string]interface{})["token"]

	fmt.Println("statusResponseSession", statusResponseSession)
	fmt.Println("dataResponseSession", dataResponseSession)
	fmt.Println("========================")

	session := models.UserSession{
		// (map[string]interface{})["Name"].(string)
		PERNR: lib.RemoveNull(dataResponseSession, "PERNR").(string),
		// dataResponseSession.(map[string]interface{})["PERNR"].(string),
		// NIP:        dataResponseSession.(map[string]interface{})["NIP"].(string),
		// SNAME:      dataResponseSession.(map[string]interface{})["SNAME"].(string),
		// WERKS:      dataResponseSession.(map[string]interface{})["WERKS"].(string),
		// BTRTL:      dataResponseSession.(map[string]interface{})["BTRTL"].(string),
		// KOSTL:      dataResponseSession.(map[string]interface{})["KOSTL"].(string),
		// ORGEH:      dataResponseSession.(map[string]interface{})["ORGEH"].(string),
		// STELL:      dataResponseSession.(map[string]interface{})["STELL"].(string),
		// WERKSTX:    dataResponseSession.(map[string]interface{})["WERKS_TX"].(string),
		// BTRTLTX:    dataResponseSession.(map[string]interface{})["BTRTL_TX"].(string),
		// KOSTLTX:    dataResponseSession.(map[string]interface{})["KOSTL_TX"].(string),
		// ORGEHTX:    dataResponseSession.(map[string]interface{})["ORGEH_TX"].(string),
		// STELLTX:    dataResponseSession.(map[string]interface{})["STELL_TX"].(string),
		// PLANSTX:    dataResponseSession.(map[string]interface{})["PLANS_TX"].(string),
		// JGPG:       dataResponseSession.(map[string]interface{})["JGPG"].(string),
		// ORGEHPGS:   dataResponseSession.(map[string]interface{})["ORGEH_PGS"].(string),
		// PLANSPGS:   dataResponseSession.(map[string]interface{})["PLANS_PGS"].(string),
		// ORGEHPGSTX: dataResponseSession.(map[string]interface{})["ORGEH_PGS_TX"].(string),
		// PLANSPGSTX: dataResponseSession.(map[string]interface{})["PLANS_PGS_TX"].(string),
		// SISACT:     dataResponseSession.(map[string]interface{})["SISA_CT"].(string),
		// SISACB:     dataResponseSession.(map[string]interface{})["SISA_CB"].(string),
		// AGAMA:      dataResponseSession.(map[string]interface{})["AGAMA"].(string),
		// TIPEUKER:   dataResponseSession.(map[string]interface{})["TIPE_UKER"].(string),
		// ADDAREA:    dataResponseSession.(map[string]interface{})["ADD_AREA"].(string),
		// PERSG:      dataResponseSession.(map[string]interface{})["PERSG"].(string),
		// PERSK:      dataResponseSession.(map[string]interface{})["PERSK"].(string),
		// STATUS:     dataResponseSession.(map[string]interface{})["STATUS"].(string),
		// BRANCH:     dataResponseSession.(map[string]interface{})["BRANCH"].(string),
		// HILFM:      dataResponseSession.(map[string]interface{})["HILFM"].(string),
		// HTEXT:      dataResponseSession.(map[string]interface{})["HTEXT"].(string),
		// HILFMPGS:   dataResponseSession.(map[string]interface{})["HILFM_PGS"].(string),
		// HTEXTPGS:   dataResponseSession.(map[string]interface{})["HTEXT_PGS"].(string),
		// KAWIN:      dataResponseSession.(map[string]interface{})["KAWIN"].(string),

		// WERKSPGS:   dataResponseSession.(map[string]interface{})["WERKS_PGS"].(string),
		// BTRTLPGS: dataResponseSession.(map[string]interface{})["BTRTL_PGS"].(string),
		// KOSTLPGS: dataResponseSession.(map[string]interface{})["KOSTL_PGS"].(string),
		KOSTLPGS: lib.RemoveNull(dataResponseSession, "KOSTL_PGS").(string),
	}
	fmt.Println("session", session)
	pernr := lib.RemoveNull(dataResponseSession, "PERNR")
	fmt.Println(pernr)
	// fmt.Println(lib.RemoveNull(dataResponseSession, "KOSTL_PGS").(string))

	return response, err
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
