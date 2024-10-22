package general

import (
	"context"
	"net/http"
)

type RerankRequest struct {
	Model string   `json:"model"`
	Query string   `json:"query"`
	Texts []string `json:"texts"`
}

type ReranKResponse struct {
	Object string   `json:"object"`
	Data   []Rerank `json:"data"`
	Model  string   `json:"model"`
}

type Rerank struct {
	Index  int     `json:"index"`
	Score  float32 `json:"score"`
	Object string  `json:"object"`
}

type OriginalReranKResponse []Rerank

func (e OriginalReranKResponse) ToReranKResponse(model string) *ReranKResponse {
	var data []Rerank

	for index, rerank := range e {
		data = append(data, Rerank{
			Object: "rerank",
			Index:  index,
			Score:  rerank.Score,
		})
	}

	return &ReranKResponse{
		Object: "list",
		Data:   data,
		Model:  model,
	}
}

// for Text Generation Inference（tgi） frame
func (c *Client) CreateRerank(
	ctx context.Context,
	url string,
	request *RerankRequest,
) (response *ReranKResponse, err error) {

	urlSuffix := "/rerank"
	req, err := c.requestBuilder.Build(ctx, http.MethodPost, c.fullURL(url, urlSuffix), request)
	if err != nil {
		return
	}

	var originalReranKResponse OriginalReranKResponse
	err = c.sendRequest(ctx, req, &originalReranKResponse)
	return originalReranKResponse.ToReranKResponse(request.Model), err
}
