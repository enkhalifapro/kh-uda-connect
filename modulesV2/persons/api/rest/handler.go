package rest

import (
	"encoding/json"
	"enkhalifapro/persons/internal"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	service Service
}

type Service interface {
	GetAll() ([]internal.Person, error)
	GetByID(id int) (*internal.Person, error)
	GetConnectionsByPersonID(personID int, startDate string, endDate string, distance float64) ([]internal.Connection, error)
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

func (h *Handler) GetLocationsByPersonID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	personID := p.ByName("personId")
	id, err := strconv.Atoi(personID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	startTime := p.ByName("startTime")
	endTime := p.ByName("endTime")
	distanceParam := p.ByName("distance")
	distance, err := strconv.ParseFloat(distanceParam, 64)
	if err != nil {
		log.Fatalf("Error converting string to float64: %v", err)
	}

	persons, err := h.service.GetConnectionsByPersonID(id, startTime, endTime, distance)
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
