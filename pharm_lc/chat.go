package pharm_lc

import (
	"context"
	"fmt"
	"net/http"
)

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	Response string `json:"response"`
	Status   int    `json:"status"`
	Time     string `json:"time"`
}

func (c *Client) CreateChat(
	ctx context.Context,
	request *ChatRequest,
) (response *ChatResponse, err error) {

	urlSuffix := ""
	req, err := c.requestBuilder.Build(ctx, http.MethodPost, c.fullURL(urlSuffix), request)
	fmt.Println(err)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
