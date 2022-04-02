package services

import (
	"eform-gateway/api/repository"
	"eform-gateway/lib"
	"eform-gateway/models"
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
func (s TransactionService) CreateTransaction(Transaction models.Transaction) error {
	err := s.repository.Save(Transaction)
	return err
}

// // UpdateTransaction updates the Transaction
// func (s TransactionService) UpdateTransaction(id uint, Transaction models.Transaction) error {

// 	TransactionDB, err := s.GetOneTransaction(id)
// 	if err != nil {
// 		return err
// 	}

// 	copier.Copy(&TransactionDB, &Transaction)

// 	TransactionDB.ID = id

// 	_, err = s.repository.Update(TransactionDB)
// 	return err
// }

// // DeleteTransaction deletes the Transaction
// func (s TransactionService) DeleteTransaction(id uint) error {
// 	return s.repository.Delete(id)
// }
