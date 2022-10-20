package service

import (
	items "items_api/domain/items"
	restError "items_api/utils/errors"
)

var (
	ItemsService itemServiceInterface = &itemsService{}
)

type itemServiceInterface interface {
	Create(items.Item) (*items.Item, restError.RestError)
	Get(string) (*items.Item, restError.RestError) 
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