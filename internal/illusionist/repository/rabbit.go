package repository

import (
	"context"
	"fmt"
	"github.com/f1xend/focus-grpc/internal/illusionist/domain"
	"github.com/f1xend/focus-grpc/pkg/hat"
)

type Logger interface {
	Println(v ...any)
}

type Rabbit struct {
	log    Logger
	source hat.HatClient
}

func NewRabbit(log Logger, source hat.HatClient) Rabbit {
	return Rabbit{
		log:    log,
		source: source,
	}
}

func (r Rabbit) Show(ctx context.Context, id int) (domain.Rabbit, error) {
	req := hat.Number{}
	resp, err := r.source.Rabbit(ctx, &req)
	if err != nil {
		r.log.Println(fmt.Errorf("error grpc show: %w", err))
		return domain.Rabbit{}, nil
	}

	res := domain.Rabbit{
		Color: resp.Color,
	}

	return res, nil
}
