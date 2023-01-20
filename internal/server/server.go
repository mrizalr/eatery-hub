package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mrizalr/eatery-hub/config"
	"gorm.io/gorm"
)

type Server struct {
	app    *fiber.App
	config *config.Config
	db     *gorm.DB
}

func New(config *config.Config, db *gorm.DB) *Server {
	return &Server{
		app:    fiber.New(),
		config: config,
		db:     db,
	}
}

func (s *Server) Run() error {
	s.MapHandlers(s.app)
	log.Printf("Server is listening on PORT: %v", s.config.Server.Port)
	return s.app.Listen(fmt.Sprintf(":%v", s.config.Server.Port))
}
