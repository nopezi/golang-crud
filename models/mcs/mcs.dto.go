package models

type McsRequest struct {
	Keyword string `json:"keywor"`
	Limit   int    `json:"limit"`
	Offset  int    `json:"offset"`
}

type McsResponse struct {
	PERNR string `json:"PERNR"`
	HTEXT string `json:"HTEXT"`
	NAMA  string `json:"NAMA"`
}
