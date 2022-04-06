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
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

// - [ ] Cronjob update expired date by timestime, Not including this service api, registered on crontab linux
// - search index where documen if expired_date = now , create to transactionExpireds and delete index from transactions

func JobsUpdate() {
	fmt.Println("jobs Update")
	// match expired_date = now
	err := MatchSearch()
	if err != nil {
		fmt.Println(err.Error())
	}
	// tampung ke struct []
	// create ke transaction.IndexTransactionExpired()
	// loop tampungan
	// delete by match expired_date = now
}

func MatchSearch() (err error) {
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
	dateNow := GetTimeNow()
	var query = `
	"match": {
	"expiredDate":`
	query += `"` + dateNow +
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
			client.Search.WithIndex(transaction.IndexTransactionOpen()),
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
			transactions := []models.Transaction{}
			hits := mapResp["hits"].(map[string]interface{})["hits"]
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

				transactions = append(transactions, models.Transaction{
					Id:            id.(string),
					Appname:       appname.(string),
					Data:          data,
					Prefix:        prefix.(string),
					ExpiredDate:   expiredDate.(string),
					ReferenceCode: referenceCode.(string),
					Status:        status.(string),
				})
				log.Println(strings.Repeat("=>", 37))
			}
			fmt.Println("transactions", transactions)
		}
	}
	return err
}
