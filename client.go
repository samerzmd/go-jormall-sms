package sms

import (
	"net/http"
)

// Config holds configuration details for the JoSMS Gateway.
type Config struct {
	BaseURL         string // e.g., https://www.josms.net
	AccountName     string
	AccountPassword string
	SenderID        string
	RequestTimeout  int // Only used for bulk SMS
}

// JormallClient represents a client for JoSMS Gateway.
type JormallClient struct {
	Config     Config
	HTTPClient *http.Client
}

// NewJormallClient creates a new JoSMS client with the given configuration.
func NewJormallClient(cfg Config, httpClient *http.Client) *JormallClient {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	return &JormallClient{
		Config:     cfg,
		HTTPClient: httpClient,
	}
}
