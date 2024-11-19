package model

type LlmRequest struct {
	Model string `json:"model"`
	Messages []Message `json:"messages"`
	MaxNewToken int `json:"max_new_tokens"`
	TopK int `json:"top_k"`
	TopP int `json:"top_p"`
	DoSample bool `json:"do_sample"`
}

type Message struct {
	Role string `json:"role"`
	Content string `json:"content"`
}

func NewLlmRequest() LlmRequest {
	request := LlmRequest{
		Model: "GawdSB/story_model",
	}

	return request
}