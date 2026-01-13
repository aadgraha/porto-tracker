package asset

import (
	"net/http"

	db "github.com/aadgraha/porto-tracker/internal/repository/postgres"
	"github.com/go-chi/chi/v5"
)

func Routes(q *db.Queries) http.Handler {
	h := NewHandler(q)

	r := chi.NewRouter()
	r.Post("/", h.Create)
	r.Get("/", h.List)
	r.Get("/{id}", h.GetByID)
	r.Put("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)

	return r
}
