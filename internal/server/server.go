package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"tinder-match/internal/config"
	"tinder-match/internal/controller"
	"tinder-match/internal/route"
	"tinder-match/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	cfg    *config.Config
	logger *zap.Logger
	http   *http.Server
}

func NewServer(
	cfg *config.Config,
	logger *zap.Logger,
	engine *gin.Engine,
) *Server {
	matchSvc := service.NewMatchService()
	matchCtrlV1 := controller.NewControllerV1(matchSvc)

	route.RegisterRESTfulV1(engine.Group("/api"), matchCtrlV1)

	return &Server{
		cfg:    cfg,
		logger: logger,
		http: &http.Server{
			Addr:    cfg.ServeAddr,
			Handler: engine,
		},
	}
}

func (s *Server) RunAsync() {
	s.logger.Info("Server running", zap.String("address", s.cfg.ServeAddr))
	go func() {
		if err := s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("Server start error", zap.Error(err))
		}
	}()
}

func (s *Server) Trap() {
	sigint := make(chan os.Signal, 1)

	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
	<-sigint

	s.logger.Info("Server exiting")
}

func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.http.Shutdown(ctx); err != nil {
		s.logger.Fatal("Server shutdown error", zap.Error(err))
	}

	s.logger.Info("Server exited")
}
