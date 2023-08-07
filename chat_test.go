package go_llm

import (
	"context"
	"github.com/Lok-Lu/go-llm/general"
	"testing"
)

func TestChat(t *testing.T) {
	url := ""
	client := NewClient().WithGeneralClient(url)
	var (
		a   float64 = 1.0
		b   float64 = 0.6
		c   float64 = 1.3
		d   int     = 2
		e   int     = 1
		f   int     = 0
		req         = general.ChatRequest{
			Inputs: "",
			Parameters: general.ChatParams{
				MaxNewTokens:      1024,
				Temperature:       &a,
				TopK:              &d,
				TopP:              &b,
				NumBeans:          &e,
				RandomSeed:        &f,
				RepetitionPenalty: &c,
			},
		}
	)
	chat, err := client.Llm.SetUrl("").CreateChatCompletionStream(context.Background(), &req)
	if err != nil {
		t.Error(err)
	}
	for {
		r, err := chat.Recv()
		if err != nil {
			break
		}
		t.Log(r)
	}
	//t.Log(chat)
}
