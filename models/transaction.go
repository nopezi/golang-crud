package models

// User model
type Transaction struct {
	Id            string
	Appname       string `json:"appname" binding:"required"`
	Object        string `json:"object" binding:"required"`
	Prefix        string `json:"prefix" binding:"required"`
	ExpiredDate   string `json:"expiredDate" binding:"required"`
	ReferenceCode string `json:"referenceCode"`
	Status        string `json:"status" binding:"required"`
}

// TableName gives table name of model
func (u Transaction) IndexName() string {
	return "transactions"
}
