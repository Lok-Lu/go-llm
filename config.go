package go_llm

import "github.com/patsnapops/go-llm/pharm_lc"

type LLMConfig struct {
	Pharm pharm_lc.ClientConfig
}

func NewLLMConfig() *LLMConfig {
	return new(LLMConfig)
}

func (l *LLMConfig) SetPharmConfig(Url string) *LLMConfig {
	l.Pharm = pharm_lc.DefaultConfig(Url)
	return l
}

func (l *LLMConfig) SetPharmConfigWithVersion(versions map[string]string) *LLMConfig {
	l.Pharm = pharm_lc.NewConfigWithVersion(versions)
	return l
}
