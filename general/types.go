package general

type DataType string

const (
	Object  DataType = "object"
	Number  DataType = "number"
	Integer DataType = "integer"
	String  DataType = "string"
	Array   DataType = "array"
	Null    DataType = "null"
	Boolean DataType = "boolean"
)

type GrammarType string

const (
	Json       GrammarType = "json"
	JsonSchema GrammarType = "json_schema"
	Regex      GrammarType = "regex"
)

type GrammarParams struct {
	Type  GrammarType  `json:"type"`
	Value SchemaParams `json:"value"`
}

type Definition map[string]any

type SchemaParams struct {
	Properties map[string]Definition `json:"properties,omitempty"`
	Required   []string              `json:"required,omitempty"`
	Title      string                `json:"title,omitempty"`
	Type       DataType              `json:"type,omitempty"`
	Items      Definition            `json:"items,omitempty"`
}

type ChatParams struct {
	MaxNewTokens      int            `json:"max_new_tokens"`
	Temperature       *float64       `json:"temperature,omitempty"`
	TopK              *int           `json:"top_k,omitempty"`
	TopP              *float64       `json:"top_p,omitempty"`
	Seed              *int           `json:"seed,omitempty"`
	Stop              []string       `json:"stop,omitempty"`
	DoSample          *bool          `json:"do_sample,omitempty"`
	RepetitionPenalty *float64       `json:"repetition_penalty,omitempty"`
	BestOf            *int           `json:"best_of,omitempty"`
	Details           *bool          `json:"details,omitempty"`
	ReturnFullText    *bool          `json:"return_full_text,omitempty"`
	Truncate          *int           `json:"truncate,omitempty"`
	TypicalP          *int           `json:"typical_p,omitempty"`
	WaterMark         *bool          `json:"watermark,omitempty"`
	NumBeams          *int           `json:"num_beams,omitempty"`
	Grammar           *GrammarParams `json:"grammar,omitempty"`
}

type ChatRequest struct {
	Inputs     string     `json:"inputs"`
	Parameters ChatParams `json:"parameters"`
}

type ChatResponse struct {
	GeneratedText string `json:"generated_text"`
}

// strea
type ChatStreamToken struct {
	ID      uint    `json:"id"`
	Text    string  `json:"text"`
	Logprob float64 `json:"logprob"`
	Special bool    `json:"special"`
}

type ChatMessage struct {
	Content string `json:"content"`
}

type ChatChoice struct {
	ID      uint        `json:"id"`
	Logprob float64     `json:"logprob"`
	Special bool        `json:"special"`
	Delta   ChatMessage `json:"delta,omitempty"`
}

type WrapperChatStreamResponse struct {
	GeneratedText string       `json:"generated_text,omitempty"`
	Details       string       `json:"details"`
	Choices       []ChatChoice `json:"choices"`
}

type ChatStreamResponse struct {
	GeneratedText string          `json:"generated_text,omitempty"`
	Token         ChatStreamToken `json:"token"`
	Details       string          `json:"details"`
}

type ChatCompletionStream struct {
	*StreamReader[ChatStreamResponse]
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
