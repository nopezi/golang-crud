package models

// User model
type Transaction struct {
	Appname       string `json:"appname"`
	Object        string `json:"object"`
	Prefix        string `json:"prefix"`
	ExpiredDate   string `json:"expiredDate"`
	ReferenceCode string `json:"referenceCode"`
	Status        string `json:"status"`
}

// TableName gives table name of model
func (u Transaction) IndexName() string {
	return "transactions"
}
