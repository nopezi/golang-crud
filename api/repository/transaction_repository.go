package repository

import (
	"bytes"
	"context"
	"eform-gateway/lib"
	"eform-gateway/models"
	"eform-gateway/requests"
	"eform-gateway/responses"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// TransactionRepository database structure
type TransactionRepository struct {
	Elastic lib.Elasticsearch
	logger  lib.Logger
	timeout time.Duration
}

// NewTransactionRepository creates a new Transaction repository
func NewTransactionRepository(elastic lib.Elasticsearch, logger lib.Logger) TransactionRepository {
	return TransactionRepository{
		Elastic: elastic,
		logger:  logger,
		timeout: time.Second * 10,
	}
}

// Create Transaction
func (r TransactionRepository) CreateTransaction(Transaction requests.TransactionRequest) (referenceCode string, err error) {
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

	res, err := req.Do(ctx, r.Elastic.Client)
	if err != nil {
		return "", fmt.Errorf("insert: request: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return "", fmt.Errorf("insert: response: %s", res.String())
	}

	return Transaction.ReferenceCode, err
}

// UpdateToExecuted
func (r TransactionRepository) UpdateTransaction(Transaction models.Transaction) (string, error) {

	transaction := requests.TransactionRequest{
		Appname:       Transaction.Appname,
		Prefix:        Transaction.Prefix,
		Data:          Transaction.Data,
		ExpiredDate:   Transaction.ExpiredDate,
		ReferenceCode: Transaction.ReferenceCode,
		Status:        Transaction.Status,
		Created:       Transaction.Created,
		LastUpdate:    Transaction.LastUpdate,
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

	res, err := req.Do(ctx, r.Elastic.Client)
	if err != nil {
		return "", fmt.Errorf("insert: request: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return "", fmt.Errorf("insert: response: %s", res.String())
	}

	statusDelete, err := r.DeleteIndex(Transaction.IndexTransactionOpen(), Transaction.Id)

	if err != nil {
		return "", fmt.Errorf("insert: request: %w", err)
	}

	if !statusDelete {
		return "", err
	}

	return transaction.ReferenceCode, err
}

func (r TransactionRepository) DeleteIndex(index string, id string) (status bool, err error) {
	// Delete by Id
	reqDelete := esapi.DeleteRequest{
		Index:      index,
		DocumentID: id,
		// Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, bdy))),
	}
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	resDelete, err := reqDelete.Do(ctx, r.Elastic.Client)
	if err != nil {
		return false, fmt.Errorf("insert: request: %w", err)
	}
	defer resDelete.Body.Close()

	if resDelete.IsError() {
		return false, fmt.Errorf("insert: response: %s", resDelete.String())
	}
	return true, err
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
		log.Printf("Error encoding query: %s", err)
	}

	res, err := r.Elastic.Client.Search(
		r.Elastic.Client.Search.WithContext(context.Background()),
		r.Elastic.Client.Search.WithIndex(transaction.IndexTransactionOpen()),
		r.Elastic.Client.Search.WithBody(&buf),
		r.Elastic.Client.Search.WithTrackTotalHits(true),
		r.Elastic.Client.Search.WithPretty(),
	)

	if err != nil {
		log.Printf("Error getting response %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			log.Printf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	var dataTrx map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&dataTrx); err != nil {
		log.Printf("Error parsing the response body: %s", err)
	}

	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(dataTrx["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(dataTrx["took"].(float64)),
	)

	// Print the ID and document source for each hit.
	// fmt.Println("hits =>", dataTrx["hits"].(map[string]interface{})["hits"])
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
		created := source.(map[string]interface{})["created"]
		lastUpdate := source.(map[string]interface{})["lastUpdate"]

		transaction = models.Transaction{
			Id:            id.(string),
			Appname:       appname.(string),
			Data:          data,
			Prefix:        prefix.(string),
			ExpiredDate:   expiredDate.(string),
			ReferenceCode: referenceCode.(string),
			Status:        status.(string),
			Created:       created.(string),
			LastUpdate:    lastUpdate.(string),
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
			Created:       "",
			LastUpdate:    "",
		}
		return transaction, false
	}

	return transaction, true
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
		log.Printf("Error encoding query: %s", err)
	}

	res, err := r.Elastic.Client.Search(
		r.Elastic.Client.Search.WithContext(context.Background()),
		r.Elastic.Client.Search.WithIndex(transaction.IndexTransactionOpen()),
		r.Elastic.Client.Search.WithBody(&buf),
		r.Elastic.Client.Search.WithTrackTotalHits(true),
		r.Elastic.Client.Search.WithPretty(),
	)

	if err != nil {
		log.Printf("Error getting response %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			log.Printf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	var dataTrx map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&dataTrx); err != nil {
		log.Printf("Error parsing the response body: %s", err)
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
		created := source.(map[string]interface{})["created"]
		lastUpdate := source.(map[string]interface{})["lastUpdate"]

		transaction = models.Transaction{
			Id:            id.(string),
			Appname:       appname.(string),
			Data:          data,
			Prefix:        prefix.(string),
			ExpiredDate:   expiredDate.(string),
			ReferenceCode: referenceCode.(string),
			Status:        status.(string),
			Created:       created.(string),
			LastUpdate:    lastUpdate.(string),
		}

		// fmt.Println(transaction)
		log.Println(strings.Repeat("=>", 37))
	}

	return transaction
}

// CreateReferenceSequence
func (r TransactionRepository) CreateReferenceSequence(referenceSequence requests.ReferenceSequenceRequest) (responseSequence responses.ReferenceSequenceResponse, err error) {
	model := models.Transaction{}
	// sequences := responses.ReferenceSequenceResponse{}
	bdy, err := json.Marshal(referenceSequence)
	if err != nil {
		return responseSequence, fmt.Errorf("insert: marshall: %w", err)
	}

	req := esapi.CreateRequest{
		Index:      model.IndexReferenceSequence(),
		DocumentID: lib.UUID(false),
		Body:       bytes.NewReader(bdy),
	}

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	res, err := req.Do(ctx, r.Elastic.Client)
	if err != nil {
		return responseSequence, fmt.Errorf("insert: request: %w", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return responseSequence, fmt.Errorf("insert: response: %s", res.String())
	}

	return responseSequence, err
}

func (r TransactionRepository) FindPrefixReferenceSequence(param string) (response responses.ReferenceSequenceResponse, status bool) {
	var transaction models.Transaction
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"prefix": param,
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s", err)
	}

	res, err := r.Elastic.Client.Search(
		r.Elastic.Client.Search.WithContext(context.Background()),
		r.Elastic.Client.Search.WithIndex(transaction.IndexReferenceSequence()),
		r.Elastic.Client.Search.WithBody(&buf),
		r.Elastic.Client.Search.WithTrackTotalHits(true),
		r.Elastic.Client.Search.WithPretty(),
	)

	if err != nil {
		log.Printf("Error getting response %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			log.Printf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
			return response, false
		}
	}
	var dataTrx map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&dataTrx); err != nil {
		log.Printf("Error parsing the response body: %s", err)
	}

	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(dataTrx["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(dataTrx["took"].(float64)),
	)

	// Print the ID and document source for each hit.
	hits := dataTrx["hits"].(map[string]interface{})["hits"]
	for _, hit := range dataTrx["hits"].(map[string]interface{})["hits"].([]interface{}) {

		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])

		source := hit.(map[string]interface{})["_source"]

		// log.Println(strings.Repeat("=>", 37))

		id := hit.(map[string]interface{})["_id"]
		// log.Println(id)
		pref := source.(map[string]interface{})["prefix"]
		// log.Println(pref)

		sequence := source.(map[string]interface{})["sequence"]
		sequenceString := fmt.Sprint(sequence)
		sequence64, _ := strconv.ParseInt(sequenceString, 10, 64)
		// log.Println(sequence)

		response = responses.ReferenceSequenceResponse{
			Id:       id.(string),
			Prefix:   pref.(string),
			Sequence: sequence64,
		}
		// log.Println(response)

		// log.Println(strings.Repeat("=>", 37))
		// fmt.Println("from FindPrefixReferenceSequence", response)
		// log.Println(strings.Repeat("=>", 37))
	}
	hitsLen := reflect.ValueOf(hits)
	if hitsLen.Len() == 0 {
		return response, false
	} else {
		return response, true
	}

}
