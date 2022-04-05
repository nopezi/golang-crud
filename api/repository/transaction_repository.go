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
	"reflect"
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

// Save Transaction
func (r TransactionRepository) Save(Transaction requests.TransactionRequest) (referenceCode string, err error) {
	model := models.Transaction{}
	bdy, err := json.Marshal(Transaction)
	if err != nil {
		return "", fmt.Errorf("insert: marshall: %w", err)
	}

	// res, err := p.elastic.client.Create()
	req := esapi.CreateRequest{
		Index:      model.IndexTransactionOpen(),
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
func (r TransactionRepository) Update(Transaction models.Transaction) (string, error) {

	transaction := requests.TransactionRequest{
		Appname:       Transaction.Appname,
		Prefix:        Transaction.Prefix,
		Data:          Transaction.Data,
		ExpiredDate:   Transaction.ExpiredDate,
		ReferenceCode: Transaction.ReferenceCode,
		Status:        Transaction.Status,
	}

	// Create data
	bdy, err := json.Marshal(transaction)
	if err != nil {
		return "", fmt.Errorf("insert: marshall: %w", err)
	}

	req := esapi.CreateRequest{
		Index:      Transaction.IndexTransactionExecuted(),
		DocumentID: lib.UUID(false),
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, bdy))),
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

	// Delete by Id
	reqDelete := esapi.DeleteRequest{
		Index:      Transaction.IndexTransactionOpen(),
		DocumentID: Transaction.Id,
		// Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, bdy))),
	}

	resDelete, err := reqDelete.Do(ctx, r.elastic.Client)
	if err != nil {
		return "", fmt.Errorf("insert: request: %w", err)
	}
	defer resDelete.Body.Close()

	if res.IsError() {
		return "", fmt.Errorf("insert: response: %s", res.String())
	}

	return transaction.ReferenceCode, err
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
		r.elastic.Client.Search.WithIndex(transaction.IndexTransactionOpen()),
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
		source := hit.(map[string]interface{})["_source"]

		id := hit.(map[string]interface{})["_id"]
		appname := source.(map[string]interface{})["appname"]
		data := source.(map[string]interface{})["data"]
		prefix := source.(map[string]interface{})["prefix"]
		expiredDate := source.(map[string]interface{})["expiredDate"]
		referenceCode := source.(map[string]interface{})["referenceCode"]
		status := source.(map[string]interface{})["status"]

		transaction = models.Transaction{
			Id:            id.(string),
			Appname:       appname.(string),
			Data:          data,
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

func (r TransactionRepository) InquiryTransaction(request requests.InquiryRequest) (transaction models.Transaction, notFound bool) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"referenceCode": request.ReferenceCode,
				// "status":        "Open",
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := r.elastic.Client.Search(
		r.elastic.Client.Search.WithContext(context.Background()),
		r.elastic.Client.Search.WithIndex(transaction.IndexTransactionOpen()),
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
	fmt.Println("hits =>", dataTrx["hits"].(map[string]interface{})["hits"])
	hits := dataTrx["hits"].(map[string]interface{})["hits"]
	for _, hit := range hits.([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
		log.Println(strings.Repeat("=>", 37))
		source := hit.(map[string]interface{})["_source"]

		id := hit.(map[string]interface{})["_id"]
		appname := source.(map[string]interface{})["appname"]
		data := source.(map[string]interface{})["data"]
		prefix := source.(map[string]interface{})["prefix"]
		expiredDate := source.(map[string]interface{})["expiredDate"]
		referenceCode := source.(map[string]interface{})["referenceCode"]
		status := source.(map[string]interface{})["status"]

		transaction = models.Transaction{
			Id:            id.(string),
			Appname:       appname.(string),
			Data:          data,
			Prefix:        prefix.(string),
			ExpiredDate:   expiredDate.(string),
			ReferenceCode: referenceCode.(string),
			Status:        status.(string),
		}

		log.Println(strings.Repeat("=>", 37))
	}

	hitsLen := reflect.ValueOf(hits)
	if hitsLen.Len() == 0 {
		transaction = models.Transaction{
			Id:            "",
			Appname:       "",
			Data:          models.Data{},
			Prefix:        "",
			ExpiredDate:   "",
			ReferenceCode: "",
			Status:        "",
		}
		return transaction, false
	}

	return transaction, true
}
