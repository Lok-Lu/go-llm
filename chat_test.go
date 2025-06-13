package go_llm

import (
	"context"
	"testing"

	"github.com/Lok-Lu/go-llm/general"
)

func TestChat(t *testing.T) {
	url := ""
	client := NewClient().WithGeneralClient(url, "")
	var (
		a float64 = 1.0
		// b   float64 = 0.6
		// c   float64 = 1.3
		// d   int     = 1024
		req = general.ChatRequest{
			Inputs: "convert to JSON: I saw a puppy a cat and a raccoon during my bike ride in the park, use this json schema",
			Parameters: general.ChatParams{
				MaxNewTokens: 1024,
				Temperature:  &a,
				// TopK:              &d,
				// TopP:              &b,
				// RepetitionPenalty: &c,
				// Truncate:          &d,
				Grammar: &general.GrammarParams{
					Type: "json",
					Value: general.SchemaParams{
						Properties: map[string]general.Definition{
							"location": {"title": "Location", "type": "string"},
							"activity": {"title": "Activity", "type": "string"},
							"animals_seen": {
								"maximum": 5,
								"minimum": 1,
								"title":   "Animals Seen",
								"type":    "integer",
							},
							"animals": {"title": "Animals", "type": "array", "items": map[string]any{"type": "string"}},
						},
						Required: []string{"location", "activity", "animals_seen", "animals"},
						Title:    "Animals",
						Type:     "object",
					},
				},
			},
		}
	)
	chat, err := client.Llm.CreateChat(context.Background(), url, &req)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", chat)
	// for {
	// 	r, err := chat.Recv()
	// 	if err != nil {
	// 		break
	// 	}
	// 	t.Log(r)
	// }
	// metric, err := client.Llm.GetMetrics(context.Background(), url, time.Second*3)
	// if err != nil {
	// 	t.Error(err)
	// }
	// a := bytes.Split(metric, []byte("\n"))
	// t.Log(string(a[0]))

}
