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
func (s TransactionService) CreateTransaction(Transaction requests.TransactionRequest) (response responses.TransactionCreateResponse, err error) {
	// refCode := lib.GenerateReferenceNumber()
	requestSequence := requests.ReferenceSequenceRequest{}
	model := models.Transaction{}
	// find prefix Transaction.Prefix in index reference_sequence
	// if exist increment sequence
	// if not exist create new
	// createReference
	// Find sequence with param prefix and  prefix + 00000 +sequence

	referenceSequence, status := s.repository.FindPrefixReferenceSequence(Transaction.Prefix)
	fmt.Println("From CreateReferenceSequence before create", referenceSequence)
	if status {
		requestSequence = requests.ReferenceSequenceRequest{
			Prefix:   Transaction.Prefix,
			Sequence: referenceSequence.Sequence + 1,
		}

		data, err := s.repository.DeleteIndex(model.IndexReferenceSequence(), referenceSequence.Id)
		fmt.Println("Delete Index=>", data, err)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		requestSequence = requests.ReferenceSequenceRequest{
			Prefix:   Transaction.Prefix,
			Sequence: 1,
		}
	}

	_, err = s.repository.CreateReferenceSequence(requestSequence)
	if err != nil {
		response = responses.TransactionCreateResponse{
			ReferenceCode: "",
		}
		return response, err
	} else {
		sequences, _ := s.repository.FindPrefixReferenceSequence(requestSequence.Prefix)
		fmt.Println("From CreateReferenceSequence after create", sequences)

		sequencePadLeft := lib.StrPadLeft(fmt.Sprint(sequences.Sequence+1), 8, "0")
		referenceCode := sequences.Prefix + sequencePadLeft

		// Create Transaction
		transaction := requests.TransactionRequest{
			Appname:       Transaction.Appname,
			Prefix:        Transaction.Prefix,
			Data:          Transaction.Data,
			ExpiredDate:   Transaction.ExpiredDate,
			ReferenceCode: referenceCode,
			Status:        "Open",
		}

		_, err = s.repository.CreateTransaction(transaction)
		if err != nil {
			return response, err
		}

		response = responses.TransactionCreateResponse{
			ReferenceCode: referenceCode,
		}
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

	referenceCode, err := s.repository.UpdateTransaction(transaction)

	response = responses.TransactionCreateResponse{
		ReferenceCode: referenceCode,
	}

	return response, err
}

// UpdateTransaction updates the Transaction
func (s TransactionService) InquiryTransaction(request requests.InquiryRequest) (response responses.Data, status bool, err error) {
	Transaction, status := s.repository.InquiryTransaction(request)
	// fmt.Println("from service:= ", Transaction)

	return Transaction.Data, status, err
}
