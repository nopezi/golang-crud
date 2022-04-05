package controllers

import (
	"eform-gateway/api/services"
	"eform-gateway/lib"
	"eform-gateway/requests"
	"eform-gateway/responses"
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"
)

// TransactionController data type
type TransactionController struct {
	service services.TransactionService
	logger  lib.Logger
}

// NewTransactionController creates new Transaction controller
func NewTransactionController(TransactionService services.TransactionService, logger lib.Logger) TransactionController {
	return TransactionController{
		service: TransactionService,
		logger:  logger,
	}
}

// SaveTransaction saves the Transaction
func (u TransactionController) CreateTransaction(c *gin.Context) {
	referenceCode := responses.TransactionCreateResponse{}
	Transaction := requests.TransactionRequest{}

	if err := c.Bind(&Transaction); err != nil {
		u.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "98", "Validasi gagal: "+err.Error(), referenceCode)
		return
	}

	if len(Transaction.Prefix) < 5 {
		lib.ReturnToJson(c, 200, "98", "Validasi gagal: Prefix harus 5 karakter", referenceCode)
		return
	}

	re := regexp.MustCompile(`^(19|20)\d\d[- /.](0[1-9]|1[012])[- /.](0[1-9]|[12][0-9]|3[01])$`)
	checkmatch := re.MatchString(Transaction.ExpiredDate)

	if !checkmatch {
		lib.ReturnToJson(c, 200, "98", "Validasi gagal: Format Tanggal yyyy-mm-dd ", referenceCode)
		return
	}

	referenceCode, err := u.service.CreateTransaction(Transaction)

	if err != nil {
		u.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "04", "exc:"+err.Error(), referenceCode)
		return
	}
	fmt.Println("insert", referenceCode)
	lib.ReturnToJson(c, 200, "00", "Insert data berhasil", referenceCode)
}

// UpdateTransaction updates Transaction
func (u TransactionController) UpdateTransaction(c *gin.Context) {
	Transaction := requests.UpdateRequest{}
	response := responses.TransactionCreateResponse{}
	if err := c.Bind(&Transaction); err != nil {
		u.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "98", "Validasi gagal: "+err.Error(), response)
		return
	}

	response, err := u.service.UpdateTransaction(Transaction)

	if err != nil {
		u.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "02", "Data tidak ditemukan", response)
		return
	}

	lib.ReturnToJson(c, 200, "00", "Update data berhasil", response)
}

// UpdateTransaction updates Transaction
func (u TransactionController) InquiryTransaction(c *gin.Context) {
	Transaction := requests.InquiryRequest{}

	if err := c.Bind(&Transaction); err != nil {
		u.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "98", "Validasi gagal: "+err.Error(), "")
		return
	}

	response, status, err := u.service.InquiryTransaction(Transaction)

	if err != nil {
		u.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "04", "exc:"+err.Error(), response)
		return
	}

	if !status {
		lib.ReturnToJson(c, 200, "02", "Data tidak ditemukan", response)
		return
	}

	lib.ReturnToJson(c, 200, "00", "Inquiry data berhasil", response)
}
