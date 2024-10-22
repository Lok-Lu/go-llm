package go_llm

import (
	"github.com/Lok-Lu/go-llm/general"
	"github.com/Lok-Lu/go-llm/pharm_lc"
)

type LLMClient struct {
	Pharm *pharm_lc.Client
	Llm   *general.Client
}

func NewClient() *LLMClient {
	return new(LLMClient)
}

func (c *LLMClient) WithPharmClient(url string) *LLMClient {
	c.Pharm = pharm_lc.NewClient(url)
	return c
}

func (c *LLMClient) WithPharmClientWithVersions(version map[string]string) *LLMClient {
	c.Pharm = pharm_lc.NewClientWithVersions(version)
	return c
}

func (c *LLMClient) WithGeneralClient(url, token string) *LLMClient {
	c.Llm = general.NewClient(url, token)
	return c
}
