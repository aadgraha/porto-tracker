package transaction

import (
	db "github.com/aadgraha/porto-tracker/internal/repository/postgres"
	"github.com/go-chi/chi/v5"
)

func Routes(q *db.Queries) chi.Router {
	r := chi.NewRouter()
	h := NewHandler(q)

	r.Post("/", h.Create)
	r.Get("/", h.List)
	r.Get("/{id}", h.GetByID)
	r.Post("/{id}/correct", h.Correct)

	return r
}
