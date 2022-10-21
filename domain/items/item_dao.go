package items

import (
	"encoding/json"
	"fmt"
	"items_api/client/elasticsearch"
	es_queries "items_api/domain/queries"
	restError "items_api/utils/errors"
	"strings"
)

const (
	indexItems = "items"
)

func (i *Item) Save() restError.RestError {
	result,err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return restError.NewInternalServerError("error when trying to save item")
	}

	i.Id = result.Id
	return nil
}

func (i *Item) Get() restError.RestError {
	result, err := elasticsearch.Client.Get(indexItems,i.Id)

	itemId := i.Id

	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return restError.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.Id))
		}
		return restError.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id))
	}

	// if err := result.Source.UnmarshalJSON(i); err != nil {

	// }

	bytes, err := result.Source.MarshalJSON()
	
	if err != nil {
		return restError.NewInternalServerError("error when trying to parse database response")
	}

	if err := json.Unmarshal(bytes, i); err != nil {
		return restError.NewInternalServerError("error when trying to parse database response")
	}

	i.Id = itemId

	return nil
}

func (i *Item) Search(query es_queries.EsQuery) ([]Item, restError.RestError) {
	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	
	if err != nil {
		return nil, restError.NewInternalServerError("error when trying to search documents")
	}

	hitResult := make([]Item, result.TotalHits())

	for index, hit := range result.Hits.Hits {
		bytes, err := hit.Source.MarshalJSON()
		if err != nil {
			return nil, restError.NewInternalServerError("error when trying to MarshalJson")
		}

		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, restError.NewInternalServerError("error when trying to parse response")
		}
		item.Id = hit.Id
		hitResult[index] = item
	}

	if len(hitResult) == 0 {
		return nil, restError.NewNotFoundError("no items found matching given criteria")
	}

	return hitResult, nil
}

func (i *Item) Update() restError.RestError {
	_, err := elasticsearch.Client.Update(indexItems,i.Id,i)
	if err != nil {
		return restError.NewInternalServerError("error when trying to update item")
	}

	return nil
}

func (i *Item) Delete() restError.RestError {
	_,err := elasticsearch.Client.Delete(indexItems,i.Id)

	if err != nil {
		return restError.NewInternalServerError("Error when trying to delete item")
	}

	return nil
}