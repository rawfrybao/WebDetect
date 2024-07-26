package server

import (
	"net/http"
	"webdetect/internal/api"
	"webdetect/internal/webhook"
)

type handlers struct {
	webhook.UpdateHandler
}

func newHandler() *handlers {
	return &handlers{
		webhook.UpdateHandler{},
	}
}

func Start() {
	apiHandler := api.Handler(newHandler())
	http.ListenAndServe(":6969", apiHandler)
}
