package handlers

import (
	"github.com/go-chi/chi"
	chiMiddle "github.com/go-chi/chi/middleware"
	"github.com/jubaer-ad/goapi/internal/middleware"
)

func Handler(cm *chi.Mux) {
	cm.Use(chiMiddle.StripSlashes)
	cm.Route("/account", func(r chi.Router) {
		r.Use(middleware.Authorization)
		r.Get("/cards", GetCardCount)
	})
}
