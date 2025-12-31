package http

import (
	"axon/server/internal/config"
	"axon/server/internal/routes"
	"context"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	app  *fiber.App
	addr string
	cfg  *config.Config
	db   *pgxpool.Pool
}

func New(addr string, cfg *config.Config, pool *pgxpool.Pool) (*Server, error) {
	app := fiber.New()
	server := &Server{
		app:  app,
		addr: addr,
		cfg:  cfg,
		db:   pool,
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
