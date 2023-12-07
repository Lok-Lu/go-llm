package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	wraperr "github.com/Lok-Lu/go-llm/error"
)

type RequestBuilder interface {
	Build(ctx context.Context, method, url string, request any) (*http.Request, error)
	Send(ctx context.Context, req *http.Request, v any) error
	SendNoCloseWithCustomClient(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error)
	SendNoClose(ctx context.Context, req *http.Request) (*http.Response, error)
}

type httpRequestBuilder struct {
	jsonSerializer JsonSerializer
	client         *http.Client
}

func NewRequestBuilder(client *http.Client) RequestBuilder {
	return &httpRequestBuilder{
		jsonSerializer: NewJsonSerializer(),
		client:         client,
	}
}

func (b *httpRequestBuilder) Build(ctx context.Context, method, url string, request any) (*http.Request, error) {
	if request == nil {
		return http.NewRequestWithContext(ctx, method, url, nil)
	}

	var reqBytes []byte
	reqBytes, err := b.jsonSerializer.Marshal(request)
	if err != nil {
		return nil, err
	}
	return http.NewRequestWithContext(
		ctx,
		method,
		url,
		bytes.NewBuffer(reqBytes),
	)
}

func (b *httpRequestBuilder) SendNoCloseWithCustomClient(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
	res, err := client.Do(req)
	if err != nil {
		return nil, handleErrorWithOutResp(err)
	}

	if isFailureStatusCode(res) {
		return nil, handleErrorResp(res)
	}
	return res, nil
}

func (b *httpRequestBuilder) SendNoClose(ctx context.Context, req *http.Request) (*http.Response, error) {
	res, err := b.client.Do(req)
	if err != nil {
		return nil, handleErrorWithOutResp(err)
	}

	if isFailureStatusCode(res) {
		return nil, handleErrorResp(res)
	}
	return res, nil
}

func (b *httpRequestBuilder) Send(ctx context.Context, req *http.Request, v any) error {
	res, err := b.client.Do(req)
	if err != nil {
		return handleErrorWithOutResp(err)
	}
	defer res.Body.Close()

	if isFailureStatusCode(res) {
		return handleErrorResp(res)
	}
	return decodeResponse(res.Body, v)
}

func handleErrorResp(resp *http.Response) error {
	var errRes wraperr.ErrorResponse
	errByte, _ := io.ReadAll(resp.Body)
	errRes.Error = &wraperr.APIError{
		Code:           resp.StatusCode,
		Message:        string(errByte),
		HTTPStatusCode: resp.StatusCode,
	}
	return fmt.Errorf("error, status code: %d, message: %w", resp.StatusCode, errRes.Error)
}

func handleErrorWithOutResp(err error) error {
	var errRes = wraperr.ServiceUnavailableError
	errRes.Error = &wraperr.APIError{
		Code:           http.StatusServiceUnavailable,
		Message:        err.Error(),
		HTTPStatusCode: http.StatusServiceUnavailable,
	}
	return fmt.Errorf("error, status code: %d, message: %w", http.StatusServiceUnavailable, errRes.Error)
}


func isFailureStatusCode(resp *http.Response) bool {
	return resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest
}

func decodeResponse(body io.Reader, v any) error {
	if v == nil {
		return nil
	}

	if result, ok := v.(*string); ok {
		return decodeString(body, result)
	}
	return json.NewDecoder(body).Decode(v)
}

func decodeString(body io.Reader, output *string) error {
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	*output = string(b)
	return nil
}
