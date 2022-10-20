package items

import (
	"encoding/json"
	"fmt"
	"items_api/client/elasticsearch"
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