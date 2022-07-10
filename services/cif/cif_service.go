package cif

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/cif"
	"os"

	"gitlab.com/golang-package-library/logger"
)

type CifDefinition interface {
	InquiryCif(cifCode string) (response models.CifResponse, err error)
}
type CifService struct {
	logger logger.Logger
	// repository repository.CifDefinition
}

func NewCifService(logger logger.Logger) CifDefinition {
	return CifService{
		logger: logger,
	}
}

// Store implements CifDefinition
func (cif CifService) InquiryCif(cifCode string) (response models.CifResponse, err error) {
	// ===============================
	// Get Session API
	// jwt := ""

	cifData := models.CifData{
		CIF: cifCode,
	}
	headers := []lib.Header{
		{
			Key:   os.Getenv("X-ESB-SERVICE_ID_Key"),
			Value: os.Getenv("X-ESB-SERVICE_ID"),
		},
		{
			Key:   os.Getenv("X-ESB-CHANNEL_ID_Key"),
			Value: os.Getenv("X-ESB-CHANNEL_ID"),
		},
	}

	options := lib.Options{
		BaseUrl: os.Getenv("ESBCifService"),
		SSL:     false,
		Payload: models.CifRequest{
			Request: cifData,
		},
		Method: "POST",
		Auth:   false,
		Header: headers,
	}

	auth := lib.Auth{
		// Authorization: "Bearer " + jwt,
	}

	// ===============================
	// Search Pekerja
	cif.logger.Zap.Info("Search cif")

	options.BaseUrl = os.Getenv("ESBCifService")
	responseObjectSession, err := lib.AuthBearer(options, auth)
	if err != nil {
		cif.logger.Zap.Error(err)
		return response, err
	}

	statusResponseSession := responseObjectSession
	dataResponseSession := responseObjectSession["RESPONSE"]

	fmt.Println("statusResponseSession", statusResponseSession)
	fmt.Println("dataResponseSession", dataResponseSession)
	fmt.Println("==================================================")
	fmt.Println("Data CIF=====================================")

	subData := models.CifResponse{
		ERRORCODE:               dataResponseSession.(map[string]interface{})["ERROR_CODE"].(string),
		RESPONSECODE:            dataResponseSession.(map[string]interface{})["RESPONSE_CODE"].(string),
		RESPONSEMESSAGE:         dataResponseSession.(map[string]interface{})["RESPONSE_MESSAGE"].(string),
		NIK:                     dataResponseSession.(map[string]interface{})["NIK"].(string),
		CIF:                     dataResponseSession.(map[string]interface{})["CIF"].(string),
		BRANCH:                  dataResponseSession.(map[string]interface{})["BRANCH"].(string),
		CUSTNAME:                dataResponseSession.(map[string]interface{})["CUST_NAME"].(string),
		HPNO:                    dataResponseSession.(map[string]interface{})["HP_NO"].(string),
		EMAIL:                   dataResponseSession.(map[string]interface{})["EMAIL"].(string),
		ADDRESS1:                dataResponseSession.(map[string]interface{})["ADDRESS1"].(string),
		ADDRESS2:                dataResponseSession.(map[string]interface{})["ADDRESS2"].(string),
		ADDRESS3:                dataResponseSession.(map[string]interface{})["ADDRESS3"].(string),
		ADDRESS4:                dataResponseSession.(map[string]interface{})["ADDRESS4"].(string),
		ZIPCODE:                 dataResponseSession.(map[string]interface{})["ZIPCODE"].(string),
		RTNO:                    dataResponseSession.(map[string]interface{})["RT_NO"].(string),
		RWNO:                    dataResponseSession.(map[string]interface{})["RW_NO"].(string),
		PLACEOFBIRTH:            dataResponseSession.(map[string]interface{})["PLACE_OF_BIRTH"].(string),
		DATEOFBIRTH:             dataResponseSession.(map[string]interface{})["DATE_OF_BIRTH"].(string),
		SEX:                     dataResponseSession.(map[string]interface{})["SEX"].(string),
		CITIZENSHIP:             dataResponseSession.(map[string]interface{})["CITIZENSHIP"].(string),
		IDTYPE:                  dataResponseSession.(map[string]interface{})["ID_TYPE"].(string),
		IDNO:                    dataResponseSession.(map[string]interface{})["ID_NO"].(string),
		RELIGION:                dataResponseSession.(map[string]interface{})["RELIGION"].(string),
		MARTIALSTATUS:           dataResponseSession.(map[string]interface{})["MARTIAL_STATUS"].(string),
		MARITALSTATUS:           dataResponseSession.(map[string]interface{})["MARITAL_STATUS"].(string),
		MOTHERNAME:              dataResponseSession.(map[string]interface{})["MOTHER_NAME"].(string),
		NPWP:                    dataResponseSession.(map[string]interface{})["NPWP"].(string),
		TUJUANPEMBUKAANREKENING: dataResponseSession.(map[string]interface{})["TUJUAN_PEMBUKAAN_REKENING"].(string),
		TYPEOFWORK:              dataResponseSession.(map[string]interface{})["TYPE_OF_WORK"].(string),
		PENDIDIKANCODE:          dataResponseSession.(map[string]interface{})["PENDIDIKAN_CODE"].(string),
		PENDIDIKANDESC:          dataResponseSession.(map[string]interface{})["PENDIDIKAN_DESC"].(string),
		ALAMATSURATMENYURAT:     dataResponseSession.(map[string]interface{})["ALAMAT_SURAT_MENYURAT"].(string),
		CITY:                    dataResponseSession.(map[string]interface{})["CITY"].(string),
		KECAMATAN:               dataResponseSession.(map[string]interface{})["KECAMATAN"].(string),
		KELURAHAN:               dataResponseSession.(map[string]interface{})["KELURAHAN"].(string),
		OFFICEADDRESS:           dataResponseSession.(map[string]interface{})["OFFICE_ADDRESS"].(string),
		OFFICECITY:              dataResponseSession.(map[string]interface{})["OFFICE_CITY"].(string),
		OFFICEKELURAHAN:         dataResponseSession.(map[string]interface{})["OFFICE_KELURAHAN"].(string),
		OFFICEKECAMATAN:         dataResponseSession.(map[string]interface{})["OFFICE_KECAMATAN"].(string),
		OFFICENAME:              dataResponseSession.(map[string]interface{})["OFFICE_NAME"].(string),
		OFFICENOTELPON:          dataResponseSession.(map[string]interface{})["OFFICE_NO_TELPON"].(string),
		OFFICEZIPCODE:           dataResponseSession.(map[string]interface{})["OFFICE_ZIP_CODE"].(string),
		PENGHASILAN:             dataResponseSession.(map[string]interface{})["PENGHASILAN"].(string),
		PROVINSI:                dataResponseSession.(map[string]interface{})["PROVINSI"].(string),
		SUMBERUTAMA:             dataResponseSession.(map[string]interface{})["SUMBER_UTAMA"].(string),
		SUMBERUTAMADECS:         dataResponseSession.(map[string]interface{})["SUMBER_UTAMA_DECS"].(string),
		PEKERJAAN:               dataResponseSession.(map[string]interface{})["PEKERJAAN"].(string),
		PEKERJAANDECS:           dataResponseSession.(map[string]interface{})["PEKERJAAN_DECS"].(string),
		DAILYTRX:                dataResponseSession.(map[string]interface{})["DAILY_TRX"].(string),
		JABATAN:                 dataResponseSession.(map[string]interface{})["JABATAN"].(string),
		JABATANDECS:             dataResponseSession.(map[string]interface{})["JABATAN_DECS"].(string),
		CCCODE:                  dataResponseSession.(map[string]interface{})["CC_CODE"].(string),
	}

	fmt.Println("Data Response CIF", subData)
	return subData, err
}
