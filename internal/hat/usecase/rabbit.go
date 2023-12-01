package usecase

import (
	"context"
	"github.com/f1xend/focus-grpc/internal/hat/domain"
	"github.com/f1xend/focus-grpc/internal/hat/repository"
)

type Logger interface {
	Println(v ...any)
}

type Rabbit struct {
	source repository.Rabbit
	logger Logger
}

func NewRabbit(rabbit repository.Rabbit, l Logger) Rabbit {
	return Rabbit{
		source: rabbit,
		logger: l,
	}
}

func (r Rabbit) RandomRabbit(ctx context.Context) domain.Rabbit {
	random := Random{}
	number := random.Random(1, 5)
	rabbit := r.source.Show(ctx, number)
	return rabbit
}
