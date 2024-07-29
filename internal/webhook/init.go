package webhook

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"webdetect/internal/logger"
)

type WebhookInfo struct {
	Url                  string    `json:"url"`
	HasCustomCertificate bool      `json:"has_custom_certificate"`
	PendingUpdateCount   int       `json:"pending_update_count"`
	IpAddress            *string   `json:"ip_address"`
	LastErrorDate        *int      `json:"last_error_date"`
	LastErrorMessage     *string   `json:"last_error_message"`
	MaxConnections       *int      `json:"max_connections"`
	AllowedUpdates       *[]string `json:"allowed_updates"`
}

type NewWebhook struct {
	Url            string  `json:"url"`
	IpAddress      *string `json:"ip_address"`
	MaxConnections *int    `json:"max_connections"`
	SecretTokenm   *string `json:"secret_token"`
}

func SetupWebhook() error {
	if CheckHasWebhook() {
		return nil
	}

	botToken := os.Getenv("TG_BOT_TOKEN")
	if botToken == "" {
		botToken = "123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11"
		//return fmt.Errorf("TG_BOT_TOKEN is not set")
	}
	logger.Log("Bot token: ", botToken)

	apiURL := os.Getenv("TG_API_URL")
	if apiURL == "" {
		apiURL = "http://127.0.0.1:8989"
		//return fmt.Errorf("TG_API_URL is not set")
	}
	apiURL = strings.TrimSuffix(apiURL, "/")

	apiURL = apiURL + "/bot" + botToken + "/setWebhook"

	domain := os.Getenv("DOMAIN_NAME")
	if domain == "" {
		domain = "https://example.com"
	}
	domain = strings.TrimSuffix(domain, "/")
	domain = strings.TrimPrefix(domain, "https://")

	webhook := NewWebhook{
		Url: "https://" + domain,
	}

	form := url.Values{}
	form.Add("url", webhook.Url)

	logger.Log("Sending request to: ", apiURL)
	logger.Log("Prarams: ", form)
	// Send the POST request
	resp, err := http.PostForm(apiURL, form)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func CheckHasWebhook() bool {
	botToken := os.Getenv("TG_BOT_TOKEN")
	if botToken == "" {
		botToken = "123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11"
	}

	apiURL := os.Getenv("TG_API_URL")
	if apiURL == "" {
		apiURL = "http://127.0.0.1:8989"
	}
	apiURL = strings.TrimSuffix(apiURL, "/")

	url := apiURL + "/bot" + botToken + "/getWebhookInfo"

	domain := os.Getenv("DOMAIN_NAME")
	domain = strings.TrimSuffix(domain, "/")
	domain = strings.TrimPrefix(domain, "https://")

	domain = "https://" + domain

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	var webhookInfo WebhookInfo
	err = json.NewDecoder(resp.Body).Decode(&webhookInfo)
	if err != nil {
		return false
	}

	if webhookInfo.Url == domain {
		return true
	}

	return false
}
