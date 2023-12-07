package error

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	ServiceUnavailable = &ErrorResponse{
		Error: &APIError{
			Code:           http.StatusServiceUnavailable,
			Message:        "service unavailable",
			HTTPStatusCode: http.StatusServiceUnavailable,
		}}
	ErrServiceUnavailable = fmt.Errorf("error, status code: %d, message: %w", ServiceUnavailable.Error.HTTPStatusCode, ServiceUnavailable.Error)
)

// APIError provides error information returned by the bard API.
type APIError struct {
	Code           any    `json:"code,omitempty"`
	Message        string `json:"message"`
	HTTPStatusCode int    `json:"-"`
}

// RequestError provides informations about generic request errors.
type RequestError struct {
	HTTPStatusCode int
	Err            error
}

type ErrorResponse struct {
	Error *APIError `json:"error,omitempty"`
}

func (e *APIError) Error() string {
	return e.Message
}

func (e *APIError) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &e.Message)
}

func (e *RequestError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return fmt.Sprintf("status code %d", e.HTTPStatusCode)
}

func (e *RequestError) Unwrap() error {
	return e.Err
}
