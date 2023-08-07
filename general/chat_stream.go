package general

import (
	"bufio"
	"context"
	"net/http"
)

func (c *Client) CreateChatCompletionStream(
	ctx context.Context,
	url string,
	request *ChatRequest,
) (stream *ChatCompletionStream, err error) {
	urlSuffix := "/generate_stream"

	req, err := c.requestBuilder.Build(ctx, http.MethodPost, c.fullURL(url, urlSuffix), request)

	resp, err := c.sendStreamRequest(ctx, req)
	if err != nil {
		return
	}

	stream = &ChatCompletionStream{
		StreamReader: NewStreamReader[ChatStreamResponse](c.config.EmptyMessagesLimit, bufio.NewReader(resp.Body), resp),
	}
	return
}
