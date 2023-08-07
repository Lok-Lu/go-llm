package internal

import "encoding/json"

type JsonSerializer interface {
	Marshal(value any) ([]byte, error)
	Unmarshal(data []byte, v any) error
}

type jsonSerializer struct{}

func NewJsonSerializer() JsonSerializer {
	return &jsonSerializer{}
}

func (jm *jsonSerializer) Marshal(value any) ([]byte, error) {
	return json.Marshal(value)
}

func (jm *jsonSerializer) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
