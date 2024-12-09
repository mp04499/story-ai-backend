package model

const MaxNewTokens = 2000
const TopK = 50
const TopP = 0.95
const NumReturnSequences = 1
const DoSample = true

type LlmRequest struct {
	Model              string    `json:"model"`
	Messages           []Message `json:"messages"`
	MaxNewToken        int       `json:"max_new_tokens"`
	TopK               int       `json:"top_k"`
	TopP               float32   `json:"top_p"`
	DoSample           bool      `json:"do_sample"`
	NumReturnSequences int       `json:"num_return_sequences"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type LlmResponse struct {
	Object            string   `json:"object"`
	Id                string   `json:"id"`
	Created           uint64   `json:"created"`
	Model             string   `json:"model"`
	SystemFingerprint string   `json:"system_fingerprint"`
	Choices           []Choice `json:"choices"`
	Usage             Usage    `json:"usage"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type StoryResponse struct {
	Data *LlmResponse `json:"data"`
}

type StoryRequest struct {
	Text string `json:"text"`
}

func NewLlmRequest() *LlmRequest {
	request := &LlmRequest{
		Model:              "GawdSB/story_model",
		MaxNewToken:        MaxNewTokens,
		TopK:               TopK,
		TopP:               TopP,
		DoSample:           DoSample,
		NumReturnSequences: NumReturnSequences,
	}

	return request
}
