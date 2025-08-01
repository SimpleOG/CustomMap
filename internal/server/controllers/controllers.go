package controllers

import (
	"awesomeProject/internal/logger"
	"awesomeProject/internal/service"
)

type Controllers struct {
	RandomNumberControllers
}

func NewControllers(service service.Service, logger logger.Logger) *Controllers {
	return &Controllers{
		NewRandomNumberControllers(service, logger),
	}
}
