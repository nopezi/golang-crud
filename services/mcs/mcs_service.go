package repoMcs

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/mcs"
	"os"
	"reflect"

	"gitlab.com/golang-package-library/logger"
)

type McsDefinition interface {
	GetMcs(request *models.McsRequest) (response []models.McsResponse, err error)
}
type McsService struct {
	logger logger.Logger
	// repository repository.McsDefinition
}

func NewMcsService(logger logger.Logger) McsDefinition {
	return McsService{
		logger: logger,
	}
}

// Store implements McsDefinition
func (mcs McsService) GetMcs(request *models.McsRequest) (response []models.McsResponse, err error) {
	// ===============================
	// Get Session API
	fmt.Println("request", request)
	jwt := ""
	options := lib.Options{
		BaseUrl: os.Getenv("OnegateURL"),
		SSL:     false,
		Payload: models.McsRequest{
			Clientid:     os.Getenv("OnegateClientID"),
			Clientsecret: os.Getenv("OnegateSecret"),
			Keyword:      request.Keyword,
			Limit:        request.Limit,
			Offset:       request.Offset,
		},
		Method: "POST",
		Auth:   false,
	}

	auth := lib.Auth{
		Authorization: "Bearer " + jwt,
	}

	options.BaseUrl = os.Getenv("OnegateURL") + "/api/v1/client_auth/request_token"
	responseObjectJwt, err := lib.AuthBearer(options, auth)
	if err != nil {
		mcs.logger.Zap.Error(err)
		return response, err
	}
	fmt.Println("responseObjectJwt", responseObjectJwt)
	statusResponseJwt := responseObjectJwt["success"]
	dataResponseJwt := responseObjectJwt["message"].(map[string]interface{})["token"].(map[string]interface{})["token"].(string)

	fmt.Println("statusResponseJwt", statusResponseJwt)
	fmt.Println("dataResponseJwt", dataResponseJwt)
	fmt.Println("========================================================")
	fmt.Println("===================JWT AUTH=============================")
	// ===============================
	// End Of get JWT

	fmt.Println("request", request)
	jwt = ""
	options = lib.Options{
		BaseUrl: os.Getenv("OnegateURL"),
		SSL:     false,
		Payload: models.McsRequest{
			Keyword: request.Keyword,
			Limit:   request.Limit,
			Offset:  request.Offset,
		},
		Method: "POST",
		Auth:   true,
	}

	auth = lib.Auth{
		Authorization: "Bearer " + dataResponseJwt,
	}
	dataResponse := []models.McsResponse{}
	status := fmt.Sprint(statusResponseJwt)
	fmt.Println("status response", reflect.TypeOf(status))
	if status == "true" {
		// ===============================
		// Search Pekerja
		mcs.logger.Zap.Info("Search Pekerja")

		options.BaseUrl = os.Getenv("OnegateURL") + "/api/v1/pekerja/searchPekerja"
		responseObjectSession, err := lib.AuthBearer(options, auth)
		if err != nil {
			mcs.logger.Zap.Error(err)
			return response, err
		}
		fmt.Println(responseObjectSession)
		statusResponseSession := responseObjectSession["success"]
		dataResponseSession := responseObjectSession["message"]

		fmt.Println("statusResponseSession", statusResponseSession)
		fmt.Println("dataResponseSession", dataResponseSession)
		fmt.Println("==========================================================")
		fmt.Println("====================DATA MCS==============================")

		for _, data := range dataResponseSession.([]interface{}) {

			subData := models.McsResponse{
				PERNR: data.(map[string]interface{})["PERNR"].(string),
				HTEXT: data.(map[string]interface{})["HTEXT"].(string),
				NAMA:  data.(map[string]interface{})["NAMA"].(string),
			}
			dataResponse = append(dataResponse, subData)
		}
	}

	// fmt.Println("response", response)
	fmt.Println("dataResponse", dataResponse)
	return dataResponse, err
}
