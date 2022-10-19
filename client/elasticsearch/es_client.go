package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"items_api/log"

	elastic "github.com/olivere/elastic/v7"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string,interface{}) (*elastic.IndexResponse,error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	logger := log.GetLogger()

	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(logger),
		elastic.SetInfoLog(logger),
	)

	if err != nil {
		panic(err)
	}

	Client.setClient(client)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index string,item interface{}) (*elastic.IndexResponse,error) {
	ctx := context.Background()
	result,err := c.client.Index().Index(index).BodyJson(item).Do(ctx)

	if err != nil {
		log.Error(fmt.Sprintf("error ewhen trying to index document in index: %s",index), err)
		return nil,err
	}
	return result, nil
}