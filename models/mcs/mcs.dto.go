package models

type McsRequest struct {
	Clientid     string `json:"clientid"`
	Clientsecret string `json:"clientsecret"`
	Keyword      string `json:"keyword"`
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
}

type PICResponse struct {
	PERNR string `json:"PERNR"`
	HTEXT string `json:"HTEXT"`
	NAMA  string `json:"NAMA"`
}

type PICResponseString struct {
	PERNR string `json:"PERNR"`
	HTEXT string `json:"HTEXT"`
	NAMA  string `json:"NAMA"`
}

type UkerResponse struct {
	BRNAME string `json:"BRNAME"`
	BRANCH string `json:"BRANCH"`
}

type UkerResponseString struct {
	BRNAME string `json:"BRNAME"`
	BRANCH string `json:"BRANCH"`
}
