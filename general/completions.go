package general

import (
	"context"
	"net/http"
)

// for completions frame 
func (c *Client) CreateCompletion(
	ctx context.Context,
	url string,
	request *ChatRequest,
) (response *ChatResponse, err error) {

	urlSuffix := "/completions"
	req, err := c.requestBuilder.Build(ctx, http.MethodPost, c.fullURL(url, urlSuffix), request)
	if err != nil {
		return
	}

	err = c.sendRequest(ctx, req, &response)
	return
}
