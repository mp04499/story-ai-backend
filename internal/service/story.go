package service

import (
	"encoding/json"
	"github.com/mp04499/story-ai-backend/internal/client"
	"github.com/mp04499/story-ai-backend/internal/model"
	"net/http"
)

const SystemMessage = "You are a imaginative story telling chat bot"

func GetStory(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var request model.StoryRequest
	err := decoder.Decode(&request)

	if err != nil {
		panic(err)
	}

	llmRequest := createLlmRequest(request.Text)
	llmResponse := client.CallStoryLlm(llmRequest)
	encoder := json.NewEncoder(rw)
	err = encoder.Encode(llmResponse)
	if err != nil {
		panic(err)
	}
}

func createLlmRequest(text string) *model.LlmRequest {
	request := model.NewLlmRequest()
	var messages = []model.Message{
		{Role: "system", Content: SystemMessage},
		{Role: "user", Content: text},
	}
	request.Messages = messages

	return request
}
