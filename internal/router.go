package internal

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mp04499/story-ai-backend/internal/service"
)

func InitRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/story-llm", func(r chi.Router) {
		r.Post("/", service.GetStory)
	})

	return r
}
