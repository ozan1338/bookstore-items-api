package app

import (
	"items_api/client/elasticsearch"
	"net/http"
	"time"

	"items_api/log"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()
	mapUrl()

	srv := &http.Server{
        Handler:      router,
        Addr:         "127.0.0.1:8000",
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 500 * time.Second,
        ReadTimeout:  2 * time.Second,
		IdleTimeout: 60 * time.Second,
    }

	log.Info("about to start application")

    if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}