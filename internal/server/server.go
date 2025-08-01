package server

import (
	"awesomeProject/internal/logger"
	"awesomeProject/internal/server/controllers"
	"awesomeProject/internal/service"
	"awesomeProject/pkg/config"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server interface {
	Run() error
	Shutdown(ctx context.Context) error
}

type server struct {
	service     service.Service
	controllers *controllers.Controllers
	router      *gin.Engine
	logger      logger.Logger
	httpServer  *http.Server
	config      config.Config
}

func NewServer(service service.Service, router *gin.Engine, config config.Config, logger logger.Logger) Server {
	return &server{
		service:     service,
		controllers: controllers.NewControllers(service, logger),
		router:      router,
		config:      config,
		logger:      logger,
	}
}
func (s *server) SetupRoutes() {
	s.router.GET("/get", s.controllers.RandomNumberControllers.GetRTPNumber)
}

func (s *server) Run() error {
	s.logger.Info("server is starting")
	s.SetupRoutes()
	s.httpServer = &http.Server{
		Addr:    s.config.ServerAddress,
		Handler: s.router,
	}
	if err := s.httpServer.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
func (s *server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
