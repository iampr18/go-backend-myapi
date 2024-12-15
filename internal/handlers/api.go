package handlers

import (
	"myapimodule/internal/middleware"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

//fuction with capital letter tells that function can be impoerted in another package

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)
	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinsBalance)
	})
}
