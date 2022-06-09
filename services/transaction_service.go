package services

import (
	"fmt"
	"infolelang/lib"
	"infolelang/models"
	"infolelang/repository"
	"infolelang/requests"
	"infolelang/responses"

	"gitlab.com/golang-package-library/logger"
)

type TransactionService interface {
	CreateTransaction(Transaction requests.TransactionRequest) (response responses.TransactionCreateResponse, err error)
	UpdateTransaction(params requests.UpdateRequest) (response responses.TransactionCreateResponse, err error)
	InquiryTransaction(request requests.InquiryRequest) (response responses.Data, status bool, err error)
}

// TransactionServiceContext service layer
type TransactionServiceContext struct {
	logger     logger.Logger
	repository repository.TransactionRepository
}

// NewTransactionServiceContext creates a new TransactionServiceContext
func NewTransactionService(logger logger.Logger, repository repository.TransactionRepository) TransactionService {
	return TransactionServiceContext{
		logger:     logger,
		repository: repository,
	}
}

// CreateTransaction call to create the Transaction
func (s TransactionServiceContext) CreateTransaction(Transaction requests.TransactionRequest) (response responses.TransactionCreateResponse, err error) {
	requestSequence := requests.ReferenceSequenceRequest{}
	referenceSequence, status := s.repository.GetPrefixReferenceSequence(Transaction.Prefix)

	fmt.Println("referenceSequence =>", referenceSequence)
	fmt.Println("status =>", status)

	var errCounter error
	var dataCounter responses.ReferenceSequenceResponse

	if status {
		requestSequence = requests.ReferenceSequenceRequest{
			Id:       referenceSequence.Id,
			Prefix:   referenceSequence.Prefix,
			Sequence: referenceSequence.Sequence + 1,
		}
		dataCounter, errCounter = s.repository.UpdateReferenceCounter(requestSequence)
	} else {
		requestSequence = requests.ReferenceSequenceRequest{
			Prefix:   Transaction.Prefix,
			Sequence: 1,
		}
		dataCounter, errCounter = s.repository.CreateReferenceCounter(requestSequence)
	}

	fmt.Println("errCounter", errCounter)
	fmt.Println("dataCounter", dataCounter)
	if errCounter != nil {
		response = responses.TransactionCreateResponse{
			ReferenceCode: "",
		}
		if err != nil {
			// filename, function, line := lib.WhereAmI()
			// lib.CreateLogErrorToDB(s.repository.Elastic.Client, filename, function, line, "CreateReferenceSequence Gagal", fmt.Sprintf("%v", err))
			return response, err
		}
		return response, err
	} else {
		// sequences, _ := s.repository.GetPrefixReferenceSequence(requestSequence.Prefix)

		sequencePadLeft := lib.StrPadLeft(fmt.Sprint(dataCounter.Sequence), 8, "0")
		referenceCode := dataCounter.Prefix + sequencePadLeft

		transaction := requests.TransactionRequest{
			Appname:       Transaction.Appname,
			Prefix:        Transaction.Prefix,
			Data:          Transaction.Data,
			ExpiredDate:   Transaction.ExpiredDate,
			ReferenceCode: referenceCode,
			Status:        "Open",
			Created:       lib.GetTimeNow("timestime"),
		}

		_, err = s.repository.CreateTransaction(transaction)
		if err != nil {
			// filename, function, line := lib.WhereAmI()
			// lib.CreateLogErrorToDB(s.repository.Elastic.Client, filename, function, line, "CreateTransaction Gagal", fmt.Sprintf("%v", err))
			return response, err
		}

		response = responses.TransactionCreateResponse{
			ReferenceCode: referenceCode,
		}
	}

	return response, err
}

// UpdateTransaction updates the Transaction
func (s TransactionServiceContext) UpdateTransaction(params requests.UpdateRequest) (response responses.TransactionCreateResponse, err error) {

	Transaction := s.repository.MatchSearch(params.ReferenceCode)

	transaction := models.Transaction{
		Id:            Transaction.Id,
		Appname:       Transaction.Appname,
		Prefix:        Transaction.Prefix,
		Data:          Transaction.Data,
		ExpiredDate:   Transaction.ExpiredDate,
		ReferenceCode: Transaction.ReferenceCode,
		Status:        "Executed",
		Created:       Transaction.Created,
		LastUpdate:    lib.GetTimeNow("timestime"),
	}

	referenceCode, err := s.repository.UpdateTransaction(transaction)
	if err != nil {
		// filename, function, line := lib.WhereAmI()
		// lib.CreateLogErrorToDB(s.repository.Elastic.Client, filename, function, line, "UpdateTransaction Gagal", fmt.Sprintf("%v", err))
		return response, err
	}

	response = responses.TransactionCreateResponse{
		ReferenceCode: referenceCode,
	}

	return response, err
}

// UpdateTransaction updates the Transaction
func (s TransactionServiceContext) InquiryTransaction(request requests.InquiryRequest) (response responses.Data, status bool, err error) {
	Transaction, status := s.repository.InquiryTransaction(request)
	// fmt.Println("from service:= ", Transaction)

	return Transaction.Data, status, err
}
