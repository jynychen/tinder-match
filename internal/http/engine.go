package http

import (
	"time"
	"tinder-match/internal/config"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewEngine(cfg *config.Config, logger *zap.Logger) *gin.Engine {
	gin.SetMode(cfg.GinMode)

	engine := gin.New()
	engine.Use(
		ginzap.Ginzap(logger, time.RFC3339, true),
		ginzap.RecoveryWithZap(logger, true),
	)

	return engine
}
