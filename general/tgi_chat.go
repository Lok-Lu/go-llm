package general

import (
	"context"
	"net/http"
)

// for Text Generation Inference（tgi） frame
func (c *Client) CreateChat(
	ctx context.Context,
	url string,
	request *ChatRequest,
) (response *ChatResponse, err error) {

	urlSuffix := "/generate"
	req, err := c.requestBuilder.Build(ctx, http.MethodPost, c.fullURL(url, urlSuffix), request)
	if err != nil {
		return
	}

	err = c.sendRequest(ctx, req, &response)
	return
}
