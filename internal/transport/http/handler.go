package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router *chi.Mux
}

// NewHander - returns a pointer to a Handler
func NewHandler() *Handler {
	return &Handler{}

}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = chi.NewRouter()
	h.Router.HandleFunc("/api/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive!")
	})
}
