package webhook

import (
	"net/http"
	"strings"
	"webdetect/internal/api"

	"github.com/go-chi/render"
)

type UpdateHandler struct{}

func (h *UpdateHandler) HandleUpdate(w http.ResponseWriter, r *http.Request) *api.Response {
	var req api.HandleUpdateJSONRequestBody
	if err := render.Bind(r, &req); err != nil {
		code := http.StatusBadRequest
		return &api.Response{Code: code}
	}

	if !strings.HasPrefix(*req.Message.Text, "/") {
		code := http.StatusOK
		return &api.Response{Code: code}
	}

	command := strings.Split(*req.Message.Text, " ")[0]
	text := strings.Join(strings.Split(*req.Message.Text, " ")[1:], " ")

	switch command {
	case "/start":
		go handleStart(req)
	case "/fetch":
		go handleFetch(req, text)
	case "/listsub":
		go handleListSubscription(req)
	case "/addsub":
		go handleAddSubscription(req, text)
	case "/delsub":
		go handleDeleteSubscription(req, text)
	case "/giveaccess":
		go handleGiveAccess(req)
	}

	code := http.StatusOK
	return &api.Response{Code: code}
}

func NewUpdateHandler() *UpdateHandler {
	return &UpdateHandler{}
}
