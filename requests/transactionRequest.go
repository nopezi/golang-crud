package requests

type TransactionRequest struct {
	Appname       string `json:"appname" binding:"required"`
	Object        string `json:"object" binding:"required"`
	Prefix        string `json:"prefix" binding:"required"`
	ExpiredDate   string `json:"expiredDate" binding:"required"`
	ReferenceCode string `json:"referenceCode"`
	Status        string `json:"status" binding:"required"`
}

type UpdateRequest struct {
	ReferenceCode string `json:"referenceCode" binding:"required"`
}

type InquiryRequest struct {
	ReferenceCode string `json:"referenceCode" binding:"required"`
	Status        string `json:"status" binding:"required"`
}
