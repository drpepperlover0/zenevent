package server

import (
	"github.com/drpepperlover0/internal/routes"
	"github.com/drpepperlover0/storage"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() error {

	router := routes.InitRoutes()
	if err := storage.CreateDB(); err != nil {
		return err
	}

	return router.Start(":8080")
}
