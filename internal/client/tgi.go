package client

import (
	"bytes"
	"encoding/json"
	"github.com/mp04499/story-ai-backend/configs"
	"github.com/mp04499/story-ai-backend/internal/model"
	"io"
	"net/http"
	"time"
)

func CallStoryLlm(llmRequest *model.LlmRequest) *model.LlmResponse {
	config := configs.LoadConfig()

	req, err := json.Marshal(llmRequest)
	if err != nil {
		panic(err)
	}

	client := http.Client{Timeout: time.Second * 30}
	resp, err := client.Post(config.LlmUrl, "application/json", bytes.NewBuffer(req))
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	decoder := json.NewDecoder(resp.Body)

	var llmResponse model.LlmResponse
	err = decoder.Decode(&llmResponse)
	if err != nil {
		panic(err)
	}

	return &llmResponse
}
