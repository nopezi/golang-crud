package jobs

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"infolelang/lib"
	"infolelang/models"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"gorm.io/gorm"
)

type JobRemoveRepository struct {
	DB      *gorm.DB
	timeout time.Duration
}

func NewJobRepository(db *gorm.DB) JobRemoveRepository {
	return JobRemoveRepository{
		DB:      db,
		timeout: time.Second * 10,
	}
}

// crontjob remove index reference_sequence
func (job JobRemoveRepository) JobsRemove() error {
	err := job.DB.Raw("truncate table reference_code_counters").Error
	if err != nil {
		lib.LogError(err.Error())
	}
	// DeleteElasticIndex()
	return nil
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
		filename, function, line := lib.WhereAmI()
		lib.CreateLogErrorToDB(client, filename, function, line, "Elasticsearch connection error", fmt.Sprintf("%v", err))
	}

	query2 := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}
	var buff bytes.Buffer
	if err := json.NewEncoder(&buff).Encode(query2); err != nil {
		log.Printf("Error encoding query: %s", err)
		filename, function, line := lib.WhereAmI()
		lib.CreateLogErrorToDB(client, filename, function, line, "Error encoding query", fmt.Sprintf("%v", err))
	}

	// Instantiate a map interface object for storing returned documents
	var mapResp map[string]interface{}
	// var buf bytes.Buffer

	// Attempt to encode the JSON query and look for errors
	if err := json.NewEncoder(&buff).Encode(query2); err != nil {
		log.Fatalf("json.NewEncoder() ERROR:", err)
		filename, function, line := lib.WhereAmI()
		lib.CreateLogErrorToDB(client, filename, function, line, "json.NewEncoder() ERROR", fmt.Sprintf("%v", err))
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
			filename, function, line := lib.WhereAmI()
			lib.CreateLogErrorToDB(client, filename, function, line, "Error getting response", fmt.Sprintf("%v", err))
		}

		// Check for any errors returned by API call to Elasticsearch
		if err != nil {
			log.Fatalf("Elasticsearch Search() API ERROR:", err)
			filename, function, line := lib.WhereAmI()
			lib.CreateLogErrorToDB(client, filename, function, line, "Elasticsearch Search() API ERROR", fmt.Sprintf("%v", err))
			// If no errors are returned, parse esapi.Response object
		} else {
			fmt.Println("res TYPE:", reflect.TypeOf(res))

			// Close the result body when the function call is complete
			defer res.Body.Close()

			// Decode the JSON response and using a pointer
			if err := json.NewDecoder(res.Body).Decode(&mapResp); err == nil {
				fmt.Println(`&mapResp:`, mapResp, "\n")
				fmt.Println(`mapResp["hits"]:`, mapResp["hits"])
				filename, function, line := lib.WhereAmI()
				lib.CreateLogErrorToDB(client, filename, function, line, "&mapResp:", fmt.Sprintf("%v", mapResp))
				lib.CreateLogErrorToDB(client, filename, function, line, "mapResp[]:", fmt.Sprintf("%v", mapResp["hits"]))
			}
		}
	}
}

// func SearchElastic(param string) {
// 	// client elasticsearch.Client
// 	var transaction models.Transaction
// 	// Allow for custom formatting of log output
// 	log.SetFlags(0)

// 	// Create a context object for the API calls
// 	ctx := context.Background()

// 	// Instantiate an Elasticsearch configuration
// 	cfg := elasticsearch.Config{
// 		Addresses: []string{
// 			os.Getenv("DBEHost"),
// 		},
// 		Username: os.Getenv("DBEUsername"),
// 		Password: os.Getenv("DBEPassword"),
// 		Transport: &http.Transport{
// 			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
// 		},
// 	}

// 	// Instantiate a new Elasticsearch client object instance
// 	client, err := elasticsearch.NewClient(cfg)

// 	// Check for connection errors to the Elasticsearch cluster
// 	if err != nil {
// 		fmt.Println("Elasticsearch connection error:", err)
// 		filename, function, line := lib.WhereAmI()
// 		lib.CreateLogErrorToDB(client, filename, function, line, "Elasticsearch connection error", fmt.Sprintf("%v", err))
// 	}

// 	// Create a new query string for the Elasticsearch method call
// 	var query = `

// 	"match": {
// 	"referenceCode":`
// 	query += `"` + param +
// 		`"}`

// 	// Pass the query string to the function and have it return a Reader object
// 	read := constructQuery(query, 2)

// 	// Example of an invalid JSON string
// 	//read = constructQuery("{bad json", 2)

// 	fmt.Println("read:", read)

// 	// Instantiate a map interface object for storing returned documents
// 	var mapResp map[string]interface{}
// 	var buf bytes.Buffer

// 	// Attempt to encode the JSON query and look for errors
// 	if err := json.NewEncoder(&buf).Encode(read); err != nil {
// 		log.Printf("json.NewEncoder() ERROR:", err)

// 		filename, function, line := lib.WhereAmI()
// 		lib.CreateLogErrorToDB(client, filename, function, line, "Attempt to encode the JSON query and look for errors", fmt.Sprintf("%v", err))

// 	} else {
// 		fmt.Println("json.NewEncoder encoded query:", read, "\n")

// 		// Pass the JSON query to the Golang client's Search() method
// 		res, err := client.Search(
// 			client.Search.WithContext(ctx),
// 			client.Search.WithIndex(transaction.IndexReferenceSequence()),
// 			client.Search.WithBody(read),
// 			client.Search.WithTrackTotalHits(true),
// 		)

// 		// Check for any errors returned by API call to Elasticsearch
// 		if err != nil {
// 			log.Fatalf("Elasticsearch Search() API ERROR:", err)
// 			filename, function, line := lib.WhereAmI()
// 			lib.CreateLogErrorToDB(client, filename, function, line, "Elasticsearch Search() API ERROR", fmt.Sprintf("%v", err))

// 			// If no errors are returned, parse esapi.Response object
// 		} else {
// 			fmt.Println("res TYPE:", reflect.TypeOf(res))

// 			// Close the result body when the function call is complete
// 			defer res.Body.Close()

// 			// Decode the JSON response and using a pointer
// 			if err := json.NewDecoder(res.Body).Decode(&mapResp); err == nil {
// 				fmt.Println(`&mapResp:`, &mapResp, "\n")
// 				fmt.Println(`mapResp["hits"]:`, mapResp["hits"])
// 			}
// 		}
// 	}
// }
