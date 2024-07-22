package handler

import (
	"encoding/json"
	"go-cassandra-crud/entity"
	"go-cassandra-crud/usecase"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	usecase usecase.Usecase
}

func New(uc usecase.Usecase) *Handler {
	return &Handler{
		usecase: uc,
	}
}

func (h *Handler) FetchAll(w http.ResponseWriter, r *http.Request) {
	cartCounts, err := h.usecase.FetchAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	data, err := json.MarshalIndent(cartCounts, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) FetchOne(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	cartCount, err := h.usecase.FetchOne(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	data, err := json.MarshalIndent(cartCount, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) Insert(w http.ResponseWriter, r *http.Request) {
	var req entity.CartCount
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if err := h.usecase.Insert(req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	if err := h.usecase.Delete(userID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
