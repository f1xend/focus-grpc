package usecase

import (
	"context"
	"fmt"
	"github.com/f1xend/focus-grpc/internal/illusionist/domain"
	"github.com/f1xend/focus-grpc/internal/illusionist/repository"
)

type Logger interface {
	Println(v ...any)
}

type Show struct {
	source repository.Rabbit
	logger Logger
}

func NewShow(source repository.Rabbit, logger Logger) Show {
	return Show{
		source: source,
		logger: logger,
	}
}

func (s Show) Rabbit(ctx context.Context) domain.Rabbit {
	res, err := s.source.Show(ctx, 1)
	if err != nil {
		s.logger.Println(fmt.Errorf("err: %w", err))
		return domain.Rabbit{}
	}
	return res
}
