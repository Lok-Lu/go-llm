package internal

import (
	"bytes"
	"encoding/json"
)

type JsonSerializer interface {
	Marshal(value any) ([]byte, error)
	Unmarshal(data []byte, v any) error
}

type jsonSerializer struct{}

func NewJsonSerializer() JsonSerializer {
	return &jsonSerializer{}
}

func (jm *jsonSerializer) Marshal(value any) ([]byte, error) {
	// for adapter json decode html code
	byteBuf := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(byteBuf)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(value)
	return byteBuf.Bytes(), err
	// return json.Marshal(value)
}

func (jm *jsonSerializer) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
