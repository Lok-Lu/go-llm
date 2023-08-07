package general

type ChatParams struct {
	MaxNewTokens      int      `json:"max_new_tokens"`
	Temperature       *float64 `json:"temperature"`
	TopK              *int     `json:"top_k"`
	TopP              *float64 `json:"top_p"`
	NumBeans          *int     `json:"num_beans"`
	RandomSeed        *int     `json:"random_seed"`
	DoSample          bool     `json:"do_sample"`
	RepetitionPenalty *float64 `json:"repetition_penalty"`
}

type ChatRequest struct {
	Inputs     string     `json:"inputs"`
	Parameters ChatParams `json:"parameters"`
}

type ChatResponse struct {
	GeneratedText string `json:"generated_text"`
}

// stream

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
