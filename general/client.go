package general

import (
	"context"
	"fmt"
	"io"
	"net/http"

	wraperr "github.com/Lok-Lu/go-llm/error"
	. "github.com/Lok-Lu/go-llm/internal"
)

type Client struct {
	config         ClientConfig
	requestBuilder RequestBuilder
}

func NewClient(url, token string) *Client {
	config := DefaultConfig(url, token)
	return NewClientWithConfig(config)
}

func NewClientWithConfig(config ClientConfig) *Client {
	return &Client{
		config:         config,
		requestBuilder: NewRequestBuilder(config.HTTPClient),
	}
}

func (c *Client) SetUrl(url string) *Client {
	c.config.BaseURL = url
	return c
}

func (c *Client) sendRequest(ctx context.Context, req *http.Request, v any) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.authToken))

	// Check whether Content-Type is already set, Upload Files API requires
	// Content-Type == multipart/form-data
	contentType := req.Header.Get("Content-Type")
	if contentType == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.config.authToken != "" {
		// for eas
		req.Header.Set("Authorization", c.config.authToken)
	}

	err := c.requestBuilder.Send(ctx, req, v)
	return err
}

func (c *Client) sendStreamRequest(ctx context.Context, req *http.Request) (*http.Response, error) {

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")

	return c.requestBuilder.SendNoClose(ctx, req)
}

func (c *Client) fullURL(url, suffix string) string {
	if url != "" {
		return fmt.Sprintf("%s%s", url, suffix)
	}
	return fmt.Sprintf("%s%s", c.config.BaseURL, suffix)
}

func (c *Client) handleErrorResp(resp *http.Response) error {
	var errRes wraperr.ErrorResponse
	errByte, _ := io.ReadAll(resp.Body)
	errRes.Error = &wraperr.APIError{
		Code:           resp.StatusCode,
		Message:        string(errByte),
		HTTPStatusCode: resp.StatusCode,
	}
	return fmt.Errorf("error, status code: %d, message: %w", resp.StatusCode, errRes.Error)
}
