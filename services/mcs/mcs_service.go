package repMcs

import (
	"fmt"
	"os"
	"reflect"
	models "riskmanagement/models/mcs"

	lib "gitlab.com/golang-package-library/goresums"

	"gitlab.com/golang-package-library/logger"
)

type McsDefinition interface {
	GetUker(request *models.McsRequest) (response []models.UkerResponse, err error)
	GetPIC(request *models.McsRequest) (response []models.PICResponse, err error)
}

type McsService struct {
	logger logger.Logger
}

func NewMcsService(logger logger.Logger) McsDefinition {
	return McsService{
		logger: logger,
	}
}

// GetUker implements McsDefinition
func (mcs McsService) GetUker(request *models.McsRequest) (response []models.UkerResponse, err error) {
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

	options.BaseUrl = os.Getenv("OnegateURL") + "api/v1/client_auth/request_token"
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
	fmt.Println("===============================================")
	fmt.Println("====================JWT AUTH===================")

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
		Auth:   false,
	}

	auth = lib.Auth{
		Authorization: "Bearer " + dataResponseJwt,
	}

	dataResponse := []models.UkerResponse{}
	status := fmt.Sprint(statusResponseJwt)
	fmt.Println("status response", reflect.TypeOf(status))
	if status == "true" {
		mcs.logger.Zap.Info("Search Uker")
		options.BaseUrl = os.Getenv("OnegateURL") + "api/v1/uker/searchUker"
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

		fmt.Println("check interface or []interface", reflect.TypeOf(dataResponseSession))
		fmt.Println("dataResponseSession", fmt.Sprint(reflect.TypeOf(dataResponseSession)))

		if fmt.Sprint(reflect.TypeOf(dataResponseSession)) == "[]interface {}" {
			for _, data := range dataResponseSession.([]interface{}) {
				BRNAME := data.(map[string]interface{})["brname"]
				BRANCH := data.(map[string]interface{})["branch"]

				if fmt.Sprint(reflect.TypeOf(BRANCH)) == "float64" {
					subData := models.UkerResponse{
						BRNAME: BRNAME.(string),
						BRANCH: fmt.Sprint(int(BRANCH.(float64))),
					}
					dataResponse = append(dataResponse, subData)
				} else {
					subData := models.UkerResponse{
						BRNAME: BRNAME.(string),
						BRANCH: BRANCH.(string),
					}
					dataResponse = append(dataResponse, subData)

				}
			}
		}

	}

	fmt.Println("dataResponse", dataResponse)
	return dataResponse, err
}

// GetPIC implements McsDefinition
func (mcs McsService) GetPIC(request *models.McsRequest) (response []models.PICResponse, err error) {
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

	options.BaseUrl = os.Getenv("OnegateURL") + "api/v1/client_auth/request_token"
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
	fmt.Println("===============================================")
	fmt.Println("====================JWT AUTH===================")

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
		Auth:   false,
	}

	auth = lib.Auth{
		Authorization: "Bearer " + dataResponseJwt,
	}

	dataResponse := []models.PICResponse{}
	status := fmt.Sprint(statusResponseJwt)
	fmt.Println("status response", reflect.TypeOf(status))
	if status == "true" {
		mcs.logger.Zap.Info("Search Pekerja")
		options.BaseUrl = os.Getenv("OnegateURL") + "api/v1/pekerja/searchPekerja"
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

		fmt.Println("check interface or []interface", reflect.TypeOf(dataResponseSession))
		fmt.Println("dataResponseSession", fmt.Sprint(reflect.TypeOf(dataResponseSession)))

		if fmt.Sprint(reflect.TypeOf(dataResponseSession)) == "[]interface {}" {
			for _, data := range dataResponseSession.([]interface{}) {
				PERNR := data.(map[string]interface{})["PERNR"]
				HTEXT := data.(map[string]interface{})["HTEXT"]
				NAMA := data.(map[string]interface{})["NAMA"]

				if fmt.Sprint(reflect.TypeOf(PERNR)) == "float64" {
					subData := models.PICResponse{
						PERNR: fmt.Sprint(int(PERNR.(float64))),
						HTEXT: HTEXT.(string),
						NAMA:  NAMA.(string),
					}
					dataResponse = append(dataResponse, subData)
				} else {
					subData := models.PICResponse{
						PERNR: PERNR.(string),
						HTEXT: HTEXT.(string),
						NAMA:  NAMA.(string),
					}
					dataResponse = append(dataResponse, subData)
				}
			}
		}

	}

	fmt.Println("dataResponse", dataResponse)
	return dataResponse, err
}
