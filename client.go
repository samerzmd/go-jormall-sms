package sms

import (
    "net/http"
)

// JormallConfig holds configuration details for Jormall SMS Gateway.
type JormallConfig struct {
    BaseURL         string
    AccountName     string
    AccountPassword string
    SenderID        string
}

// JormallClient represents a client for Jormall SMS service.
type JormallClient struct {
    Config     *JormallConfig
    HTTPClient *http.Client
}

// NewJormallClient creates a new Jormall SMS client with the given configuration.
func NewJormallClient(config *JormallConfig) *JormallClient {
    return &JormallClient{
        Config:     config,
        HTTPClient: &http.Client{},
    }
}
