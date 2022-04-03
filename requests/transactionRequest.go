package requests

type TransactionRequest struct {
	Appname       string `json:"appname"`
	Object        string `json:"object"`
	Prefix        string `json:"prefix"`
	ExpiredDate   string `json:"expiredDate"`
	ReferenceCode string `json:"referenceCode"`
	Status        string `json:"status"`
}

type UpdateRequest struct {
	ReferenceCode string `json:"referenceCode"`
}
