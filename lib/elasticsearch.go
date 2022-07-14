package lib

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	elastic "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	Env "gitlab.com/golang-package-library/env"
	logger "gitlab.com/golang-package-library/logger"
)

type RequestElastic struct {
	DocumentID string
	Index      string
	Body       interface{}
}
type ElasticsearchDefinition interface {
	Store() (response bool, err error)
	Update() (response bool, err error)
	Search() (response []interface{}, err error)
	Delete() (response bool, err error)
}
type Elasticsearch struct {
	Client    *elastic.Client
	zapLogger logger.Logger
	timeout   time.Duration
}

func NewElastic(env Env.Env, zapLogger logger.Logger) Elasticsearch {
	url := env.DBEHost
	username := env.DBEUsername
	password := env.DBEPassword

	cfg := elastic.Config{
		Addresses: []string{
			url,
		},
		Username: username,
		Password: password,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		},
	}

	client, _ := elastic.NewClient(cfg)
	_, err := client.Info()
	if err != nil {
		zapLogger.Zap.Info("Url: ", url)
		zapLogger.Zap.Panic(err)
		zapLogger.Zap.Info("Elasticsearch Connection Refused")
	}

	zapLogger.Zap.Info("Elasticsearch Connection Established")
	return Elasticsearch{
		Client:  client,
		timeout: time.Second * 10,
	}
}

func (e Elasticsearch) Store(request RequestElastic) (response bool, err error) {
	body, err := json.Marshal(request.Body)
	if err != nil {
		e.zapLogger.Zap.Error(err)
		return false, err
	}

	req := esapi.CreateRequest{
		Index:      request.Index,
		DocumentID: request.DocumentID,
		Body:       bytes.NewReader(body),
	}

	ctx, cancel := context.WithTimeout(context.Background(), e.timeout)
	defer cancel()

	res, err := req.Do(ctx, e.Client)
	fmt.Println(res)
	if err != nil {
		fmt.Println("error", reflect.TypeOf(err))
		e.zapLogger.Zap.Error(err)
		return false, err
	}
	errExeeded := fmt.Sprint(err)
	if errExeeded == "context.deadlineExceededError" {
		fmt.Println("error", err)
		e.zapLogger.Zap.Error(err)
		return false, err
	}

	defer res.Body.Close()
	fmt.Println(res)

	if res.IsError() {
		e.zapLogger.Zap.Error(res.String())
		return false, err
	}

	return true, err
}

func (e Elasticsearch) Update(request RequestElastic) (response bool, err error) {
	body, err := json.Marshal(request.Body)
	if err != nil {
		e.zapLogger.Zap.Error(err)
	}

	req := esapi.UpdateRequest{
		Index:      request.Index,
		DocumentID: request.DocumentID,
		Body:       bytes.NewReader(body),
	}

	ctx, cancel := context.WithTimeout(context.Background(), e.timeout)
	defer cancel()

	res, err := req.Do(ctx, e.Client)
	if err != nil {
		e.zapLogger.Zap.Error(err)
		return false, err
	}
	defer res.Body.Close()

	if res.IsError() {
		e.zapLogger.Zap.Error(res.String())
		return false, err
	}

	return true, err
}

func (e Elasticsearch) Search(request RequestElastic) (response interface{}, err error) {
	var buf bytes.Buffer
	// query := map[string]interface{}{
	// 	"query": map[string]interface{}{
	// 		"match_all": map[string]interface{}{
	// 			// "name": "SHM",
	// 			// "status":        "Open",
	// 		},
	// 	},
	// }
	// query, err := json.Marshal(data)
	// if err != nil {
	// 	e.zapLogger.Zap.Error(err)
	// 	return response, err
	// }

	if err := json.NewEncoder(&buf).Encode(request.Body); err != nil {
		e.zapLogger.Zap.Error(err)
		return response, err
	}

	res, err := e.Client.Search(
		e.Client.Search.WithContext(context.Background()),
		e.Client.Search.WithIndex(request.Index),
		e.Client.Search.WithBody(&buf),
		e.Client.Search.WithTrackTotalHits(true),
		e.Client.Search.WithPretty(),
	)

	if err != nil {
		e.zapLogger.Zap.Error(err)
		return response, err
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
		return response, err
	}

	// Print the response status, number of results, and request duration.
	// e.zapLogger.Zap.Info("Data is empty!!")
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(dataTrx["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(dataTrx["took"].(float64)),
	)
	total := dataTrx["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)
	log.Println(int(total))
	if total == 0 {
		return response, err
	}
	data := dataTrx["hits"].(map[string]interface{})["hits"].([]interface{})
	return data, err
}

func (e Elasticsearch) Delete(request RequestElastic) (response bool, err error) {
	reqDelete := esapi.DeleteRequest{
		Index:      request.Index,
		DocumentID: request.DocumentID,
	}
	ctx, cancel := context.WithTimeout(context.Background(), e.timeout)
	defer cancel()
	resDelete, err := reqDelete.Do(ctx, e.Client)
	if err != nil {
		e.zapLogger.Zap.Error(err)
		return false, err
	}
	defer resDelete.Body.Close()

	if resDelete.IsError() {
		e.zapLogger.Zap.Error(err)
		return false, err
	}
	return true, err
}
