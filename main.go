package main

import (
	"github.com/mp04499/story-ai-backend/internal"
	"net/http"
)

func main() {
	r := internal.InitRouter()

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
