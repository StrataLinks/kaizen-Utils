package handle

import "github.com/olivere/elastic/v7"

type DatadogHandler struct {
	client *elastic.Client
	index  string
}
