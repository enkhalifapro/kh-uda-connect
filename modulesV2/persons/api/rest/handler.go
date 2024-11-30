package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
}

type Response struct {
	Message string `json:"message"`
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hhhh222")
	response := Response{Message: "Ok!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
