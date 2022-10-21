package app

import (
	"items_api/controller"
	"net/http"
)

func mapUrl() {
	router.HandleFunc("/ping", controller.PingController.Get).Methods(http.MethodGet)
	router.HandleFunc("/items", controller.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/search", controller.ItemsController.Search).Methods(http.MethodGet)
	router.HandleFunc("/items/get/{id}", controller.ItemsController.Get).Methods(http.MethodGet)
	router.HandleFunc("/items/update/{id}", controller.ItemsController.Update).Methods(http.MethodPatch)
	router.HandleFunc("/items/delete/{id}", controller.ItemsController.Delete).Methods(http.MethodDelete)
}