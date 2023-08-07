package general

type ChatParams struct {
	MaxNewTokens      int      `json:"max_new_tokens"`
	Temperature       *float64 `json:"temperature"`
	TopK              *int     `json:"top_k"`
	TopP              *float64 `json:"top_p"`
	NumBeans          *int     `json:"num_beans"`
	RandomSeed        *int     `json:"random_seed"`
	RepetitionPenalty *float64 `json:"repetition_penalty"`
}

type ChatRequest struct {
	Inputs     string     `json:"inputs"`
	Parameters ChatParams `json:"parameters"`
}
