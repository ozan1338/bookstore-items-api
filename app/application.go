package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrl()

	srv := &http.Server{
        Handler:      router,
        Addr:         "127.0.0.1:8000",
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 500 * time.Second,
        ReadTimeout:  2 * time.Second,
		IdleTimeout: 60 * time.Second,
    }

    if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
	log.Println("server start at localhost:8000")
}