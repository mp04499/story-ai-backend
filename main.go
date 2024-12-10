package main

import (
	"github.com/mp04499/story-ai-backend/internal"
	"net/http"
)

func main() {
	r := internal.InitRouter()

	println("Starting Story API on port 8080 🌬️")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
