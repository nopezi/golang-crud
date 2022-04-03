package repository

import (
	"bytes"
	"context"
	"eform-gateway/lib"
	"eform-gateway/models"
	"eform-gateway/requests"
	"encoding/json"
	"fmt"
	"log"
	"strings"
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
func (r TransactionRepository) Save(Transaction requests.TransactionRequest) (referenceCode string, err error) {
	model := models.Transaction{}
	bdy, err := json.Marshal(Transaction)
	if err != nil {
		return "", fmt.Errorf("insert: marshall: %w", err)
	}

	// res, err := p.elastic.client.Create()
	req := esapi.CreateRequest{
		Index:      model.IndexName(),
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

// Update updates Transaction
// UpdateToExecuted
func (r TransactionRepository) Update(Transaction models.Transaction) (bool, error) {
	transaction := requests.TransactionRequest{
		Appname:       Transaction.Appname,
		Object:        Transaction.Object,
		Prefix:        Transaction.Prefix,
		ExpiredDate:   Transaction.ExpiredDate,
		ReferenceCode: Transaction.ReferenceCode,
		Status:        Transaction.Status,
	}

	bdy, err := json.Marshal(transaction)
	if err != nil {
		return false, fmt.Errorf("insert: marshall: %w", err)
	}

	req := esapi.UpdateRequest{
		Index:      Transaction.IndexName(),
		DocumentID: Transaction.Id,
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, bdy))),
	}

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	res, err := req.Do(ctx, r.elastic.Client)
	if err != nil {
		return false, fmt.Errorf("insert: request: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return false, fmt.Errorf("insert: response: %s", res.String())
	}

	return true, err
}

func (r TransactionRepository) MatchSearch(param string) (transaction models.Transaction) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"referenceCode": param,
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := r.elastic.Client.Search(
		r.elastic.Client.Search.WithContext(context.Background()),
		r.elastic.Client.Search.WithIndex(transaction.IndexName()),
		r.elastic.Client.Search.WithBody(&buf),
		r.elastic.Client.Search.WithTrackTotalHits(true),
		r.elastic.Client.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Error getting response %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	var dataTrx map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&dataTrx); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(dataTrx["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(dataTrx["took"].(float64)),
	)

	// Print the ID and document source for each hit.
	for _, hit := range dataTrx["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
		log.Println(strings.Repeat("=>", 37))
		data := hit.(map[string]interface{})["_source"]

		id := hit.(map[string]interface{})["_id"]
		appname := data.(map[string]interface{})["appname"]
		object := data.(map[string]interface{})["object"]
		prefix := data.(map[string]interface{})["prefix"]
		expiredDate := data.(map[string]interface{})["expiredDate"]
		referenceCode := data.(map[string]interface{})["referenceCode"]
		status := data.(map[string]interface{})["status"]

		transaction = models.Transaction{
			Id:            id.(string),
			Appname:       appname.(string),
			Object:        object.(string),
			Prefix:        prefix.(string),
			ExpiredDate:   expiredDate.(string),
			ReferenceCode: referenceCode.(string),
			Status:        status.(string),
		}

		// fmt.Println(transaction)
		log.Println(strings.Repeat("=>", 37))
	}

	return transaction
}

// // GetOne gets ont Transaction
// - [ ] Inquiry search by reference number
//       - inquery where reference_code and status = Open
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
