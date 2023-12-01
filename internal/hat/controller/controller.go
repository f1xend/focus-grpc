package controller

import (
	"github.com/f1xend/focus-grpc/internal/hat/usecase"
	"github.com/f1xend/focus-grpc/pkg/hat"
)

type Logger interface {
	Println(v ...any)
}

type Controller struct {
	hat.HatServer
	rabbitUseCase usecase.Rabbit
	logger        Logger
}

func NewController(rabbit usecase.Rabbit, l Logger) *Controller {
	return &Controller{
		rabbitUseCase: rabbit,
		logger:        l,
	}
}
