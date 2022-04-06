package jobs

import (
	"bytes"
	"context"
	"crypto/tls"
	"eform-gateway/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type JobRemoveRepository struct {
	Elastic *elasticsearch.Client
	timeout time.Duration
}

func NewJobRepository(elastic *elasticsearch.Client) JobRemoveRepository {
	return JobRemoveRepository{
		Elastic: elastic,
		timeout: time.Second * 10,
	}
}

// crontjob remove index reference_sequence
func (job JobRemoveRepository) JobsRemove() error {
	// var job JobRemoveRepository
	url := os.Getenv("DBEHost")
	fmt.Println("jobs Remove", url)

	// select all and remove one by one

	// data := job.MatchSearch("DPS0100000002")
	// fmt.Println(data)
	// SearchElastic("DPS0100000002")
	DeleteElasticIndex()
	// *job.Elastic
	return nil
}

func SearchElastic(param string) {
	// client elasticsearch.Client
	var transaction models.Transaction
	// Allow for custom formatting of log output
	log.SetFlags(0)

	// Create a context object for the API calls
	ctx := context.Background()

	// Instantiate an Elasticsearch configuration
	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("DBEHost"),
		},
		Username: os.Getenv("DBEUsername"),
		Password: os.Getenv("DBEPassword"),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		},
	}

	// Instantiate a new Elasticsearch client object instance
	client, err := elasticsearch.NewClient(cfg)

	// Check for connection errors to the Elasticsearch cluster
	if err != nil {
		fmt.Println("Elasticsearch connection error:", err)
	}

	// Create a new query string for the Elasticsearch method call
	var query = `
	
	"match": {
	"referenceCode":`
	query += `"` + param +
		`"}`

	// Pass the query string to the function and have it return a Reader object
	read := constructQuery(query, 2)

	// Example of an invalid JSON string
	//read = constructQuery("{bad json", 2)

	fmt.Println("read:", read)

	// Instantiate a map interface object for storing returned documents
	var mapResp map[string]interface{}
	var buf bytes.Buffer

	// Attempt to encode the JSON query and look for errors
	if err := json.NewEncoder(&buf).Encode(read); err != nil {
		log.Fatalf("json.NewEncoder() ERROR:", err)

		// Query is a valid JSON object
	} else {
		fmt.Println("json.NewEncoder encoded query:", read, "\n")

		// Pass the JSON query to the Golang client's Search() method
		res, err := client.Search(
			client.Search.WithContext(ctx),
			client.Search.WithIndex(transaction.IndexReferenceSequence()),
			client.Search.WithBody(read),
			client.Search.WithTrackTotalHits(true),
		)

		// Check for any errors returned by API call to Elasticsearch
		if err != nil {
			log.Fatalf("Elasticsearch Search() API ERROR:", err)

			// If no errors are returned, parse esapi.Response object
		} else {
			fmt.Println("res TYPE:", reflect.TypeOf(res))

			// Close the result body when the function call is complete
			defer res.Body.Close()

			// Decode the JSON response and using a pointer
			if err := json.NewDecoder(res.Body).Decode(&mapResp); err == nil {
				fmt.Println(`&mapResp:`, &mapResp, "\n")
				fmt.Println(`mapResp["hits"]:`, mapResp["hits"])
			}
		}
	}
}

func DeleteElasticIndex() {
	// client elasticsearch.Client
	var transaction models.Transaction
	// Allow for custom formatting of log output
	log.SetFlags(0)

	// Instantiate an Elasticsearch configuration
	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("DBEHost"),
		},
		Username: os.Getenv("DBEUsername"),
		Password: os.Getenv("DBEPassword"),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		},
	}

	// Instantiate a new Elasticsearch client object instance
	client, err := elasticsearch.NewClient(cfg)

	// Check for connection errors to the Elasticsearch cluster
	if err != nil {
		fmt.Println("Elasticsearch connection error:", err)
	}

	query2 := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}
	var buff bytes.Buffer
	if err := json.NewEncoder(&buff).Encode(query2); err != nil {
		log.Printf("Error encoding query: %s", err)
	}

	// Instantiate a map interface object for storing returned documents
	var mapResp map[string]interface{}
	// var buf bytes.Buffer

	// Attempt to encode the JSON query and look for errors
	if err := json.NewEncoder(&buff).Encode(query2); err != nil {
		log.Fatalf("json.NewEncoder() ERROR:", err)

		// Query is a valid JSON object
	} else {
		fmt.Println("json.NewEncoder encoded query:", query2, "\n")

		// Pass the JSON query to the Golang client's Search() method
		req := esapi.DeleteByQueryRequest{
			Index: []string{transaction.IndexReferenceSequence()},
			Body:  &buff,
		}

		res, err := req.Do(context.Background(), client)
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}

		// Check for any errors returned by API call to Elasticsearch
		if err != nil {
			log.Fatalf("Elasticsearch Search() API ERROR:", err)

			// If no errors are returned, parse esapi.Response object
		} else {
			fmt.Println("res TYPE:", reflect.TypeOf(res))

			// Close the result body when the function call is complete
			defer res.Body.Close()

			// Decode the JSON response and using a pointer
			if err := json.NewDecoder(res.Body).Decode(&mapResp); err == nil {
				fmt.Println(`&mapResp:`, mapResp, "\n")
				fmt.Println(`mapResp["hits"]:`, mapResp["hits"])
			}
		}
	}
}

func constructQuery(q string, size int) *strings.Reader {

	// Build a query string from string passed to function
	var query = `{"query": {`

	// Concatenate query string with string passed to method call
	query = query + q

	// Use the strconv.Itoa() method to convert int to string
	query = query + `}, "size": ` + strconv.Itoa(size) + `}`
	fmt.Println("\nquery:", query)

	// Check for JSON errors
	isValid := json.Valid([]byte(query)) // returns bool

	// Default query is "{}" if JSON is invalid
	if isValid == false {
		fmt.Println("constructQuery() ERROR: query string not valid:", query)
		fmt.Println("Using default match_all query")
		query = "{}"
	} else {
		fmt.Println("constructQuery() valid JSON:", isValid)
	}

	// Build a new string from JSON query
	var b strings.Builder
	b.WriteString(query)

	// Instantiate a *strings.Reader object from string
	read := strings.NewReader(b.String())

	// Return a *strings.Reader object
	return read
}

func (job JobRemoveRepository) MatchSearch(param string) (transaction models.Transaction) {
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

	// res, err := job.elastic.Client.Search(
	// 	job.elastic.Client.Search.WithContext(context.Background()),
	// 	job.elastic.Client.Search.WithIndex(transaction.IndexTransactionOpen()),
	// 	job.elastic.Client.Search.WithBody(&buf),
	// 	job.elastic.Client.Search.WithTrackTotalHits(true),
	// 	job.elastic.Client.Search.WithPretty(),
	// )

	// if err != nil {
	// 	log.Printf("Error getting response %s", err)
	// }

	// defer res.Body.Close()

	// if res.IsError() {
	// 	var e map[string]interface{}
	// 	if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
	// 		log.Printf("Error parsing the response body: %s", err)
	// 	} else {
	// 		log.Printf("[%s] %s: %s",
	// 			res.Status(),
	// 			e["error"].(map[string]interface{})["type"],
	// 			e["error"].(map[string]interface{})["reason"],
	// 		)
	// 	}
	// }
	// var dataTrx map[string]interface{}
	// if err := json.NewDecoder(res.Body).Decode(&dataTrx); err != nil {
	// 	log.Printf("Error parsing the response body: %s", err)
	// }

	// // Print the response status, number of results, and request duration.
	// log.Printf(
	// 	"[%s] %d hits; took: %dms",
	// 	res.Status(),
	// 	int(dataTrx["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
	// 	int(dataTrx["took"].(float64)),
	// )

	// // Print the ID and document source for each hit.
	// for _, hit := range dataTrx["hits"].(map[string]interface{})["hits"].([]interface{}) {
	// 	log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	// 	log.Println(strings.Repeat("=>", 37))
	// 	source := hit.(map[string]interface{})["_source"]

	// 	id := hit.(map[string]interface{})["_id"]
	// 	appname := source.(map[string]interface{})["appname"]
	// 	data := source.(map[string]interface{})["data"]
	// 	prefix := source.(map[string]interface{})["prefix"]
	// 	expiredDate := source.(map[string]interface{})["expiredDate"]
	// 	referenceCode := source.(map[string]interface{})["referenceCode"]
	// 	status := source.(map[string]interface{})["status"]

	// 	transaction = models.Transaction{
	// 		Id:            id.(string),
	// 		Appname:       appname.(string),
	// 		Data:          data,
	// 		Prefix:        prefix.(string),
	// 		ExpiredDate:   expiredDate.(string),
	// 		ReferenceCode: referenceCode.(string),
	// 		Status:        status.(string),
	// 	}

	// 	// fmt.Println(transaction)
	// 	log.Println(strings.Repeat("=>", 37))
	// }

	return transaction
}
