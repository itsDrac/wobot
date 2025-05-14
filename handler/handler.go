package handler

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/itsDrac/wobot/service"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

var validate *validator.Validate

func NewChiRouter() chi.Router {
	// Initialize the validator
	validate = validator.New()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // The url pointing to API definition
	))
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
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/health", h.Health)
			r.Post("/users/create", h.CreateUser)
			r.Post("/users/login", h.LoginUser)
			// Authenticate routes.
			r.Group(func(r chi.Router) {
				r.Use(h.AuthUser)
				r.Post("/upload", h.UploadFile)
				r.Get("/storage/remaining", h.GetStorage)
				r.Get("/files", h.GetFiles)
			})
		})
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
