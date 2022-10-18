package items

import (
	"items_api/client/elasticsearch"
	restError "items_api/utils/errors"
)

const (
	indexItems = "items"
)

func (i *Item) Save() *restError.RestError {
	result,err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return restError.NewInternalServerError("error when trying to save item")
	}

	i.Id = result.Id
	return nil
}