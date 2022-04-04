package services

import (
	"eform-gateway/api/repository"
	"eform-gateway/lib"
	"eform-gateway/models"
	"eform-gateway/requests"
	"eform-gateway/responses"
	"fmt"
)

// TransactionService service layer
type TransactionService struct {
	logger     lib.Logger
	repository repository.TransactionRepository
}

// NewTransactionService creates a new Transactionservice
func NewTransactionService(logger lib.Logger, repository repository.TransactionRepository) TransactionService {
	return TransactionService{
		logger:     logger,
		repository: repository,
	}
}

// CreateTransaction call to create the Transaction
func (s TransactionService) CreateTransaction(Transaction requests.TransactionRequest) (responses.TransactionCreateResponse, error) {
	refCode := lib.GenerateReferenceNumber()
	transaction := requests.TransactionRequest{
		Appname:       Transaction.Appname,
		Prefix:        Transaction.Prefix,
		Data:          Transaction.Data,
		ExpiredDate:   Transaction.ExpiredDate,
		ReferenceCode: refCode,
		Status:        "Open",
	}

	referenceCode, err := s.repository.Save(transaction)
	response := responses.TransactionCreateResponse{
		ReferenceCode: referenceCode,
	}

	return response, err
}

// UpdateTransaction updates the Transaction
func (s TransactionService) UpdateTransaction(params requests.UpdateRequest) (response responses.TransactionCreateResponse, err error) {
	Transaction := s.repository.MatchSearch(params.ReferenceCode)

	transaction := models.Transaction{
		Id:            Transaction.Id,
		Appname:       Transaction.Appname,
		Prefix:        Transaction.Prefix,
		Data:          Transaction.Data,
		ExpiredDate:   Transaction.ExpiredDate,
		ReferenceCode: Transaction.ReferenceCode,
		Status:        "Executed",
	}

	referenceCode, err := s.repository.Update(transaction)

	response = responses.TransactionCreateResponse{
		ReferenceCode: referenceCode,
	}

	return response, err
}

// UpdateTransaction updates the Transaction
func (s TransactionService) InquiryTransaction(request requests.InquiryRequest) (response responses.Data, status bool, err error) {
	Transaction, status := s.repository.InquiryTransaction(request)
	fmt.Println("from service:= ", Transaction)

	return Transaction.Data, status, err
}
