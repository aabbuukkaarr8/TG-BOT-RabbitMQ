package apiserver

import (
	"github.com/aabbuukkaarr8/TG-BOT/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *gin.Engine
}

func New(config *Config) *APIServer {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	return &APIServer{
		config: config,
		logger: logger,
		router: gin.Default(),
	}

}
func (s *APIServer) Run() error {
	if err := s.configLogger(); err != nil {
		return err
	}

	s.logger.Info("Starting API server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}
func (s *APIServer) configLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) ConfigureRouter(handler *handler.Handler) {
	s.router.POST("/notify", handler.Create)
	s.router.DELETE("/notify/:id", handler.Delete)
	s.router.GET("/notify/:id", handler.Status)
}
