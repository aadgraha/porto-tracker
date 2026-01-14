package transaction

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	db "github.com/aadgraha/porto-tracker/internal/repository/postgres"
	util "github.com/aadgraha/porto-tracker/utils"
)

type Handler struct {
	q *db.Queries
}

func NewHandler(q *db.Queries) *Handler {
	return &Handler{q: q}
}

/* ---------- CREATE ---------- */

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	qty, err := util.ParseNumeric(req.Quantity)
	if err != nil {
		http.Error(w, "invalid quantity", http.StatusBadRequest)
		return
	}

	price, err := util.ParseOptionalNumeric(req.Price)
	if err != nil {
		http.Error(w, "invalid price", http.StatusBadRequest)
		return
	}

	fee, err := util.ParseOptionalNumeric(req.Fee)
	if err != nil {
		http.Error(w, "invalid fee", http.StatusBadRequest)
		return
	}

	tx, err := h.q.CreateTransaction(r.Context(), db.CreateTransactionParams{
		AssetID:    req.AssetID,
		TxType:     req.TxType,
		Quantity:   qty,
		Price:      price,
		Fee:        fee,
		OccurredAt: util.ParseTimestamptz(req.OccurredAt),
	})

	if err != nil {
		log.Println("DB error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tx)
}

/* ---------- LIST ACTIVE ---------- */

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	txs, err := h.q.ListActiveTransactions(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(txs)
}

/* ---------- GET BY ID ---------- */

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	tx, err := h.q.GetTransactionByID(r.Context(), id)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tx)
}

/* ---------- CORRECT ---------- */

func (h *Handler) Correct(w http.ResponseWriter, r *http.Request) {
	originalID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var req CorrectTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	qty, err := util.ParseNumeric(req.Quantity)
	if err != nil {
		http.Error(w, "invalid quantity", http.StatusBadRequest)
		return
	}

	price, err := util.ParseOptionalNumeric(req.Price)
	if err != nil {
		http.Error(w, "invalid price", http.StatusBadRequest)
		return
	}

	fee, err := util.ParseOptionalNumeric(req.Fee)
	if err != nil {
		http.Error(w, "invalid fee", http.StatusBadRequest)
		return
	}

	tx, err := h.q.CorrectTransaction(r.Context(), db.CorrectTransactionParams{
		ID:         originalID,
		AssetID:    req.AssetID,
		TxType:     req.TxType,
		Quantity:   qty,
		Price:      price,
		Fee:        fee,
		OccurredAt: util.ParseTimestamptz(req.OccurredAt),
	})
	if err != nil {
		log.Println("DB error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tx)
}
