package models

// User model
type Transaction struct {
	Id            string
	Appname       string      `json:"appname"`
	Prefix        string      `json:"prefix"`
	Data          interface{} `json:"data"`
	ExpiredDate   string      `json:"expiredDate"`
	ReferenceCode string      `json:"referenceCode"`
	Status        string      `json:"status"`
}

type Data map[string]interface{}

// TableName gives table name of model
func (u Transaction) IndexName() string {
	return "transactions"
}
