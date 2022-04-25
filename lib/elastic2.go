package lib

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
)

type ElasticSearch2 struct {
	Client *elasticsearch.Client
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

	client, _ := elasticsearch.NewClient(cfg)
	info, err := client.Info()
	if err != nil {
		LogChecklist("Elasticsearch Connection Refused", false)
		return client, err
	}
	fmt.Println("info", info)
	// fmt.Println("err", err)
	LogChecklist("Elasticsearch Connection Established", true)
	return client, nil
}
