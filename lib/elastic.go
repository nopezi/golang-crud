package lib

import (
	"crypto/tls"
	"net/http"

	elastic "github.com/elastic/go-elasticsearch/v8"
)

type Elasticsearch struct {
	Client *elastic.Client
	// index  string
	// alias  string
}

func NewElastic(env Env, zapLogger Logger) Elasticsearch {
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

	client, err := elastic.NewClient(cfg)

	if err != nil {
		zapLogger.Zap.Info("Url: ", url)
		zapLogger.Zap.Panic(err)
		zapLogger.Zap.Info("elastic connection failed")
	}
	zapLogger.Zap.Info("elastic connection established")
	return Elasticsearch{
		Client: client,
	}
}

// func (e *Elasticsearch) CreateIndex(index string) error {
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
