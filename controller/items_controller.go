package controller

import (
	"encoding/json"
	"io/ioutil"
	oauth "items_api/api/oauth"
	"items_api/domain/items"
	es_queries "items_api/domain/queries"
	"items_api/service"
	restError "items_api/utils/errors"
	"items_api/utils/http_utils"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var (
	ItemsController itemsControllerInterface = &itemsController{} 
)

type itemsControllerInterface interface{
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request) 
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
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
	vars := mux.Vars(r)

	itemId := strings.TrimSpace(vars["id"])

	item,err := service.ItemsService.Get(itemId)
	if err != nil {
		http_utils.ResponseError(w,err)
		return
	}

	http_utils.ResponseJson(w, http.StatusOK, item)
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	var query es_queries.EsQuery

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := restError.NewBadRequestError("invalid json body")
		http_utils.ResponseError(w,apiErr)
		return
	}

	defer r.Body.Close()

	if err := json.Unmarshal(bytes, &query);err != nil {
		apiErr := restError.NewBadRequestError("invalid json body when unmarshall")
		http_utils.ResponseError(w,apiErr)
		return
	}


	items, searchErr := service.ItemsService.Search(query); 
	if searchErr != nil {
		http_utils.ResponseError(w,searchErr)
		return
	}

	http_utils.ResponseJson(w, http.StatusOK, items)
}

func (c *itemsController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	itemId := strings.TrimSpace(vars["id"])

	var newItem items.Item

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := restError.NewBadRequestError("invalid json body")
		http_utils.ResponseError(w,apiErr)
		return
	}

	defer r.Body.Close()

	if err := json.Unmarshal(bytes, &newItem); err != nil {
		apiErr := restError.NewBadRequestError("invalid json body when unmarshall")
		http_utils.ResponseError(w,apiErr)
		return
	}

	newItem.Id = itemId

	item,updateErr := service.ItemsService.Update(newItem)

	if updateErr != nil {
		http_utils.ResponseError(w,updateErr)
		return
	}

	http_utils.ResponseJson(w,http.StatusOK, item)
}

func (c *itemsController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	itemId := strings.TrimSpace(vars["id"])

	if err := service.ItemsService.Delete(itemId); err != nil {
		http_utils.ResponseError(w,err)
		return
	}

	http_utils.ResponseJson(w, http.StatusOK, map[string]interface{}{"status":"deleted"})
}