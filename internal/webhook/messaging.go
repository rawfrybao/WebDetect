package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func SendMessage(message TgMessage) {
	botToken := os.Getenv("TG_BOT_TOKEN")
	botToken = strings.TrimPrefix(botToken, "bot")

	apiURL := os.Getenv("TG_API_URL")
	apiURL = strings.TrimSuffix(apiURL, "/")

	url := apiURL + "/bot" + botToken + "/sendMessage"

	fmt.Println(url)
	fmt.Println(message)

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
