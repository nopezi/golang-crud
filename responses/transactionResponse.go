package responses

type TransactionResponse struct {
	Appname       string `json:"appname"`
	Object        string `json:"object"`
	Prefix        string `json:"prefix"`
	ExpiredDate   string `json:"expiredDate"`
	ReferenceCode string `json:"referenceCode"`
	Status        string `json:"status"`
}
