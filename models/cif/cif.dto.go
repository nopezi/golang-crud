package models

type CifRequest struct {
	Request CifData `json:"REQUEST"`
}
type CifData struct {
	CIF string `json:"CIF"`
}

type CifResponse struct {
	ERRORCODE               string `json:"ERROR_CODE"`
	RESPONSECODE            string `json:"RESPONSE_CODE"`
	RESPONSEMESSAGE         string `json:"RESPONSE_MESSAGE"`
	NIK                     string `json:"NIK"`
	CIF                     string `json:"CIF"`
	BRANCH                  string `json:"BRANCH"`
	CUSTNAME                string `json:"CUST_NAME"`
	HPNO                    string `json:"HP_NO"`
	EMAIL                   string `json:"EMAIL"`
	ADDRESS1                string `json:"ADDRESS1"`
	ADDRESS2                string `json:"ADDRESS2"`
	ADDRESS3                string `json:"ADDRESS3"`
	ADDRESS4                string `json:"ADDRESS4"`
	ZIPCODE                 string `json:"ZIPCODE"`
	RTNO                    string `json:"RT_NO"`
	RWNO                    string `json:"RW_NO"`
	PLACEOFBIRTH            string `json:"PLACE_OF_BIRTH"`
	DATEOFBIRTH             string `json:"DATE_OF_BIRTH"`
	SEX                     string `json:"SEX"`
	CITIZENSHIP             string `json:"CITIZENSHIP"`
	IDTYPE                  string `json:"ID_TYPE"`
	IDNO                    string `json:"ID_NO"`
	RELIGION                string `json:"RELIGION"`
	MARTIALSTATUS           string `json:"MARTIAL_STATUS"`
	MARITALSTATUS           string `json:"MARITAL_STATUS"`
	MOTHERNAME              string `json:"MOTHER_NAME"`
	NPWP                    string `json:"NPWP"`
	TUJUANPEMBUKAANREKENING string `json:"TUJUAN_PEMBUKAAN_REKENING"`
	TYPEOFWORK              string `json:"TYPE_OF_WORK"`
	PENDIDIKANCODE          string `json:"PENDIDIKAN_CODE"`
	PENDIDIKANDESC          string `json:"PENDIDIKAN_DESC"`
	ALAMATSURATMENYURAT     string `json:"ALAMAT_SURAT_MENYURAT"`
	CITY                    string `json:"CITY"`
	KECAMATAN               string `json:"KECAMATAN"`
	KELURAHAN               string `json:"KELURAHAN"`
	OFFICEADDRESS           string `json:"OFFICE_ADDRESS"`
	OFFICECITY              string `json:"OFFICE_CITY"`
	OFFICEKELURAHAN         string `json:"OFFICE_KELURAHAN"`
	OFFICEKECAMATAN         string `json:"OFFICE_KECAMATAN"`
	OFFICENAME              string `json:"OFFICE_NAME"`
	OFFICENOTELPON          string `json:"OFFICE_NO_TELPON"`
	OFFICEZIPCODE           string `json:"OFFICE_ZIP_CODE"`
	PENGHASILAN             string `json:"PENGHASILAN"`
	PROVINSI                string `json:"PROVINSI"`
	SUMBERUTAMA             string `json:"SUMBER_UTAMA"`
	SUMBERUTAMADECS         string `json:"SUMBER_UTAMA_DECS"`
	PEKERJAAN               string `json:"PEKERJAAN"`
	PEKERJAANDECS           string `json:"PEKERJAAN_DECS"`
	DAILYTRX                string `json:"DAILY_TRX"`
	JABATAN                 string `json:"JABATAN"`
	JABATANDECS             string `json:"JABATAN_DECS"`
	CCCODE                  string `json:"CC_CODE"`
}
