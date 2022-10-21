package service

import (
	items "items_api/domain/items"
	es_queries "items_api/domain/queries"
	restError "items_api/utils/errors"
)

var (
	ItemsService itemServiceInterface = &itemsService{}
)

type itemServiceInterface interface {
	Create(items.Item) (*items.Item, restError.RestError)
	Get(string) (*items.Item, restError.RestError) 
	Search(query es_queries.EsQuery) ([]items.Item, restError.RestError)
	Update(items.Item) (*items.Item, restError.RestError)
	Delete(string) restError.RestError
}

type itemsService struct {}

func (s *itemsService) Create(itemRequest items.Item) (*items.Item, restError.RestError) {
	if err := itemRequest.Save(); err != nil {
		return nil, err
	}

	return &itemRequest,nil
}

func (s *itemsService) Get(id string) (*items.Item, restError.RestError) {
	item := items.Item{Id:id}

	if err := item.Get(); err != nil {
		return nil, err
	}


	return &item,nil
}

func (s *itemsService) Search(query es_queries.EsQuery) ([]items.Item, restError.RestError) {
	dao := items.Item{}
	
	return dao.Search(query)
	
}

func (s *itemsService) Update(newItem items.Item) (*items.Item, restError.RestError) {
	currentItem, err := s.Get(newItem.Id)
	if err != nil {
		return nil, err
	}

	if newItem.Seller > 0 {
		currentItem.Seller = newItem.Seller
	}

	if newItem.Title != "" {
		currentItem.Title = newItem.Title
	}

	if newItem.AvailableQuantity > 0 {
		currentItem.AvailableQuantity = newItem.AvailableQuantity
	}

	if newItem.Description.Html != "" {
		currentItem.Description.Html = newItem.Description.Html
	}

	if newItem.Description.PlainText != "" {
		currentItem.Description.PlainText = newItem.Description.PlainText
	}

	if len(newItem.Picture) > 0 {
		currentItem.Picture = newItem.Picture
	}

	if newItem.Video != "" {
		currentItem.Video = newItem.Video
	}

	if newItem.Price > 0 {
		currentItem.Price = newItem.Price
	}

	if newItem.SoldQuantity > 0 {
		currentItem.SoldQuantity = newItem.SoldQuantity
	}

	if newItem.Status != "" {
		currentItem.Status = newItem.Status
	}

	if err := currentItem.Update(); err != nil {
		return nil,err
	}

	return currentItem,nil

}

func (s *itemsService) Delete(id string) restError.RestError {
	item := &items.Item{Id: id}

	return item.Delete()
}