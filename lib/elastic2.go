package lib

import (
	"crypto/tls"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
)

type ElasticSearch2 struct {
	Client *elasticsearch.Client
	index  string
	alias  string
}

func New(addresses []string, username string, password string) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: addresses,
		Username:  username,
		Password:  password,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return client, err
	}

	return client, nil
}

// func (e *ElasticSearch2) CreateIndex(index string) error {
// 	e.index = index
// 	e.alias = index + "_alias"

// 	res, err := e.client.Indices.Exists([]string{e.index})
// 	if err != nil {
// 		return fmt.Errorf("cannot check index existence: %w", err)
// 	}
// 	if res.StatusCode == 200 {
// 		return nil
// 	}
// 	if res.StatusCode != 404 {
// 		return fmt.Errorf("error in index existence response: %s", res.String())
// 	}

// 	res, err = e.client.Indices.Create(e.index)
// 	if err != nil {
// 		return fmt.Errorf("cannot create index: %w", err)
// 	}
// 	if res.IsError() {
// 		return fmt.Errorf("error in index creation response: %s", res.String())
// 	}

// 	res, err = e.client.Indices.PutAlias([]string{e.index}, e.alias)
// 	if err != nil {
// 		return fmt.Errorf("cannot create index alias: %w", err)
// 	}
// 	if res.IsError() {
// 		return fmt.Errorf("error in index alias creation response: %s", res.String())
// 	}

// 	return nil
// }

// // document represents a single document in Get API response body.
// type document struct {
// 	Source interface{} `json:"_source"`
// }