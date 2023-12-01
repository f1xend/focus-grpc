package controller

import "github.com/f1xend/focus-grpc/internal/illusionist/usecase"

type Logger interface {
	Println(v ...any)
}

type Controller struct {
	log  Logger
	show usecase.Show
}

func NewController(l Logger, s usecase.Show) Controller {
	return Controller{
		log:  l,
		show: s,
	}
}
