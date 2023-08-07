package general

import (
	"context"
	"net/http"
)

type ChatResponse struct {
	GeneratedText string `json:"generated_text"`
}

func (c *Client) CreateChat(
	ctx context.Context,
	request *ChatRequest,
) (response *ChatResponse, err error) {

	urlSuffix := "/generate"
	req, err := c.requestBuilder.Build(ctx, http.MethodPost, c.fullURL(urlSuffix), request)
	if err != nil {
		return
	}

	err = c.sendRequest(ctx, req, &response)
	return
}
