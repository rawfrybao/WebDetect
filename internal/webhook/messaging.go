package webhook

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"webdetect/internal/logger"
)

func SendMessage(message TgMessage) {
	botToken := os.Getenv("TG_BOT_TOKEN")
	botToken = strings.TrimPrefix(botToken, "bot")

	apiURL := os.Getenv("TG_API_URL")
	apiURL = strings.TrimSuffix(apiURL, "/")

	url := apiURL + "/bot" + botToken + "/sendMessage"

	log.Println(url)
	logger.Log(url)
	log.Println(message)
	logger.Log(message)

	jsonData, err := json.Marshal(message)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
}
