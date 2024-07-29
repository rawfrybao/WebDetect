package webhook

import (
	"strings"
	"webdetect/internal/api"
	"webdetect/internal/db"
	"webdetect/internal/logger"
)

func addSubscription(tgID int64, name, url, xpath string) error {

	err := db.NewSubscription(tgID, name, url, xpath)
	if err != nil {
		return err
	}

	return nil
}

func handleAddSubscription(req api.HandleUpdateJSONRequestBody, text string) {
	hasAccess, err := db.CheckHasAccess(*req.Message.From.ID)
	if err != nil {
		logger.Log(err.Error())
	}

	if !hasAccess {
		logger.Log("User does not have access")
	}

	if len(strings.Split(text, " ")) != 3 {
		logger.Log("Invalid number of arguments")
	}

	args := strings.Split(text, " ")
	name := args[0]
	url := args[1]
	xpath := args[2]

	err = addSubscription(*req.Message.From.ID, name, url, xpath)
	if err != nil {
		logger.Log(err.Error())
	}
}
