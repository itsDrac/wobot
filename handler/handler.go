package handler

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/itsDrac/wobot/service"
)

func NewChiRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	return r
}

type ChiHandler struct {
	Service *service.Service
}

func (h ChiHandler) Mount(router http.Handler) {
	r, ok := router.(chi.Router)
	if !ok {
		log.Fatal("Handler is not a chi.Router")
		return
	}
	r.Get("/api/v1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
}

func (h ChiHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (h ChiHandler) Run(r http.Handler) {
	// Start the server
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
