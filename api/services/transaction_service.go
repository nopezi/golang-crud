package services

import (
	"eform-gateway/api/repository"
	"eform-gateway/lib"
	"eform-gateway/models"
	"eform-gateway/requests"
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

// // WithTrx delegates transaction to repository database
// func (s TransactionService) WithTrx(trxHandle *gorm.DB) TransactionService {
// 	s.repository = s.repository.WithTrx(trxHandle)
// 	return s
// }

// // GetOneTransaction gets one Transaction
// func (s TransactionService) GetOneTransaction(id uint) (models.Transaction, error) {
// 	Transaction, err := s.repository.GetOne(id)
// 	return Transaction, err
// }

// // GetOneTransaction gets one Transaction
// func (s TransactionService) GetOneTransactionEmail(email *string) (models.Transaction, error) {
// 	Transaction, err := s.repository.GetTransactionByEmail(email)
// 	return Transaction, err
// }

// // GetAllTransaction get all the Transaction
// func (s TransactionService) GetAllTransaction() ([]models.Transaction, error) {
// 	Transactions, err := s.repository.GetAll()
// 	return Transactions, err
// }

// CreateTransaction call to create the Transaction
func (s TransactionService) CreateTransaction(Transaction requests.TransactionRequest) (string, error) {
	refCode := lib.GenerateReferenceNumber()
	transaction := requests.TransactionRequest{
		Appname:       Transaction.Appname,
		Object:        Transaction.Object,
		Prefix:        Transaction.Prefix,
		ExpiredDate:   Transaction.ExpiredDate,
		ReferenceCode: refCode,
		Status:        Transaction.Status,
	}

	referenceCode, err := s.repository.Save(transaction)

	return referenceCode, err
}

// UpdateTransaction updates the Transaction
func (s TransactionService) UpdateTransaction(params requests.UpdateRequest) (response bool, err error) {
	Transaction := s.repository.MatchSearch(params.ReferenceCode)
	fmt.Println("from service:= ", Transaction)
	transaction := models.Transaction{
		Id:            Transaction.Id,
		Appname:       Transaction.Appname,
		Object:        Transaction.Object,
		Prefix:        Transaction.Prefix,
		ExpiredDate:   Transaction.ExpiredDate,
		ReferenceCode: Transaction.ReferenceCode,
		Status:        "Executed",
	}

	response, err = s.repository.Update(transaction)

	return response, err
}

// // DeleteTransaction deletes the Transaction
// func (s TransactionService) DeleteTransaction(id uint) error {
// 	return s.repository.Delete(id)
// }
