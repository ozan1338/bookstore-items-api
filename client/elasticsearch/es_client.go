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
	Get(string,string) (*elastic.GetResult, error)
	Search(string, elastic.Query) (*elastic.SearchResult, error)
	Update(string,string, interface{}) (*elastic.UpdateResponse, error)
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
		log.Error(fmt.Sprintf("error when trying to index document in index: %s",index), err)
		return nil,err
	}
	return result, nil
}

func (c *esClient) Get(index string,id string) (*elastic.GetResult,error) {
	ctx := context.Background()
	result, err := c.client.Get().
    Index(index).
    Id(id).
    Do(ctx)

	if err != nil {
		log.Error(fmt.Sprintf("error when trying to get id %s",id), err)
		return nil,err
	}

	return result,nil
}

func (c *esClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()

	result, err := c.client.Search(index).Query(query).RestTotalHitsAsInt(true).Do(ctx)

	if err != nil {
		log.Error(fmt.Sprintf("error when trying to search document in index %s",index), err)
		return nil,err
	}

	return result, nil
}

func (c *esClient) Update(index, id string, newItems interface{}) (*elastic.UpdateResponse, error) {
	ctx := context.Background()

	result, err := c.client.Update().Index(index).Id(id).Doc(newItems).Do(ctx)

	if err != nil {
		log.Error(fmt.Sprintf("error when trying to update document in id %s",id), err)
		return nil,err
	}

	return result,nil
}