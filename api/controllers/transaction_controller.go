package controllers

import (
	"eform-gateway/api/services"
	"eform-gateway/lib"
	"eform-gateway/requests"

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

// // GetOneTransaction gets one Transaction
// func (u TransactionController) GetOneTransaction(c *gin.Context) {
// 	paramID := c.Param("id")

// 	id, err := strconv.Atoi(paramID)
// 	if err != nil {
// 		u.logger.Zap.Error(err)
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}
// 	Transaction, err := u.service.GetOneTransaction(uint(id))

// 	if err != nil {
// 		u.logger.Zap.Error(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"data": Transaction,
// 	})

// }

// // GetTransaction gets the Transaction
// func (u TransactionController) GetTransaction(c *gin.Context) {
// 	Transactions, err := u.service.GetAllTransaction()
// 	if err != nil {
// 		u.logger.Zap.Error(err)
// 	}
// 	c.JSON(200, gin.H{"data": Transactions})
// }

// SaveTransaction saves the Transaction
func (u TransactionController) SaveTransaction(c *gin.Context) {
	referenceCode := ""
	Transaction := requests.TransactionRequest{}

	if err := c.Bind(&Transaction); err != nil {
		u.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "98", "Validasi parameter gagal: "+err.Error(), referenceCode)
		return
	}

	referenceCode, err := u.service.CreateTransaction(Transaction)

	if err != nil {
		u.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "04", "exc:"+err.Error(), referenceCode)
		return
	}

	lib.ReturnToJson(c, 200, "00", "Insert data berhasil", referenceCode)
}

// UpdateTransaction updates Transaction
func (u TransactionController) UpdateTransaction(c *gin.Context) {
	Transaction := requests.UpdateRequest{}
	var response bool
	if err := c.Bind(&Transaction); err != nil {
		u.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "98", "Validasi parameter gagal: "+err.Error(), response)
		return
	}

	response, err := u.service.UpdateTransaction(Transaction)

	if err != nil {
		u.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "04", "exc:"+err.Error(), response)
		return
	}

	lib.ReturnToJson(c, 200, "00", "Update data berhasil", response)
}

// // SaveTransactionWOTrx saves the Transaction without transaction for comparision
// func (u TransactionController) SaveTransactionWOTrx(c *gin.Context) {
// 	Transaction := models.Transaction{}

// 	if err := c.Bind(&Transaction); err != nil {
// 		u.logger.Zap.Error(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error bind JSON": err.Error(),
// 		})
// 		return
// 	}

// 	if err := u.service.CreateTransaction(Transaction); err != nil {
// 		u.logger.Zap.Error(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H{"data": "Transaction created"})
// }

// // DeleteTransaction deletes Transaction
// func (u TransactionController) DeleteTransaction(c *gin.Context) {
// 	paramID := c.Param("id")

// 	id, err := strconv.Atoi(paramID)
// 	if err != nil {
// 		u.logger.Zap.Error(err)
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}

// 	if err := u.service.DeleteTransaction(uint(id)); err != nil {
// 		u.logger.Zap.Error(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H{"data": "Transaction deleted"})
// }
