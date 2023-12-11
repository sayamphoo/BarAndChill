package repo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sayamphoo/microservice/config"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const idKey = "_id"

type Database struct {
	Index   string
	Elastic *elasticsearch.TypedClient
}

// ConnectElastic establishes a connection to Elasticsearch.
func (db *Database) ConnectElastic() {
	db.Elastic = config.PingElasticsearchClient()
}

// RepoSave saves data to the database.
func (db *Database) RepoSave(entity interface{}) (string, error) {
	db.ConnectElastic()

	result, err := db.Elastic.
		Index(db.Index).
		Request(entity).
		Do(context.TODO())

	if err != nil {
		return "", fmt.Errorf("failed to save document: %w", err)
	}

	if result.Id_ != "" {
		return result.Id_, nil
	}

	return "", errors.New("failed to save document")
}

// RepoFindByID finds data by ID.
func (db *Database) RepoFindByID(id string) (*types.HitsMetadata, error) {
	db.ConnectElastic()
	return db.RepoFindByWord(idKey, id)
}

// RepoFindByWord finds data by a specified key and value.
func (db *Database) RepoFindByWord(key string, value string) (*types.HitsMetadata, error) {
	db.ConnectElastic()
	querySize := 1000
	result, err := db.Elastic.Search().
		Index(db.Index).
		Request(&search.Request{
			Size: &querySize,
			Query: &types.Query{
				Match: map[string]types.MatchQuery{
					key: {Query: value},
				},
			},
		}).Do(context.Background())

	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	if result.Hits.Total.Value > 0 {
		return &result.Hits, nil
	}
	return nil, fmt.Errorf("%s not found", key)
}

// RepoFindByWordRangeLte finds data with a key-value range condition.
func (db *Database) RepoFindByWordRangeLte(key string, date string) (*types.HitsMetadata, error) {
	db.ConnectElastic()
	querySize := 1000
	result, err := db.Elastic.Search().
		Index(db.Index).
		Request(&search.Request{
			Size: &querySize,
			Query: &types.Query{
				Range: map[string]types.RangeQuery{
					key: map[string]interface{}{
						"lte": date,
					},
				},
			},
		}).Do(context.Background())

	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	if result.Hits.Total.Value > 0 {
		return &result.Hits, nil
	}
	return nil, fmt.Errorf("%s not found", key)
}

// RepoUpdating updates data by ID.
func (db *Database) RepoUpdating(id string, wrapper interface{}) (*update.Response, error) {
	db.ConnectElastic()

	result, err := db.Elastic.Update(db.Index, id).Doc(wrapper).Do(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("update error: %w", err)
	}
	return result, nil
}

// RepoGetIndex retrieves all data from the index.
func (db *Database) RepoGetIndex() (*types.HitsMetadata, error) {
	path := fmt.Sprintf("%s/%s/_search?size=1000", os.Getenv("ELASTIC_HOST"), db.Index)
	result, err := http.Get(path)
	if err != nil {
		return nil, fmt.Errorf("update error: %w", err)
	}
	defer result.Body.Close()

	body, _ := io.ReadAll(result.Body)
	var res search.Response
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	return &res.Hits, nil
}
