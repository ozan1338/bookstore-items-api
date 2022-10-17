package service

import (
	items "items_api/domain/items"
	restError "items_api/utils/errors"
	"net/http"
)

var (
	ItemsService itemServiceInterface = &itemsService{}
)

type itemServiceInterface interface {
	Create(items.Item) (*items.Item, *restError.RestError)
	Get(string) (*items.Item, *restError.RestError) 
}

type itemsService struct {}

func (s *itemsService) Create(items.Item) (*items.Item, *restError.RestError) {
	return nil,&restError.RestError{Message: "not ready",Status: http.StatusNotImplemented, Error: "not_implemented"}
}

func (s *itemsService) Get(title string) (*items.Item, *restError.RestError) {
	return nil,&restError.RestError{Message: "not ready",Status: http.StatusNotImplemented, Error: "not_implemented"}
}