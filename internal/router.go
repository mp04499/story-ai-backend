package internal

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mp04499/story-ai-backend/internal/service"
	"net/http"
)

func InitRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/story-llm", func(r chi.Router) {
		r.Use(llmCtx)
		r.Post("/", service.GetStory)
	})

	return r
}

func llmCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		//w.Header().Set("Cache-Control", "no-cache")
		//w.Header().Set("Connection", "keep-alive")

		next.ServeHTTP(w, r)
	})
}
