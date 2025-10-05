package apiserver

import (
	"net/http"

	"github.com/aabbuukkaarr8/TG-BOT/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/wb-go/wbf/zlog"
)

type APIServer struct {
	config *Config
	router *gin.Engine
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: gin.Default(),
	}

}
func (s *APIServer) Run() error {
	zlog.Logger.Info().Msg("Starting API server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}
func (s *APIServer) configLogger() error { return nil }

func (s *APIServer) ConfigureRouter(handler *handler.Handler) {
	s.router.POST("/notify", handler.Create)
	s.router.DELETE("/notify/:id", handler.Delete)
	s.router.GET("/notify/:id", handler.Status)
}
