package go_llm

import "github.com/patsnapops/go-llm/pharm_lc"

type LLMClient struct {
	Pharm *pharm_lc.Client
}

func NewClient(config *LLMConfig) *LLMClient {
	return &LLMClient{
		Pharm: pharm_lc.NewClientWithConfig(config.Pharm),
	}
}
