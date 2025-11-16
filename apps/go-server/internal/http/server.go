package http

import (
	"context"
	"database/sql"
	"axon/server/internal/config"
	"axon/server/internal/routes"
	"log"

	"github.com/gofiber/fiber/v3"
)

type Server struct {
	app  *fiber.App
	addr string
	cfg  *config.Config
	_    *sql.DB // TODO: include db
}

func New(addr string, cfg *config.Config) (*Server, error) {
	app := fiber.New()
	server := &Server{
		app:  app,
		addr: addr,
		cfg:  cfg,
	}
	routes.Register(app, cfg)
	return server, nil
}

func (server *Server) Start() error {
	log.Printf("Server listening at %s\n", server.addr)
	return server.app.Listen(server.addr)
}

func (server *Server) Shutdown(ctx context.Context) error {
	return server.app.Shutdown()
}
