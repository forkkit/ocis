package svc

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/owncloud/ocis-webdav/pkg/config"
)

// Service defines the extension handlers.
type Service interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
	Dummy(http.ResponseWriter, *http.Request)
}

// NewService returns a service implementation for Service.
func NewService(opts ...Option) Service {
	options := newOptions(opts...)

	m := chi.NewMux()
	m.Use(options.Middleware...)

	svc := Webdav{
		config: options.Config,
		mux:    m,
	}

	m.Route(options.Config.HTTP.Root, func(r chi.Router) {
		r.Get("/", svc.Dummy)
	})

	return svc
}

// Webdav defines implements the business logic for Service.
type Webdav struct {
	config *config.Config
	mux    *chi.Mux
}

// ServeHTTP implements the Service interface.
func (g Webdav) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.mux.ServeHTTP(w, r)
}

// Dummy implements the Service interface.
func (g Webdav) Dummy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(http.StatusText(http.StatusOK)))
}
