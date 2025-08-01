package service

import (
	"awesomeProject/internal/logger"
	"awesomeProject/pkg/config"
	"gonum.org/v1/gonum/stat/distuv"
	"math"
)

const (
	A = iota + 1 //Левая граница интервала
	B = 10000    //Правая граница интервала
)

type Service interface {
	RandomMultiplier() float64
}
type rtpService struct {
	config config.Config
	logger logger.Logger
}

func NewService(config config.Config, logger logger.Logger) Service {
	return &rtpService{
		config: config,
		logger: logger,
	}
}
func (r *rtpService) RandomMultiplier() float64 {
	r.logger.Info("generating number")
	//формулу вывел в пояснении
	c := math.Sqrt(float64(A*A + 2*(B-A)*r.config.RTPNumber))
	normal := distuv.Normal{
		Mu:    c,
		Sigma: math.Sqrt(c),
	}
	a := normal.Rand()
	//проверка чтоб число не выходило за границы
	if a < A {
		return A
	} else if a > B {
		return B
	}
	return a

}
