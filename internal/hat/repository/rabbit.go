package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/f1xend/focus-grpc/internal/hat/domain"
)

type Logger interface {
	Println(v ...any)
}

type Rabbit struct {
	client *sql.DB
	logger Logger
}

func NewRabbit(db *sql.DB, logger Logger) Rabbit {
	return Rabbit{
		client: db,
		logger: logger,
	}
}

func (r Rabbit) Show(ctx context.Context, id int) domain.Rabbit {
	var color string

	q := "select color from rabbit where id=$1"
	row := r.client.QueryRowContext(ctx, q, id)

	if err := row.Scan(&color); err != nil {
		if err == sql.ErrNoRows {
			r.logger.Println(fmt.Errorf("no rabbit with id %d\n", id))
		} else {
			r.logger.Println(fmt.Errorf("query error: %v\n", err))
		}
		return domain.Rabbit{}
	}
	return domain.Rabbit{
		Color: color,
	}
}
