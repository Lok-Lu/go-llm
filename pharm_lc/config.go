package pharm_lc

import (
	"net/http"
)

// ClientConfig for pharm
type ClientConfig struct {
	BaseURL    string
	HTTPClient *http.Client
}

func DefaultConfig(baseUrl string) ClientConfig {
	return ClientConfig{
		BaseURL:    baseUrl,
		HTTPClient: &http.Client{},
	}
}

func (c ClientConfig) WithHttpClientConfig(client *http.Client) ClientConfig {
	c.HTTPClient = client
	return c
}
