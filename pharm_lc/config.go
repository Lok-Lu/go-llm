package pharm_lc

import (
	"net/http"
)

// ClientConfig for pharm
type ClientConfig struct {
	BaseURL    string
	VersionUrl map[string]string
	HTTPClient *http.Client
}

func DefaultConfig(baseUrl string) ClientConfig {
	return ClientConfig{
		BaseURL:    baseUrl,
		HTTPClient: &http.Client{},
	}
}

func NewConfigWithVersion(versionUrl map[string]string) ClientConfig {
	return ClientConfig{
		VersionUrl: versionUrl,
		HTTPClient: &http.Client{},
	}
}

func (c ClientConfig) WithHttpClientConfig(client *http.Client) ClientConfig {
	c.HTTPClient = client
	return c
}
