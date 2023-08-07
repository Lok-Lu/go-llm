package general

import (
	"bufio"
	"context"
	"net/http"
)

type ChatStreamToken struct {
	ID      uint    `json:"id"`
	Text    string  `json:"text"`
	Logprob float64 `json:"logprob"`
	Special bool    `json:"special"`
}

type ChatStreamResponse struct {
	GeneratedText string          `json:"generated_text,omitempty"`
	Token         ChatStreamToken `json:"token"`
	Details       string          `json:"details"`
}

type ChatCompletionStream struct {
	*StreamReader[ChatStreamResponse]
}

func (c *Client) CreateChatCompletionStream(
	ctx context.Context,
	request *ChatRequest,
) (stream *ChatCompletionStream, err error) {
	urlSuffix := "/generate_stream"

	req, err := c.requestBuilder.Build(ctx, http.MethodPost, c.fullURL(urlSuffix), request)

	resp, err := c.sendStreamRequest(ctx, req)
	if err != nil {
		return
	}

	stream = &ChatCompletionStream{
		StreamReader: NewStreamReader[ChatStreamResponse](c.config.EmptyMessagesLimit, bufio.NewReader(resp.Body), resp),
	}
	return
}
