package general

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type RerankInferenceFrame string

const (
	RerankInferenceFrameTEI      RerankInferenceFrame = "tei"
	RerankInferenceFrameInfinity RerankInferenceFrame = "infinity"
)

type RerankRequest struct {
	Model      string   `json:"model"`
	Query      string   `json:"query"`
	Texts      []string `json:"texts,omitempty"`
	RawScores  bool     `json:"raw_scores,omitempty" example:"false"`
	ReturnText bool     `json:"return_text,omitempty" example:"false"`
	// Documents      []string             `json:"documents,omitempty"`
	// ReturnDocument bool                 `json:"return_document,omitempty" example:"false"`
	TopN           int                  `json:"top_n,omitempty" example:"1"`
	InferenceFrame RerankInferenceFrame `json:"inference_frame"`
}

func (r *RerankRequest) MarshalJSON() ([]byte, error) {
	switch r.InferenceFrame {
	case RerankInferenceFrameTEI:
		return json.Marshal(r)
	case RerankInferenceFrameInfinity:
		req := struct {
			Model          string               `json:"model"`
			Query          string               `json:"query"`
			Texts          []string             `json:"documents,omitempty"`
			RawScores      bool                 `json:"raw_scores,omitempty" example:"false"`
			ReturnText     bool                 `json:"return_document,omitempty" example:"false"`
			TopN           int                  `json:"top_n,omitempty" example:"1"`
			InferenceFrame RerankInferenceFrame `json:"inference_frame"`
		}(*r)

		return json.Marshal(req)
	default:
		return nil, fmt.Errorf("invalid inference frame: %s", r.InferenceFrame)
	}
}

func (r *RerankRequest) UnmarshalJSON(data []byte) error {
	var err error

	switch r.InferenceFrame {
	case RerankInferenceFrameTEI:
		err = json.Unmarshal(data, r)
		return err
	case RerankInferenceFrameInfinity:
		req := struct {
			Model          string               `json:"model"`
			Query          string               `json:"query"`
			Texts          []string             `json:"documents,omitempty"`
			RawScores      bool                 `json:"raw_scores,omitempty" example:"false"`
			ReturnText     bool                 `json:"return_document,omitempty" example:"false"`
			TopN           int                  `json:"top_n,omitempty" example:"1"`
			InferenceFrame RerankInferenceFrame `json:"inference_frame"`
		}(*r)

		return json.Unmarshal(data, &req)
	default:
		return fmt.Errorf("invalid inference frame: %s", r.InferenceFrame)
	}
}

type ReranKResponse struct {
	Object string   `json:"object"`
	Data   []Rerank `json:"data"`
	Model  string   `json:"model"`
	Usage  *Usage   `json:"usage"`
}

type Rerank struct {
	Index          int      `json:"index"`
	Score          *float32 `json:"score,omitempty"`
	Object         string   `json:"object,omitempty"`
	RelevanceScore *float32 `json:"relevance_score,omitempty"`
	Document       string   `json:"document,omitempty"`
}

type TEIOriginalReranKResponse []Rerank

func (e TEIOriginalReranKResponse) ToReranKResponse(model string) *ReranKResponse {
	var data []Rerank

	for _, rerank := range e {
		data = append(data, Rerank{
			Object: "rerank",
			Index:  rerank.Index,
			Score:  rerank.Score,
		})
	}

	return &ReranKResponse{
		Object: "list",
		Data:   data,
		Model:  model,
	}
}

type InfinityOriginalReranKResponse struct {
	Object  string   `json:"object"`
	Results []Rerank `json:"results"`
	Model   string   `json:"model"`
	Usage   *Usage   `json:"usage"`
	ID      string   `json:"id"`
	Created int      `json:"created"`
}

func (e InfinityOriginalReranKResponse) ToReranKResponse(model string) *ReranKResponse {
	var data []Rerank
	for _, rerank := range e.Results {
		fmt.Println(rerank)
		data = append(data, Rerank{
			Object: e.Object,
			Index:  rerank.Index,
			Score:  rerank.RelevanceScore,
		})
	}

	return &ReranKResponse{
		Object: "list",
		Data:   data,
		Model:  model,
		Usage:  e.Usage,
	}
}

func (c *Client) switchRerankResponse(ctx context.Context, req *http.Request, frame RerankInferenceFrame, model string) (*ReranKResponse, error) {
	var err error
	switch frame {
	case RerankInferenceFrameTEI:
		var originalReranKResponse TEIOriginalReranKResponse
		err = c.sendRequest(ctx, req, &originalReranKResponse)
		if err != nil {
			return nil, err
		}
		return originalReranKResponse.ToReranKResponse(model), nil
	case RerankInferenceFrameInfinity:
		var originalReranKResponse InfinityOriginalReranKResponse
		err = c.sendRequest(ctx, req, &originalReranKResponse)
		if err != nil {
			return nil, err
		}
		return originalReranKResponse.ToReranKResponse(model), nil
	default:
		return nil, fmt.Errorf("invalid inference frame: %s", frame)
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

	return c.switchRerankResponse(ctx, req, request.InferenceFrame, request.Model)
}
