package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
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
		return fmt.Errorf("TG_BOT_TOKEN is not set")
	}

	apiURL := os.Getenv("TG_API_URL")
	if apiURL == "" {
		return fmt.Errorf("TG_API_URL is not set")
	}
	apiURL = strings.TrimSuffix(apiURL, "/")

	url := apiURL + "/bot" + botToken + "/setWebhook"

	domain := os.Getenv("DOMAIN_NAME")
	domain = strings.TrimSuffix(domain, "/")
	domain = strings.TrimPrefix(domain, "https://")

	webhook := NewWebhook{
		Url: "https://" + domain,
	}

	jsonData, err := json.Marshal(webhook)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func CheckHasWebhook() bool {
	botToken := os.Getenv("TG_BOT_TOKEN")
	if botToken == "" {
		return false
	}

	apiURL := os.Getenv("TG_API_URL")
	if apiURL == "" {
		return false
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
