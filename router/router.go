package router

import (
	"sis-pustaka/handler" //import handler yang kita buat sebelumnya
	"sis-pustaka/middleware" //import middleware yang kita buat sebelumnya
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware" //import chi untuk routing
	
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	
	r.Use(chiMiddleware.Logger) //gunakan middleware logger untuk mencatat request yang masuk
	r.Use(middleware.LoggerMiddleware) //gunakan middleware logger untuk mencatat request yang masuk

	r.Route("/sis-pustaka", func(r chi.Router) {
		r.Get("/", book_handler.GetAllBooks)
		r.Post("/", book_handler.CreateBook)
		r.Get("/{id}", book_handler.GetBookByID)
		r.Put("/{id}", book_handler.UpdateBook)
		r.Delete("/{id}", book_handler.DeleteBook)
	})
	return r
}