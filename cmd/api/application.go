package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/weebNeedWeed/leetcode-ai-explainer/internal/storage"
)

type application struct {
	storage storage.Storage
}

func (a *application) mount() *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"*"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/problems/{id}/explanation", a.getExplanationHandler)

	return r
}

func (*application) run(mux *chi.Mux) error {
	srv := http.Server{
		Addr:    ":9090",
		Handler: mux,
	}

	fmt.Printf("Application is listening on: %s\n", srv.Addr)

	return srv.ListenAndServe()
}
