package repository

import (
	"bytes"
	"context"
	"eform-gateway/lib"
	"eform-gateway/models"
	"encoding/json"
	"fmt"
	"time"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// TransactionRepository database structure
type TransactionRepository struct {
	elastic lib.Elasticsearch
	logger  lib.Logger
	timeout time.Duration
}

// NewTransactionRepository creates a new Transaction repository
func NewTransactionRepository(elastic lib.Elasticsearch, logger lib.Logger) TransactionRepository {
	return TransactionRepository{
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 10,
	}
}

// // GetAll gets all Transactions
// func (r TransactionRepository) GetAll() (Transactions []models.Transaction, err error) {
// 	return Transactions, r.db.DB.Find(&Transactions).Error
// }

// Save Transaction
func (r TransactionRepository) Save(Transaction models.Transaction) (referenceCode string, err error) {
	bdy, err := json.Marshal(Transaction)
	if err != nil {
		return "", fmt.Errorf("insert: marshall: %w", err)
	}

	// res, err := p.elastic.client.Create()
	req := esapi.CreateRequest{
		Index:      Transaction.IndexName(),
		DocumentID: lib.UUID(false),
		Body:       bytes.NewReader(bdy),
	}

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	res, err := req.Do(ctx, r.elastic.Client)
	if err != nil {
		return "", fmt.Errorf("insert: request: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return "", fmt.Errorf("insert: response: %s", res.String())
	}

	return Transaction.ReferenceCode, err
}

// // Update updates Transaction
// func (r TransactionRepository) Update(Transaction models.Transaction) (models.Transaction, error) {
// 	return Transaction, r.db.DB.Save(&Transaction).Error
// }

// // GetOne gets ont Transaction
// func (r TransactionRepository) GetOne(id uint) (Transaction models.Transaction, err error) {
// 	return Transaction, r.db.DB.Where("id = ?", id).First(&Transaction).Error
// }

// // GetOne gets Transaction by email
// func (r TransactionRepository) GetTransactionByEmail(email *string) (Transaction models.Transaction, err error) {
// 	return Transaction, r.db.DB.Where("email = ?", email).First(&Transaction).Error
// }

// // Delete deletes the row of data
// func (r TransactionRepository) Delete(id uint) error {
// 	return r.db.DB.Where("id = ?", id).Delete(&models.Transaction{}).Error
// }
