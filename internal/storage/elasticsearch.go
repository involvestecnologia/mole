package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/involvestecnologia/mole/errors"
	"github.com/involvestecnologia/mole/internal/collectors"
	"github.com/involvestecnologia/mole/models"
)

type elasticsearch struct {
	client      *elasticsearch7.Client
	source      string
	batchSize   int
	currentSize int
	buffer      *bytes.Buffer
	collector   collectors.StorageCollector
}

//Elasticsearch - Opens communication with elasticsearch
func Elasticsearch(conf models.Elasticsearch, collector collectors.StorageCollector) Storage {

	cfg := elasticsearch7.Config{
		Addresses: []string{conf.Hosts},
		Username:  conf.Username,
		Password:  conf.Password,
	}

	client, err := elasticsearch7.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &elasticsearch{
		client:      client,
		source:      conf.Source,
		batchSize:   conf.BatchSize,
		currentSize: 0,
		buffer:      &bytes.Buffer{},
		collector:   collector,
	}
}

//Add - Adds the record to the batch and writes it to elasticsearch when the quantity is complete
func (e *elasticsearch) Add(oplog *models.Oplog) error {

	resource, err := e.prepare(oplog)
	if err != nil {
		return err
	}

	routing := resource.Timestamp.Format("02-01-2006")

	meta := []byte(fmt.Sprintf(`{ "index" : { "_index" : "%s", "_type": "_doc", "_routing": "%s" } }%s`, e.source, routing, "\n"))
	data, err := json.Marshal(&resource)
	if err != nil {
		return err
	}

	data = append(data, "\n"...)
	e.buffer.Grow(len(meta) + len(data))
	e.buffer.Write(meta)
	e.buffer.Write(data)

	e.currentSize++

	if e.currentSize == e.batchSize {
		res, err := e.client.Bulk(bytes.NewReader(e.buffer.Bytes()), e.client.Bulk.WithIndex(e.source))
		if err != nil {
			return err
		}

		defer res.Body.Close()

		if res.IsError() {
			return &errors.CouldNotSaveOplogOnElasticsearch{Message: res.Body}
		}

		e.buffer.Reset()
		e.currentSize = 0
		e.collector.IncreasesStorageMetrics(e.batchSize)
	}

	return nil
}

func (e *elasticsearch) prepare(oplog *models.Oplog) (*models.OplogAnalysis, error) {

	query, err := json.Marshal(oplog.Object)
	if err != nil {
		return &models.OplogAnalysis{}, err
	}

	oplogAnalysis := &models.OplogAnalysis{}
	oplogAnalysis.Operation = oplog.GetOperationName()
	oplogAnalysis.Query = string(query)
	oplogAnalysis.Timestamp = time.Unix(int64(oplog.Timestamp.T), 0)

	namespace := strings.Split(oplog.Namespace, ".")
	if len(namespace) > 1 {
		oplogAnalysis.Database = namespace[0]
		oplogAnalysis.Collection = namespace[1]
	}

	return oplogAnalysis, nil
}

//StartTime - Searches for the date of the last stored record
func (e *elasticsearch) StartTime() (time.Time, error) {
	var buf bytes.Buffer
	var result models.SearchResults

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"size": 1,
		"sort": map[string]interface{}{
			"timestamp": "desc",
		},
	}

	if err := json.NewEncoder(&buf).Encode(&query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := e.client.Search(
		e.client.Search.WithContext(context.Background()),
		e.client.Search.WithIndex(e.source),
		e.client.Search.WithBody(&buf),
		e.client.Search.WithTrackTotalHits(true),
		e.client.Search.WithPretty(),
	)

	defer res.Body.Close()

	if err != nil {
		return time.Time{}, fmt.Errorf("error getting response: %s", err)
	}

	if res.StatusCode != 200 {
		return time.Time{}, &errors.CouldNotReadLastTimeOnElasticsearch{Message: res.Body}
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return time.Time{}, err
	}

	if len(result.Hits.Querys) > 0 {
		return time.Parse(time.RFC3339, result.Hits.Querys[0].Source.Timestamp)
	}

	return time.Now(), nil
}
