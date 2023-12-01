package controller

import (
	"context"
	"github.com/f1xend/focus-grpc/pkg/hat"
)

func (c *Controller) Rabbit(ctx context.Context, req *hat.Number) (*hat.Rabbit, error) {
	r := c.rabbitUseCase.RandomRabbit(ctx)
	res := &hat.Rabbit{
		Color: r.Color,
	}
	return res, nil
}
