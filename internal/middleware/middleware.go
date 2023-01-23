package middleware

import "github.com/mrizalr/eatery-hub/config"

type MiddlewareManager struct {
	config *config.Config
}

func NewMiddlewareManager(cfg *config.Config) *MiddlewareManager {
	return &MiddlewareManager{cfg}
}
