package go_llm

import (
	"bytes"
	"context"
	"testing"
)

func TestChat(t *testing.T) {
	url := ""
	client := NewClient().WithGeneralClient(url)
	// var (
	// 	a   float64 = 1.0
	// 	b   float64 = 0.6
	// 	c   float64 = 1.3
	// 	d   int     = 1024
	// 	req         = general.ChatRequest{
	// 		Inputs: "你是谁[AI]",
	// 		Parameters: general.ChatParams{
	// 			MaxNewTokens:      1024,
	// 			Temperature:       &a,
	// 			TopK:              &d,
	// 			TopP:              &b,
	// 			RepetitionPenalty: &c,
	// 			Truncate:          &d,
	// 		},
	// 	}
	// )
	// chat, err := client.Llm.CreateChat(context.Background(), url, &req)
	// if err != nil {
	// 	t.Error(err)
	// }
	//for {
	//	r, err := chat.Recv()
	//	if err != nil {
	//		break
	//	}
	//	t.Log(r)
	//}
	metric, err := client.Llm.GetMetrics(context.Background(), url)
	if err != nil {
		t.Error(err)
	}
	a := bytes.Split(metric, []byte("\n"))
	t.Log(string(a[0]))

}
