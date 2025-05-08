package router

import (
	"github.com/go-chi/chi/v5"
	
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Route("sis-pustaka", func(r chi.Router) {
		r.Get("/", handler.GetAllBooks)
		r.Post("/", handler.CreateBook)
		r.Get("/{id}", handler.GetBookByID)
		r.Put("/{id}", handler.UpdateBook)
		r.Delete("/{id}", handler.DeleteBook)
	})
	return r
}