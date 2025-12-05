package app

import (
	"net/http"
	"subscription/internal/config"
	"time"
)

type Server struct {
	srv *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		srv: &http.Server{
			Addr:              cfg.Server.Host + ":" + cfg.Server.Port,
			Handler:           handler,
			ReadHeaderTimeout: time.Second * 10,
		},
	}
}

func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}
