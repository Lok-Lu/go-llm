package request

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type marshaller interface {
	marshal(value any) ([]byte, error)
}

type jsonMarshaller struct{}

func (jm *jsonMarshaller) marshal(value any) ([]byte, error) {
	return json.Marshal(value)
}

type RequestBuilder interface {
	Build(ctx context.Context, method, url string, request any) (*http.Request, error)
}

type httpRequestBuilder struct {
	marshaller marshaller
}

func NewRequestBuilder() RequestBuilder {
	return &httpRequestBuilder{
		marshaller: &jsonMarshaller{},
	}
}

func (b *httpRequestBuilder) Build(ctx context.Context, method, url string, request any) (*http.Request, error) {
	if request == nil {
		return http.NewRequestWithContext(ctx, method, url, nil)
	}

	var reqBytes []byte
	reqBytes, err := b.marshaller.marshal(request)
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
