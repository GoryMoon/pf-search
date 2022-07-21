package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/meilisearch/meilisearch-go"
)

type Server struct {
	port      int
	host      string
	meiliHost string
	meiliKey  string
	meili     *meilisearch.Client
	fiber     *fiber.App
}

func (s *Server) Run() error {
	s.setupRoutes()
	s.initMelli()

	err := s.fiber.Listen(fmt.Sprintf("%s:%d", s.host, s.port))
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Shutdown() {
	_ = s.fiber.Shutdown()
}

func CreateNewServer(host string, port int, melliHost string, melliKey string) *Server {
	return &Server{
		host:      host,
		port:      port,
		meiliHost: melliHost,
		meiliKey:  melliKey,
		fiber:     setupFiber(),
		meili:     setupMelli(melliHost, melliKey),
	}
}
