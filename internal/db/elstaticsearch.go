package db

import (
	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	"sync"
)

var ES *elasticsearch.Client
var es_once sync.Once

func NewElasticsearch(address []string, user,password string) {
	var err error
	es_once.Do(func() {
		cfg := elasticsearch.Config{
			Addresses: address,
			Username:              user,
			Password:              password,
			CloudID:               "",
			APIKey:                "",
			RetryOnStatus:         nil,
			DisableRetry:          false,
			EnableRetryOnTimeout:  false,
			MaxRetries:            0,
			DiscoverNodesOnStart:  false,
			DiscoverNodesInterval: 0,
			EnableMetrics:         false,
			EnableDebugLogger:     false,
			RetryBackoff:          nil,
			Transport:             nil,
			Logger:                nil,
			Selector:              nil,
			ConnectionPoolFunc:    nil,
		}
		ES, err = elasticsearch.NewClient(cfg)
		if err != nil {
			panic(err)
		}
	})
}

