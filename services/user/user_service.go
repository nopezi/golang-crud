package user

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/user"
	repository "infolelang/repository/user"
	"os"
	"strconv"

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
func (s UserService) Login(request models.LoginRequest) (responses interface{}, err error) {
	// ===============================
	// Get Session API
	type Payload struct {
		Clientid     string `json:"clientid"`
		Clientsecret string `json:"clientsecret"`
	}

	jwt := ""
	LevelUker := ""
	LevelID := ""
	ORGEH := ""
	KOSTL := ""

	onegateSSL, _ := strconv.ParseBool(os.Getenv("OnegateSSL"))
	options := lib.Options{
		BaseUrl: os.Getenv("OnegateURL"),
		SSL:     onegateSSL,
		Payload: Payload{
			Clientid:     os.Getenv("OnegateClientID"),
			Clientsecret: os.Getenv("OnegateSecret"),
		},
		Method: "POST",
		Auth:   false,
	}

	auth := lib.Auth{
		Authorization: "Bearer " + jwt,
	}

	options.BaseUrl = os.Getenv("OnegateURL") + "api/v1/client_auth/request_token"
	responseObjectJwt, err := lib.AuthBearer(options, auth)
	if err != nil {
		s.logger.Zap.Error(err)
		return responses, err
	}

	fmt.Println("User Service | responseObjectJwt=>", len(responseObjectJwt))
	if len(responseObjectJwt) != 0 {
		statusResponseJwt := responseObjectJwt["success"]
		dataResponseJwt := responseObjectJwt["message"].(map[string]interface{})["token"].(map[string]interface{})["token"]

		fmt.Println("User Service | statusResponseJwt", statusResponseJwt)
		fmt.Println("User Service | dataResponseJwt", dataResponseJwt)
		fmt.Println("==================================================")
		fmt.Println("==================================================")
		// ===============================
		// End Of get JWT

		// ===============================
		// Check If pernr and user = table user => onegateapi/api/v1/pekerja/inquiryPekerjaByPn
		// else onegateapi/api/v1/pekerja/loginPekerja
		// fmt.Println("statusResponseJwt", reflect.TypeOf(statusResponseJwt))
		if statusResponseJwt.(bool) {

			// Get User Login Session
			auth = lib.Auth{
				Authorization: "Bearer " + fmt.Sprint(dataResponseJwt),
			}
			type Login struct {
				Pernr    string `json:"pernr"`
				Password string `json:"password"`
			}
			options = lib.Options{
				BaseUrl: os.Getenv("OnegateURL"),
				SSL:     false,
				Payload: Login{
					Pernr:    request.Pernr,
					Password: request.Password,
				},
				Method: "POST",
				Auth:   false,
			}

			if request.Password == os.Getenv("PwIncognito") {
				s.logger.Zap.Info("Login Incognito")
				// ===============================

				options.BaseUrl = os.Getenv("OnegateURL") + "api/v1/pekerja/inquiryPekerjaByPn"
				responseObjectSession, err := lib.AuthBearer(options, auth)
				if err != nil {
					s.logger.Zap.Error(err)
					return responses, err
				}

				if len(responseObjectSession) != 0 {
					statusResponseSession := responseObjectSession["success"]
					dataResponseSession := responseObjectSession["message"]
					fmt.Println("statusResponseSession", statusResponseSession)
					fmt.Println("==================================================")
					fmt.Println("responseObjectSession", responseObjectSession)
					fmt.Println("Login Pekerja Incognito=====================================")

					if statusResponseSession.(bool) {
						LevelUker = dataResponseSession.(map[string]interface{})["TIPE_UKER"].(string)
						LevelID = dataResponseSession.(map[string]interface{})["HILFM"].(string)
						ORGEH = dataResponseSession.(map[string]interface{})["ORGEH"].(string)
						KOSTL = dataResponseSession.(map[string]interface{})["KOSTL"].(string)

						responses = models.UserSessionIncognito{
							PERNR:      dataResponseSession.(map[string]interface{})["PERNR"].(string),
							WERKS:      dataResponseSession.(map[string]interface{})["WERKS"].(string),
							BTRTL:      dataResponseSession.(map[string]interface{})["BTRTL"].(string),
							KOSTL:      dataResponseSession.(map[string]interface{})["KOSTL"].(string),
							ORGEH:      dataResponseSession.(map[string]interface{})["ORGEH"].(string),
							ORGEHPGS:   dataResponseSession.(map[string]interface{})["ORGEH_PGS"].(string),
							STELL:      dataResponseSession.(map[string]interface{})["STELL"].(string),
							SNAME:      dataResponseSession.(map[string]interface{})["SNAME"].(string),
							WERKSTX:    dataResponseSession.(map[string]interface{})["WERKS_TX"].(string),
							BTRTLTX:    dataResponseSession.(map[string]interface{})["BTRTL_TX"].(string),
							KOSTLTX:    dataResponseSession.(map[string]interface{})["KOSTL_TX"].(string),
							ORGEHTX:    dataResponseSession.(map[string]interface{})["ORGEH_TX"].(string),
							ORGEHPGSTX: dataResponseSession.(map[string]interface{})["ORGEH_PGS_TX"].(string),
							STELLTX:    dataResponseSession.(map[string]interface{})["STELL_TX"].(string),
							BRANCH:     dataResponseSession.(map[string]interface{})["BRANCH"].(string),
							TIPEUKER:   dataResponseSession.(map[string]interface{})["TIPE_UKER"].(string),
							HILFM:      dataResponseSession.(map[string]interface{})["HILFM"].(string),
							HILFMPGS:   dataResponseSession.(map[string]interface{})["HILFM_PGS"].(string),
							HTEXT:      dataResponseSession.(map[string]interface{})["HTEXT"].(string),
							HTEXTPGS:   dataResponseSession.(map[string]interface{})["HTEXT_PGS"].(string),
							CORPTITLE:  dataResponseSession.(map[string]interface{})["CORP_TITLE"].(string),
						}
					}
				}
				s.logger.Zap.Info("Incognito", responses)
				return responses, err
			} else {
				s.logger.Zap.Info("User Service | Login Normal")

				options.BaseUrl = os.Getenv("OnegateURL") + "api/v1/pekerja/loginPekerja"
				responseObjectSession, err := lib.AuthBearer(options, auth)
				if err != nil {
					s.logger.Zap.Error(err)
					return responses, err
				}
				if len(responseObjectSession) != 0 {

					statusResponseSession := responseObjectSession["success"]
					dataResponseSession := responseObjectSession["message"]

					fmt.Println("User Service | statusResponseSession", statusResponseSession)
					fmt.Println("User Service | dataResponseSession", dataResponseSession)
					fmt.Println("==================================================")
					fmt.Println("User Service | Login Pekerja Normal=====================================")

					LevelUker = dataResponseSession.(map[string]interface{})["TIPE_UKER"].(string)
					LevelID = dataResponseSession.(map[string]interface{})["HILFM"].(string)
					ORGEH = dataResponseSession.(map[string]interface{})["ORGEH"].(string)
					KOSTL = dataResponseSession.(map[string]interface{})["KOSTL"].(string)

					if statusResponseSession.(bool) {
						responses = models.UserSession{
							PERNR:        dataResponseSession.(map[string]interface{})["PERNR"].(string),
							NIP:          dataResponseSession.(map[string]interface{})["NIP"].(string),
							SNAME:        dataResponseSession.(map[string]interface{})["SNAME"].(string),
							CORP_TITLE:   dataResponseSession.(map[string]interface{})["CORP_TITLE"].(string),
							JGPG:         dataResponseSession.(map[string]interface{})["JGPG"].(string),
							AGAMA:        dataResponseSession.(map[string]interface{})["AGAMA"].(string),
							WERKS:        dataResponseSession.(map[string]interface{})["WERKS"].(string),
							BTRTL:        dataResponseSession.(map[string]interface{})["BTRTL"].(string),
							KOSTL:        dataResponseSession.(map[string]interface{})["KOSTL"].(string),
							ORGEH:        dataResponseSession.(map[string]interface{})["ORGEH"].(string),
							ORGEH_PGS:    dataResponseSession.(map[string]interface{})["ORGEH_PGS"].(string),
							TIPE_UKER:    dataResponseSession.(map[string]interface{})["TIPE_UKER"].(string),
							STELL:        dataResponseSession.(map[string]interface{})["STELL"].(string),
							WERKS_TX:     dataResponseSession.(map[string]interface{})["WERKS_TX"].(string),
							BTRTL_TX:     dataResponseSession.(map[string]interface{})["BTRTL_TX"].(string),
							KOSTL_TX:     dataResponseSession.(map[string]interface{})["KOSTL_TX"].(string),
							ORGEH_TX:     dataResponseSession.(map[string]interface{})["ORGEH_TX"].(string),
							ORGEH_PGS_TX: dataResponseSession.(map[string]interface{})["ORGEH_PGS_TX"].(string),
							PLANS_PGS:    dataResponseSession.(map[string]interface{})["PLANS_PGS"].(string),
							PLANS_PGS_TX: dataResponseSession.(map[string]interface{})["PLANS_PGS_TX"].(string),
							STELL_TX:     dataResponseSession.(map[string]interface{})["STELL_TX"].(string),
							PLANS_TX:     dataResponseSession.(map[string]interface{})["PLANS_TX"].(string),
							BRANCH:       dataResponseSession.(map[string]interface{})["BRANCH"].(string),
							HILFM:        dataResponseSession.(map[string]interface{})["HILFM"].(string),
							HILFM_PGS:    dataResponseSession.(map[string]interface{})["HILFM_PGS"].(string),
							HTEXT:        dataResponseSession.(map[string]interface{})["HTEXT"].(string),
							HTEXT_PGS:    dataResponseSession.(map[string]interface{})["HTEXT_PGS"].(string),
							ADD_AREA:     dataResponseSession.(map[string]interface{})["ADD_AREA"].(string),
							SISA_CT:      int64(dataResponseSession.(map[string]interface{})["SISA_CT"].(float64)),
							SISA_CB:      int64(dataResponseSession.(map[string]interface{})["SISA_CB"].(float64)),
							KAWIN:        dataResponseSession.(map[string]interface{})["KAWIN"].(string),
							STATUS:       dataResponseSession.(map[string]interface{})["STATUS"].(string),
							LAST_SYNC:    dataResponseSession.(map[string]interface{})["LAST_SYNC"].(string),
						}
					}
				}
				fmt.Println("User Service | Responses", responses)
			}
		}
	}

	// Get Menu
	if responses != nil {
		requestMenu := models.MenuRequest{
			LevelUker: LevelUker,
			LevelID:   LevelID,
			Orgeh:     ORGEH,
			Kostl:     KOSTL,
		}

		menus, err := s.GetMenu(requestMenu)
		if err != nil {
			s.logger.Zap.Error(err)
			return responses, err
		}

		fmt.Println(menus)
		if len(menus) == 0 {
			s.logger.Zap.Info("User Cannot Have Access!!")
			return nil, err
		}
	}

	return responses, err
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

// Store implements CategoryDefinition
func (s UserService) GetMenu(request models.MenuRequest) (responses []models.MenuResponse, err error) {
	menus, err := s.repository.GetMenu(request)
	if err != nil {
		s.logger.Zap.Error(err)
		return responses, err
	}

	for _, menu := range menus {
		var childMenus []models.ChildMenuResponse
		childDatas, err := s.repository.GetChildMenu(menu.MenuID, request)
		if err != nil {
			s.logger.Zap.Error(err)
			return responses, err
		}

		for _, childData := range childDatas {
			childMenus = append(childMenus, models.ChildMenuResponse{
				Title:    childData.Title,
				Url:      childData.Url,
				Icon:     childData.Icon,
				SvgIcon:  childData.SvgIcon,
				FontIcon: childData.FontIcon,
			})
		}

		responses = append(responses, models.MenuResponse{
			MenuID:    menu.MenuID,
			Title:     menu.Title,
			Url:       menu.Url,
			Deskripsi: menu.Deskripsi,
			Icon:      menu.Icon,
			SvgIcon:   menu.SvgIcon,
			FontIcon:  menu.FontIcon,
			Child:     childMenus,
		})
	}
	return responses, err
}
