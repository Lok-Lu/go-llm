package general

import (
	"context"
	"io"
	"net/http"
)

func (c *Client) GetMetrics(ctx context.Context, url string) (response []byte, err error) {
	urlSuffix := "/metrics"
	req, err := c.requestBuilder.Build(ctx, http.MethodGet, c.fullURL(url, urlSuffix), nil)
	if err != nil {
		return
	}
	resp, err := c.requestBuilder.SendNoClose(ctx, req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	response, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	return response, nil
}
