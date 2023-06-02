package go_llm

import (
	"context"
	"github.com/patsnapops/go-llm/pharm_lc"
	"testing"
)

func TestChat(t *testing.T) {
	url := ""
	config := NewLLMConfig().SetPharmConfig(url)
	client := NewClient(config)

	var req = pharm_lc.ChatRequest{
		Message: "你是gpt几",
	}
	chat, err := client.Pharm.CreateChat(context.Background(), &req)
	if err != nil {
		t.Error(err)
	}
	t.Log(chat)
}
