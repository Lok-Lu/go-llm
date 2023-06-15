package pharm_lc

import (
	"encoding/json"
	"fmt"
	myerr "github.com/patsnapops/go-llm/error"
	llm "github.com/patsnapops/go-llm/request"
	"io"
	"net/http"
)

type Client struct {
	config         ClientConfig
	requestBuilder llm.RequestBuilder
}

func NewClient(url string) *Client {
	config := DefaultConfig(url)
	return NewClientWithConfig(config)
}

func NewClientWithVersions(versions map[string]string) *Client {
	config := NewConfigWithVersion(versions)
	return NewClientWithConfig(config)
}

func NewClientWithConfig(config ClientConfig) *Client {
	return &Client{
		config:         config,
		requestBuilder: llm.NewRequestBuilder(),
	}
}

func (c *Client) sendRequest(req *http.Request, v any) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.authToken))

	// Check whether Content-Type is already set, Upload Files API requires
	// Content-Type == multipart/form-data
	contentType := req.Header.Get("Content-Type")
	if contentType == "" {
		req.Header.Set("Content-Type", "text/plain")
	}

	res, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return c.handleErrorResp(res)
	}

	return decodeResponse(res.Body, v)
}

func decodeResponse(body io.Reader, v any) error {
	if v == nil {
		return nil
	}

	if result, ok := v.(*string); ok {
		return decodeString(body, result)
	}
	return json.NewDecoder(body).Decode(v)
}

func decodeString(body io.Reader, output *string) error {
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	*output = string(b)
	return nil
}

func (c *Client) fullURL(suffix, version string) string {
	if version == "" {
		return fmt.Sprintf("%s%s", c.config.BaseURL, suffix)
	}
	return fmt.Sprintf("%s%s", c.config.VersionUrl[version], suffix)
}

func (c *Client) handleErrorResp(resp *http.Response) error {
	var errRes myerr.ErrorResponse
	errByte, _ := io.ReadAll(resp.Body)

	errRes.Error = &myerr.APIError{
		Code:           resp.StatusCode,
		Message:        string(errByte),
		HTTPStatusCode: resp.StatusCode,
	}
	return fmt.Errorf("error, status code: %d, message: %w", resp.StatusCode, errRes.Error)
}
