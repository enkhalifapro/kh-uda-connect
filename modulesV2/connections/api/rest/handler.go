package rest

import (
	"encoding/json"
	"enkhalifapro/connections/internal"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type Handler struct {
	service Service
}

type Service interface {
	GetAll() ([]internal.Person, error)
	GetByID(id int) (*internal.Person, error)
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

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	persons, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	idParam := p.ByName("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	persons, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req internal.CreatePayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
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
