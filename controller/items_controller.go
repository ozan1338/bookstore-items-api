package controller

import (
	"encoding/json"
	"io/ioutil"
	oauth "items_api/api/oauth"
	"items_api/domain/items"
	"items_api/service"
	restError "items_api/utils/errors"
	"items_api/utils/http_utils"
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
		http_utils.ResponseError(w, err)
		return
	}

	sellerId := oauth.GetCienId(r)

	if sellerId == 0 {
		respErr := restError.NewUnauthorizedError("invalid credential")
		http_utils.ResponseError(w,respErr)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respError := restError.NewBadRequestError("invalid request body")
		http_utils.ResponseError(w, respError)
		return
	}

	defer r.Body.Close()

	var itemRequest items.Item

	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respError := restError.NewBadRequestError("invalid json body")
		http_utils.ResponseError(w, respError)
		return
	}

	itemRequest.Seller = sellerId

	result, saveErr := service.ItemsService.Create(itemRequest)
	if err != nil {
		//TODO return error to the user
		http_utils.ResponseError(w, saveErr)
		return
	}

	//TODO: return created item with http status 201 - created
	http_utils.ResponseJson(w, http.StatusOK, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}