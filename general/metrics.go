package general

import (
	"context"
	"io"
	"net/http"
	"time"
)

func (c *Client) GetMetrics(ctx context.Context, url string, timeout time.Duration) (response []byte, err error) {
	urlSuffix := "/metrics"
	req, err := c.requestBuilder.Build(ctx, http.MethodGet, c.fullURL(url, urlSuffix), nil)
	if err != nil {
		return
	}
	resp, err := c.requestBuilder.SendNoCloseWithCustomClient(ctx, &http.Client{
		Timeout: timeout,
	}, req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	response, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return response, nil
}
