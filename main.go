package main

import (
	"awesomeProject/internal/logger"
	"awesomeProject/internal/server"
	"awesomeProject/internal/service"
	"awesomeProject/pkg/config"
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var rtp = flag.Float64("rtp", 0.4, "number for multiplier generation")
	flag.Parse()
	if *rtp > 1 {
		*rtp = 1
	}
	if *rtp <= 0 {
		*rtp = 0.1
	}
	logger, err := logger.NewLogger(0)
	if err != nil {
		log.Println("logger wasn't created via error: ", zap.Error(err))
		return
	}
	config, err := config.NewConfig("./", "env", "app", *rtp)
	if err != nil {
		logger.Error("config wasn't created via error: ", zap.Error(err))
		return
	}
	service := service.NewService(config, logger)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	server := server.NewServer(service, router, config, logger)
	errChan := make(chan error, 1)
	go func() {
		err = server.Run()
		if err != nil {
			errChan <- err
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case err = <-errChan:
		logger.Error("server error occured : ", zap.Error(err))
		os.Exit(1)
	case <-quit:
		logger.Info("shutdown started ")
		if err := server.Shutdown(ctx); err != nil {
			logger.Error("error while trying to shutdown server :", zap.Error(err))
			os.Exit(1)
		}
	}
}
