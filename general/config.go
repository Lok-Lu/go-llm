package general

import (
	"net/http"
)

const EmptyMessagesLimit = 300

// ClientConfig for pharm
type ClientConfig struct {
	BaseURL            string
	HTTPClient         *http.Client
	EmptyMessagesLimit uint
}

func DefaultConfig(baseUrl string) ClientConfig {
	return ClientConfig{
		BaseURL:            baseUrl,
		HTTPClient:         &http.Client{},
		EmptyMessagesLimit: EmptyMessagesLimit,
	}
}

func (c ClientConfig) WithHttpClientConfig(client *http.Client) ClientConfig {
	c.HTTPClient = client
	return c
}
