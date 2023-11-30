package controller

import (
	"context"

	"github.com/f1xend/focus-grpc/pkg/hat"
)

// Healthz выполняет проверку состояния здоровья.
func (c *Controller) Healthz(ctx context.Context, req *hat.HealthzRequest) (*hat.HealthzResponse, error) {
	return &hat.HealthzResponse{}, nil
}
