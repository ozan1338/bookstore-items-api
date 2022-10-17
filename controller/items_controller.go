package controller

import (
	"fmt"
	oauth "items_api/api/oauth"
	"items_api/domain/items"
	"items_api/service"
	"net/http"
)

var (
	ItemsController itemsControllerInterface = &itemsController{} 
)

type itemsControllerInterface interface{
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request) 
}

type itemsController struct {}

func (c *itemsController)Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//TODO return error to the user
		return
	}

	item := items.Item{
		Seller: oauth.GetUserId(r),
	}

	result, err := service.ItemsService.Create(item)
	if err != nil {
		//TODO return error to the user
		return
	}

	fmt.Println(result)
	//TODO: return created item with http status 201 - created
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}