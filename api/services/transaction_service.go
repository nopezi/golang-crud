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

	requestSequence := requests.ReferenceSequenceRequest{}
	model := models.Transaction{}

	referenceSequence, status := s.repository.FindPrefixReferenceSequence(Transaction.Prefix)
	// fmt.Println("From CreateReferenceSequence before create", referenceSequence)
	if status {
		requestSequence = requests.ReferenceSequenceRequest{
			Prefix:   Transaction.Prefix,
			Sequence: referenceSequence.Sequence + 1,
		}

		_, err := s.repository.DeleteIndex(model.IndexReferenceSequence(), referenceSequence.Id)
		// fmt.Println("Delete Index=>", data, err)
		if err != nil {
			fmt.Println(err.Error())
		}
		if err != nil {
			filename, function, line := lib.WhereAmI()
			lib.CreateLogErrorToDB(s.repository.Elastic.Client, filename, function, line, "Delete Index Gagal", fmt.Sprintf("%v", err))
			return response, err
		}
	} else {
		requestSequence = requests.ReferenceSequenceRequest{
			Prefix:   Transaction.Prefix,
			Sequence: 1,
		}
		fmt.Println("false, =>>", requestSequence)
	}

	_, err = s.repository.CreateReferenceSequence(requestSequence)
	if err != nil {
		response = responses.TransactionCreateResponse{
			ReferenceCode: "",
		}
		if err != nil {
			filename, function, line := lib.WhereAmI()
			lib.CreateLogErrorToDB(s.repository.Elastic.Client, filename, function, line, "CreateReferenceSequence Gagal", fmt.Sprintf("%v", err))
			return response, err
		}
		return response, err
	} else {
		sequences, _ := s.repository.FindPrefixReferenceSequence(requestSequence.Prefix)
		// fmt.Println("From CreateReferenceSequence after create", sequences)

		sequencePadLeft := lib.StrPadLeft(fmt.Sprint(sequences.Sequence+1), 8, "0")
		referenceCode := Transaction.Prefix + sequencePadLeft
		// fmt.Println("referenceCode", referenceCode)
		// Create Transaction
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
			filename, function, line := lib.WhereAmI()
			lib.CreateLogErrorToDB(s.repository.Elastic.Client, filename, function, line, "CreateTransaction Gagal", fmt.Sprintf("%v", err))
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
		Created:       Transaction.Created,
		LastUpdate:    lib.GetTimeNow("timestime"),
	}

	referenceCode, err := s.repository.UpdateTransaction(transaction)
	if err != nil {
		filename, function, line := lib.WhereAmI()
		lib.CreateLogErrorToDB(s.repository.Elastic.Client, filename, function, line, "UpdateTransaction Gagal", fmt.Sprintf("%v", err))
		return response, err
	}

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
