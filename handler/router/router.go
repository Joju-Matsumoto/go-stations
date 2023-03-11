package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()

	// healthz
	mux.Handle("/healthz", &handler.HealthzHandler{})

	// todo
	mux.Handle("/todos", handler.NewTODOHandler(service.NewTODOService(todoDB)))
	return mux
}
