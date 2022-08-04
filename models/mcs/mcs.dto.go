package models

type McsRequest struct {
	Clientid     string `json:"clientid"`
	Clientsecret string `json:"clientsecret"`
	Keyword      string `json:"keyword"`
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
}

type McsResponse struct {
	PERNR string `json:"PERNR"`
	HTEXT string `json:"HTEXT"`
	NAMA  string `json:"NAMA"`
}
