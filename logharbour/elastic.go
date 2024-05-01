package logharbour

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// ElasticsearchWriter defines methods for Elasticsearch writer
type ElasticsearchWriter interface {
	Write(index string, documentID string, body string) error
}

type ElasticsearchClient struct {
	client *elasticsearch.Client
}

// NewElasticsearchClient creates a new Elasticsearch client with the given configuration
func NewElasticsearchClient(cfg elasticsearch.Config) (*ElasticsearchClient, error) {
	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &ElasticsearchClient{client: esClient}, nil
}

// Write sends a document to Elasticsearch. It implements ElasticsearchWriter.
func (ec *ElasticsearchClient) Write(index string, documentID string, body string) error {
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: documentID,
		Body:       strings.NewReader(body),
	}

	res, err := req.Do(context.Background(), ec.client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("Error response from Elasticsearch: %s", res.String())
		return errors.New(res.String())
	}

	return nil
}

func (ec *ElasticsearchClient) CreateIndex(indexName, mapping string) error {
	res, err := ec.client.Indices.Create(indexName, ec.client.Indices.Create.WithBody(strings.NewReader(mapping)))
	if err != nil {
		return err
	}
	if res.IsError() {
		return fmt.Errorf("error creating index: %s", res.String())
	}
	return nil
}

func (ec *ElasticsearchClient) IndexExists(indexName string) (bool, error) {
	res, err := ec.client.Indices.Exists([]string{indexName})
	if err != nil {
		return false, err
	}
	return res.StatusCode == 200, nil
}

// Write sends a document to Elasticsearch with retry logic.
// func (ec *ElasticsearchClient) Write(index string, documentID string, body string) error {
// 	var maxAttempts = 5
// 	var initialBackoff = 1 * time.Second

// 	operation := func() error {
// 		req := esapi.IndexRequest{
// 			Index:      index,
// 			DocumentID: documentID,
// 			Body:       strings.NewReader(body),
// 		}

// 		res, err := req.Do(context.Background(), ec.client)
// 		if err != nil {
// 			return err
// 		}
// 		defer res.Body.Close()

// 		if res.IsError() {
// 			log.Printf("Error response from Elasticsearch: %s", res.String())
// 			return errors.New(res.String())
// 		}

// 		return nil
// 	}

// 	for attempt := 1; attempt <= maxAttempts; attempt++ {
// 		err := operation()
// 		if err == nil {
// 			return nil // Success
// 		}

// 		if attempt == maxAttempts {
// 			return fmt.Errorf("after %d attempts, last error: %s", attempt, err)
// 		}

// 		wait := initialBackoff * time.Duration(1<<(attempt-1)) // Exponential backoff
// 		log.Printf("Attempt %d failed, retrying in %v: %v", attempt, wait, err)
// 		time.Sleep(wait)
// 	}

// 	return fmt.Errorf("reached max attempts without success")
// }
