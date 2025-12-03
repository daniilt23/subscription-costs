package app

import (
	"log"
	"net/http"
	"subscription/internal/config"
)

type Server struct {
	srv *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	log.Printf("Server run on port: %s", cfg.Server.Port)
	return &Server{
		srv: &http.Server{
			Addr: cfg.Server.Host + ":" + cfg.Server.Port,
			Handler: handler,
		},
	}
}

func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}