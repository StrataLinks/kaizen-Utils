package handle

import (
	"context"
	"time"

	"github.com/StrataLinks/kaizen-Utils/internal/logger"
	"github.com/olivere/elastic/v7"
)

// ElasticsearchHandler indexes log messages to Elasticsearch.
type ElasticsearchHandler struct {
	client *elastic.Client
	index  string
}

// NewElasticsearchHandler creates a new handler for Elasticsearch.
func NewElasticsearchHandler(url, index string) (*ElasticsearchHandler, error) {
	client, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		return nil, err
	}

	return &ElasticsearchHandler{
		client: client,
		index:  index,
	}, nil
}

// Log sends a log message to Elasticsearch.
func (e *ElasticsearchHandler) Log(level logger.LogLevel, message string) {
	logEntry := map[string]interface{}{
		"timestamp": time.Now(),
		"level":     level.String(),
		"message":   message,
	}
	_, err := e.client.Index().
		Index(e.index).
		BodyJson(logEntry).
		Do(context.Background())
	if err != nil {
		// Handle error
	}
}
