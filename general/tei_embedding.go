package general

import (
	"context"
	"net/http"
)

type Embedding struct {
	Object    string `json:"object"`
	Embedding any    `json:"embedding"` // []float32、string
	Index     int    `json:"index"`
}

type EmbeddingResponse struct {
	Object string      `json:"object"`
	Data   []Embedding `json:"data"`
	Model  string      `json:"model"`
	Usage  Usage       `json:"usage"`
}

type OriginalEmbeddingResponse [][]float32

func (e OriginalEmbeddingResponse) ToEmbeddingResponse(model string) *EmbeddingResponse {
	var data []Embedding

	for index, embedding := range e {
		data = append(data, Embedding{
			Object:    "embedding",
			Embedding: embedding,
			Index:     index,
		})
	}

	return &EmbeddingResponse{
		Object: "list",
		Data:   data,
		Model:  model,
	}
}

type EmbeddingRequest struct {
	Model  string `json:"model"`
	Inputs any    `json:"inputs"`
}

// for Text Generation Inference（tgi） frame
func (c *Client) CreateEmbedding(
	ctx context.Context,
	url string,
	request *EmbeddingRequest,
) (response *EmbeddingResponse, err error) {

	urlSuffix := "/embed"
	req, err := c.requestBuilder.Build(ctx, http.MethodPost, c.fullURL(url, urlSuffix), request)
	if err != nil {
		return
	}

	var originalEmbeddingResponse OriginalEmbeddingResponse
	err = c.sendRequest(ctx, req, &originalEmbeddingResponse)
	if err != nil {
		return
	}
	return originalEmbeddingResponse.ToEmbeddingResponse(request.Model), nil
}

type OpenaiEmbeddingRequest struct {
	Input          any    `json:"input"`
	Model          string `json:"model"`
	User           string `json:"user"`
	EncodingFormat string `json:"encoding_format,omitempty"`
	Dimensions     int    `json:"dimensions,omitempty"`
}

func (c *Client) CreateEmbeddingLikeOpenai(
	ctx context.Context,
	url string,
	request *OpenaiEmbeddingRequest,
) (response *EmbeddingResponse, err error) {
	urlSuffix := "/v1/embeddings"
	req, err := c.requestBuilder.Build(ctx, http.MethodPost, c.fullURL(url, urlSuffix), request)
	if err != nil {
		return
	}

	err = c.sendRequest(ctx, req, &response)

	if err != nil {
		return
	}
	return
}
