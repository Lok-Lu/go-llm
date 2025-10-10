package general

import (
	"context"
	"net/http"

	"github.com/spf13/cast"
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

	header, err := c.sendRequest(ctx, req, &response)
	if err != nil {
		return
	}

	promptTokens := header.Get("x-prompt-tokens")

	response.Details.InputLength = cast.ToInt32(promptTokens)
	return
}
