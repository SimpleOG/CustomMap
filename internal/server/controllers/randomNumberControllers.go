package controllers

import (
	"awesomeProject/internal/logger"
	"awesomeProject/internal/service"
	"github.com/gin-gonic/gin"
)

type RandomNumberControllers interface {
	GetRTPNumber(ctx *gin.Context)
}
type RTPHandlers struct {
	service service.Service
	logger  logger.Logger
}

func NewRandomNumberControllers(service service.Service, logger logger.Logger) RandomNumberControllers {
	return &RTPHandlers{
		service: service,
		logger:  logger,
	}
}
func (r *RTPHandlers) GetRTPNumber(ctx *gin.Context) {
	r.logger.Info("Request is accepted")
	number := r.service.RandomMultiplier()
	ctx.JSON(200, gin.H{"result": number})
}
