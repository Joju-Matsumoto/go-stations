package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/handler/middleware"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()

	// healthz
	mux.Handle("/healthz", handler.NewHealthzHandler())

	// todo
	mux.Handle("/todos", handler.NewTODOHandler(service.NewTODOService(todoDB)))

	recoveryMiddleware := middleware.NewRecovery()
	// do-panic
	mux.Handle("/do-panic", recoveryMiddleware.Recovery(handler.NewDoPanicHandler()))
	return mux
}
