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

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Println("elastic connection failed")
		return client, err
	}
	fmt.Println("elastic connection established")
	return client, nil
}
