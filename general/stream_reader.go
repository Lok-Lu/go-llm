package general

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"

	wraperr "github.com/Lok-Lu/go-llm/error"
	"github.com/Lok-Lu/go-llm/internal"
)

var (
	ErrTooManyEmptyStreamMessages = errors.New("stream has sent too many empty messages")
	headerData                    = []byte("data:")
	errorPrefix                   = []byte(`data:{"error":`)
)

type StreamType interface {
	ChatStreamResponse
}

type StreamReader[T StreamType] struct {
	EmptyMessagesLimit uint
	isFinished         bool

	reader         *bufio.Reader
	response       *http.Response
	errAccumulator internal.ErrorAccumulator
	jsonSerializer internal.JsonSerializer
}

func NewStreamReader[T StreamType](emptyMessagesLimit uint, reader *bufio.Reader, resp *http.Response) *StreamReader[T] {
	return &StreamReader[T]{
		EmptyMessagesLimit: emptyMessagesLimit,
		isFinished:         false,
		reader:             reader,
		response:           resp,
		errAccumulator:     internal.NewErrorAccumulator(),
		jsonSerializer:     internal.NewJsonSerializer(),
	}
}

func (stream *StreamReader[T]) Recv() (response T, err error) {
	if stream.isFinished {
		err = io.EOF
		return
	}

	response, err = stream.processLines()
	return
}

func (stream *StreamReader[T]) processLines() (T, error) {
	var (
		emptyMessagesCount uint
		hasErrorPrefix     bool
	)
	for {
		rawLine, readErr := stream.reader.ReadBytes('\n')
		if readErr != nil {
			if errors.Is(readErr, io.EOF) {
				stream.isFinished = true
				return *new(T), io.EOF
			}

			respErr := stream.unmarshalError()
			if respErr != nil {
				return *new(T), fmt.Errorf("error, %w", respErr.Error)
			}
			return *new(T), readErr
		}

		noSpaceLine := bytes.TrimSpace(rawLine)
		if bytes.HasPrefix(noSpaceLine, errorPrefix) {
			hasErrorPrefix = true
		}

		if !bytes.HasPrefix(noSpaceLine, headerData) {
			if hasErrorPrefix {
				noSpaceLine = bytes.TrimPrefix(noSpaceLine, headerData)
			}

			writeErr := stream.errAccumulator.Write(noSpaceLine)
			if writeErr != nil {
				return *new(T), writeErr
			}
			emptyMessagesCount++
			if emptyMessagesCount > stream.EmptyMessagesLimit {
				return *new(T), ErrTooManyEmptyStreamMessages
			}

			respErr := stream.unmarshalError()
			if respErr != nil {
				return *new(T), fmt.Errorf("error, %w", respErr.Error)
			}
			continue
		}

		noPrefixLine := bytes.TrimSpace(bytes.TrimPrefix(noSpaceLine, headerData))

		var response T
		unmarshalErr := stream.jsonSerializer.Unmarshal(noPrefixLine, &response)
		if unmarshalErr != nil {
			return *new(T), unmarshalErr
		}

		return response, nil
	}
}

func (stream *StreamReader[T]) unmarshalError() (errResp *wraperr.ErrorResponse) {
	errBytes := stream.errAccumulator.Bytes()
	if len(errBytes) == 0 {
		return
	}

	err := stream.jsonSerializer.Unmarshal(errBytes, &errResp)
	if err != nil {
		errResp = nil
	}

	return
}

func (stream *StreamReader[T]) Close() {
	stream.response.Body.Close()
}
