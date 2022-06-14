package models

type McsRequest struct {
	Keyword string `json:"keywor"`
	Limit   string `json:"limit"`
	Offset  string `json:"offset"`
}

type McsResponse struct {
	PERNR string `json:"PERNR"`
	HTEXT string `json:"HTEXT"`
	NAMA  string `json:"NAMA"`
}
