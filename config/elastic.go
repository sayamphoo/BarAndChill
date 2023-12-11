package config

import (
	"context"
	"os"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
)

var once sync.Once
var Database *elasticsearch.TypedClient

func initializeElasticsearch() {
	var err error

	host := os.Getenv("ELASTIC_HOST")
	user := os.Getenv("ELASTIC_USERNAME")
	pass := os.Getenv("ELASTIC_PASSWORD")

	Database, err = elasticsearch.NewTypedClient(
		elasticsearch.Config{
			Addresses: []string{
				host,
			},
			Username: user,
			Password: pass,
		},
	)

	if err != nil {
		panic("Elastic Config Error")
	}

	if ping() != nil {
		panic("Elastic Config Error")
	}
}

func PingElasticsearchClient() *elasticsearch.TypedClient {
	once.Do(initializeElasticsearch)

	if ping() != nil {
		return nil
	}

	return Database
}

func ping() error {
	_, err := Database.Ping().Do(context.Background())
	return err
}
