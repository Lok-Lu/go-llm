package general

import (
	"context"
	"net/http"
)

type VllmEmbeddingRequest struct {
	Model           string   `json:"model"`
	InputType       string   `json:"input_type,omitempty"`
	Texts           []string `json:"texts,omitempty"`
	Images          []string `json:"images,omitempty"`
	Inputs          any      `json:"inputs,omitempty"`
	EmbeddingTypes  []string `json:"embedding_types,omitempty"`
	OutputDimension int      `json:"output_dimension,omitempty"`
	Truncate        string   `json:"truncate,omitempty"`
}

type VllmEmbeddingResponse struct {
	ID         string             `json:"id"`
	Embeddings VllmEmbeddingTypes `json:"embeddings"`
	Texts      []string           `json:"texts"`
	Meta       VllmEmbeddingMeta  `json:"meta"`
}

type VllmEmbeddingTypes struct {
	Float [][]float32 `json:"float"`
}

type VllmEmbeddingMeta struct {
	APIVersion  VllmAPIVersion  `json:"api_version"`
	BilledUnits VllmBilledUnits `json:"billed_units"`
}

type VllmAPIVersion struct {
	Version string `json:"version"`
}

type VllmBilledUnits struct {
	InputTokens int `json:"input_tokens"`
	ImageTokens int `json:"image_tokens"`
}

func (r *VllmEmbeddingResponse) ToEmbeddingResponse(model string) *EmbeddingResponse {
	var data []Embedding

	for index, embedding := range r.Embeddings.Float {
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
		Usage: Usage{
			PromptTokens: r.Meta.BilledUnits.InputTokens + r.Meta.BilledUnits.ImageTokens,
			TotalTokens:  r.Meta.BilledUnits.InputTokens + r.Meta.BilledUnits.ImageTokens,
		},
	}
}

// for vLLM Cohere-compatible embed API, supports multi-modal (texts and images)
func (c *Client) CreateMultiModalEmbedding(
	ctx context.Context,
	url string,
	request *VllmEmbeddingRequest,
) (response *EmbeddingResponse, err error) {

	urlSuffix := "/v2/embed"
	req, err := c.requestBuilder.Build(ctx, http.MethodPost, c.fullURL(url, urlSuffix), request)
	if err != nil {
		return
	}

	var vllmEmbeddingResponse VllmEmbeddingResponse
	_, err = c.sendRequest(ctx, req, &vllmEmbeddingResponse)
	if err != nil {
		return
	}
	return vllmEmbeddingResponse.ToEmbeddingResponse(request.Model), nil
}
