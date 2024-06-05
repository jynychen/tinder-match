package main

import (
	"tinder-match/internal/config"
	"tinder-match/internal/http"
	"tinder-match/internal/logger"
	"tinder-match/internal/server"
)

// @title           tinder-match
// @version         1.0
// @description     server for the Tinder matching system
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	logger, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}

	h := http.NewEngine(cfg, logger)
	s := server.NewServer(cfg, logger, h)

	s.RunAsync()
	defer s.Shutdown()

	s.Trap()
}
