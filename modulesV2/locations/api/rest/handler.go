package rest

import (
	"encoding/json"
	"enkhalifapro/locations/internal"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

type Handler struct {
	service Service
}

type Service interface {
	Add(person *internal.CreatePayload) error
}

type HealthResponse struct {
	Message string `json:"message"`
}

func NewHandler(srv Service) *Handler {
	return &Handler{
		service: srv,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req internal.CreatePayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	req.CreationTS = time.Now().Unix()
	if err := h.service.Add(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 204 No Content
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response := HealthResponse{Message: "Ok!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
