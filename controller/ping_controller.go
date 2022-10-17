package controller

import "net/http"

var (
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Get(w http.ResponseWriter,r *http.Request)
}

type pingController struct {}

func (c *pingController) Get(w http.ResponseWriter,r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ping"))
}

