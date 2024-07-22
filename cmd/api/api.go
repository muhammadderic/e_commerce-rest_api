package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muhammadderic/ecomrest/services/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Create a new user store and handler, the register user router to the subrouter
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	return http.ListenAndServe(s.addr, router)
}
